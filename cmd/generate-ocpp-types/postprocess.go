package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

// attributeStructs lists inner struct types whose fields should be encoded as XML attributes rather than child
// elements. This handles the OCPP 1.5 <value context="..." format="...">chardata</value> wire form, which differs from
// the OCPP 1.6+ representation where each metadata property is a child element. The Go struct itself is identical
// between JSON and XML; only the XML tags differ.
//
// Keyed by generated Go type name; value lists field names that should become XML attributes and, if non-empty, the
// field that should carry the element's character data.
type attributeSpec struct {
	attrFields  []string // fields to tag as XML attributes
	charDataFor string   // field to tag as chardata (empty if not applicable)
}

// attributeSpecs describes any struct that uses the XML-attribute wire form. Currently only OCPP 1.5 uses this shape;
// OCPP 1.6+ moved everything to child elements. Two struct variants exist in 1.5 because the JSON schema inlines the
// same meter-value shape inside StopTransaction.transactionData rather than $ref'ing MeterValueType, so the code
// generator hoists it to its own anonymous type name.
var attributeSpecs = map[string]map[string]attributeSpec{
	"1.5": {
		"MeterValueType": {
			attrFields:  []string{"Context", "Format", "Location", "Measurand", "Unit"},
			charDataFor: "Value",
		},
		"StopTransactionTransactionDataElemValuesElemValueElem": {
			attrFields:  []string{"Context", "Format", "Location", "Measurand", "Unit"},
			charDataFor: "Value",
		},
	},
}

// postprocessForSOAP mutates a generated schema.go so the same struct set works for OCPP-J (JSON) and OCPP-S (SOAP/XML)
// transports.
//
// It does four things:
//  1. Rewrites `type Foo map[string]interface{}` to `type Foo struct{}` for known message PDUs so an XMLName field can
//     be attached (encoding/xml cannot annotate a map).
//  2. Inserts `XMLName xml.Name` on each message PDU struct, tagged with the namespace URN and lower-camelCase element
//     name pulled from the WSDL.
//  3. Rewrites tags on version-specific attribute-style inner structs (see attributeSpecs).
//  4. Reorders element fields to match the WSDL <s:sequence> declaration, since Go's encoding/xml emits child
//     elements in struct-declaration order and strict SOAP validators reject out-of-sequence children.
//
// The file is only rewritten if a change is actually made, so re-running the generator without a WSDL change is a
// no-op.
func postprocessForSOAP(filename, version string, bindings wsdlBindings, sequences []wsdlSequence) error {
	if len(bindings) == 0 {
		return nil
	}

	src, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Deliberately parse without comments. go-jsonschema emits per-field descriptor comments (e.g. "IDTag corresponds
	// to the JSON schema field \"idTag\".") that carry no useful information beyond the JSON tag. Keeping them makes it
	// hard to inject XMLName as the first struct field: the printer re-anchors the pre-existing comments around our new
	// field's type expression, producing broken output like `xml.\n// old comment\nName`. Dropping them sidesteps that
	// entirely.
	fset := token.NewFileSet()
	fileAST, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return fmt.Errorf("parsing %s: %w", filename, err)
	}

	changed := false

	// Pass 1: convert message PDUs declared as `map[string]interface{}` into empty structs. go-jsonschema emits maps
	// for object schemas with no declared properties (e.g. Heartbeat). Maps cannot carry an XMLName tag, so we widen
	// them to empty structs first.
	for _, decl := range fileAST.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, wanted := lookupBinding(bindings, ts.Name.Name); !wanted {
				continue
			}
			if _, isMap := ts.Type.(*ast.MapType); isMap {
				ts.Type = &ast.StructType{Fields: &ast.FieldList{}}
				changed = true
			}
		}
	}

	// Pass 2: attach XMLName to each message PDU struct so encoding/xml emits the SOAP element name and namespace
	// correctly. For bi-directional messages (DataTransfer), the tag carries only the element name; callers pin the
	// namespace at marshal time via generated helper methods.
	var bidirectionalStructs []string // Go type names of structs that need direction helpers
	for _, decl := range fileAST.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			binding, wanted := lookupBinding(bindings, ts.Name.Name)
			if !wanted {
				continue
			}
			st, isStruct := ts.Type.(*ast.StructType)
			if !isStruct {
				continue
			}
			if hasFieldNamed(st, "XMLName") {
				continue
			}
			fields := []*ast.Field{xmlNameField(binding)}
			if binding.IsBidirectional() {
				// Bi-directional messages need a runtime-settable namespace. Baking it into the XMLName tag would
				// freeze the direction (tag wins over the runtime XMLName value when it has a non-empty local name), so
				// instead the direction helpers write the namespace to this Xmlns attribute field.
				fields = append(fields, xmlnsAttrField())
			}
			st.Fields.List = append(fields, st.Fields.List...)
			changed = true
			if binding.IsBidirectional() {
				bidirectionalStructs = append(bidirectionalStructs, ts.Name.Name)
			}
		}
	}

	// Pass 2b: emit direction helper methods for bi-directional messages.
	for _, name := range bidirectionalStructs {
		binding, _ := lookupBinding(bindings, name)
		fileAST.Decls = append(fileAST.Decls,
			directionMethod(name, "FromChargePoint", binding.ElementName, binding.CSNamespace),
			directionMethod(name, "FromCentralSystem", binding.ElementName, binding.CPNamespace),
		)
		changed = true
	}

	// Pass 3: rewrite tags on attribute-style structs.
	if versionSpecs, ok := attributeSpecs[version]; ok {
		for _, decl := range fileAST.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				attrSpec, wanted := versionSpecs[ts.Name.Name]
				if !wanted {
					continue
				}
				st, isStruct := ts.Type.(*ast.StructType)
				if !isStruct {
					continue
				}
				if rewriteAttributeStruct(st, attrSpec) {
					changed = true
				}
			}
		}
	}

	// Pass 4: reorder element fields to match the WSDL <s:sequence> ordering. Fields are matched to complexTypes by
	// xml-element-name set, so a struct whose fields don't collectively match any WSDL sequence is skipped.
	if len(sequences) > 0 {
		for _, decl := range fileAST.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.TYPE {
				continue
			}
			for _, spec := range gen.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, isStruct := ts.Type.(*ast.StructType)
				if !isStruct {
					continue
				}
				if reorderFieldsBySequence(st, sequences) {
					changed = true
				}
			}
		}
	}

	if !changed {
		return nil
	}

	ensureImport(fileAST, "encoding/xml")

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, fileAST); err != nil {
		return err
	}

	// Reinsert the "generated code" marker that go-jsonschema emitted originally.
	output := append([]byte("// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.\n\n"), buf.Bytes()...)
	return os.WriteFile(filename, output, 0o644)
}

// lookupBinding maps a Go type name (e.g. "Authorize" or "AuthorizeResponse") to its WSDL binding. OCPP 1.5/1.6
// JSON-side struct names drop the "Request" suffix on the request side but keep the "Response" suffix on the response
// side, so we peel Response and lookup with isResponse=true.
func lookupBinding(bindings wsdlBindings, structName string) (wsdlBinding, bool) {
	if strings.HasSuffix(structName, "Response") {
		return bindings.Get(strings.TrimSuffix(structName, "Response"), true)
	}
	return bindings.Get(structName, false)
}

func xmlNameField(binding wsdlBinding) *ast.Field {
	// json:"-" is required because xml.Name has exported Space/Local fields; without it, the JSON encoder would emit
	// "XMLName":{"Space":"...","Local":"..."} on every marshalled OCPP-J message.
	//
	// For uni-directional messages the tag carries "<namespace> <elementName>". For bi-directional messages
	// (DataTransfer) the tag carries just "<elementName>" and callers pin the namespace via the generated
	// FromChargePoint / FromCentralSystem helpers.
	var xmlTag string
	if ns := binding.DefaultNamespace(); ns != "" {
		xmlTag = ns + " " + binding.ElementName
	} else {
		xmlTag = binding.ElementName
	}
	tagValue := fmt.Sprintf("`xml:%q json:\"-\"`", xmlTag)
	return &ast.Field{
		Names: []*ast.Ident{ast.NewIdent("XMLName")},
		Type:  &ast.SelectorExpr{X: ast.NewIdent("xml"), Sel: ast.NewIdent("Name")},
		Tag:   &ast.BasicLit{Kind: token.STRING, Value: tagValue},
	}
}

// xmlnsAttrField builds an AST node for a field that emits an explicit xmlns="" attribute on the containing struct's
// XML element. Direction helpers write to it at marshal time.
func xmlnsAttrField() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{ast.NewIdent("Xmlns")},
		Type:  ast.NewIdent("string"),
		Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`xml:\"xmlns,attr,omitempty\" json:\"-\"`"},
	}
}

// directionMethod builds an AST node for a helper method that pins the SOAP target namespace of a bi-directional
// message, e.g.:
//
//	func (d *DataTransfer) FromChargePoint() *DataTransfer {
//	    d.Xmlns = "urn://Ocpp/Cs/2015/10/"
//	    return d
//	}
//
// The pointer receiver mutates the caller's value; returning the pointer lets callers chain the call straight into
// xml.Marshal.
func directionMethod(structName, methodName, _, namespace string) *ast.FuncDecl {
	receiver := &ast.Field{
		Names: []*ast.Ident{ast.NewIdent("d")},
		Type:  &ast.StarExpr{X: ast.NewIdent(structName)},
	}
	returnType := &ast.StarExpr{X: ast.NewIdent(structName)}

	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent("d"), Sel: ast.NewIdent("Xmlns")}},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", namespace)}},
	}
	ret := &ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("d")}}

	return &ast.FuncDecl{
		Recv: &ast.FieldList{List: []*ast.Field{receiver}},
		Name: ast.NewIdent(methodName),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{},
			Results: &ast.FieldList{List: []*ast.Field{{Type: returnType}}},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{assign, ret}},
	}
}

// reorderFieldsBySequence reorders the element-mode fields of a struct to match a WSDL <s:sequence> declaration.
// XMLName, xmlns attributes, other XML attributes, and chardata fields are left in their original positions; only the
// ordering of ordinary element-carrying fields is affected. The struct is matched to a complexType by comparing its
// set of xml element names to each sequence's element name set — a unique match reorders; ambiguous or missing matches
// leave the struct untouched.
func reorderFieldsBySequence(st *ast.StructType, sequences []wsdlSequence) bool {
	// Collect ordinary element-mode fields (skip XMLName, attributes, chardata, and any xml:"-" fields).
	type elementField struct {
		index   int
		xmlName string
	}
	var elems []elementField
	for i, f := range st.Fields.List {
		if len(f.Names) == 0 || f.Names[0].Name == "XMLName" {
			continue
		}
		name, isElem := parseElementXMLName(f.Tag)
		if !isElem {
			continue
		}
		elems = append(elems, elementField{index: i, xmlName: name})
	}
	if len(elems) < 2 {
		return false
	}

	nameSet := make(map[string]bool, len(elems))
	for _, e := range elems {
		nameSet[e.xmlName] = true
	}

	// Find the unique complexType sequence whose element names match this set exactly.
	var match *wsdlSequence
	for i := range sequences {
		seq := &sequences[i]
		if len(seq.ElementNames) != len(nameSet) {
			continue
		}
		ok := true
		for _, n := range seq.ElementNames {
			if !nameSet[n] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		if match != nil {
			return false // ambiguous — two complexTypes have the same field-name set
		}
		match = seq
	}
	if match == nil {
		return false
	}

	byName := make(map[string]*ast.Field, len(elems))
	for _, e := range elems {
		byName[e.xmlName] = st.Fields.List[e.index]
	}
	targetOrder := make([]*ast.Field, len(match.ElementNames))
	for i, n := range match.ElementNames {
		targetOrder[i] = byName[n]
	}

	changed := false
	for i, e := range elems {
		if st.Fields.List[e.index] != targetOrder[i] {
			changed = true
			break
		}
	}
	if !changed {
		return false
	}
	for i, e := range elems {
		st.Fields.List[e.index] = targetOrder[i]
	}
	return true
}

// parseElementXMLName returns the xml tag's element name and reports whether the field participates in XML wire
// ordering. Attribute fields, chardata fields, xml:"-" fields, and fields whose tag name contains a space
// (namespace-qualified — used by XMLName) are all considered non-ordering.
func parseElementXMLName(tag *ast.BasicLit) (string, bool) {
	if tag == nil {
		return "", false
	}
	inner, ok := unquoteTagLiteral(tag.Value)
	if !ok {
		return "", false
	}
	for _, t := range parseTagList(inner) {
		if t.key != "xml" {
			continue
		}
		name, opts, _ := strings.Cut(t.value, ",")
		if name == "" || name == "-" {
			return "", false
		}
		if strings.Contains(name, " ") {
			return "", false
		}
		optSet := strings.Split(opts, ",")
		for _, o := range optSet {
			if o == "attr" || o == "chardata" || o == "innerxml" || o == "comment" || o == "cdata" {
				return "", false
			}
		}
		return name, true
	}
	return "", false
}

func hasFieldNamed(st *ast.StructType, name string) bool {
	for _, f := range st.Fields.List {
		for _, n := range f.Names {
			if n.Name == name {
				return true
			}
		}
	}
	return false
}

// rewriteAttributeStruct rewrites XML tags on a struct so that named fields become attributes and one designated field
// carries the element's character data. JSON tags are preserved as-is.
func rewriteAttributeStruct(st *ast.StructType, spec attributeSpec) bool {
	attrSet := map[string]struct{}{}
	for _, name := range spec.attrFields {
		attrSet[name] = struct{}{}
	}

	changed := false
	for _, field := range st.Fields.List {
		if field.Tag == nil || len(field.Names) == 0 {
			continue
		}
		fieldName := field.Names[0].Name
		if _, isAttr := attrSet[fieldName]; isAttr {
			if newTag, ok := rewriteXMLTag(field.Tag.Value, attrModeTag); ok {
				field.Tag.Value = newTag
				changed = true
			}
		} else if fieldName == spec.charDataFor {
			if newTag, ok := rewriteXMLTag(field.Tag.Value, charDataTag); ok {
				field.Tag.Value = newTag
				changed = true
			}
		}
	}
	return changed
}

// tagRewriter transforms an existing xml tag value ("<name>,<opts>") into a new form.
type tagRewriter func(name, opts string) string

func attrModeTag(name, opts string) string {
	if strings.Contains(opts, "attr") {
		return name + "," + opts
	}
	if opts == "" {
		return name + ",attr"
	}
	return name + ",attr," + opts
}

func charDataTag(_, _ string) string { return ",chardata" }

// rewriteXMLTag replaces the xml tag inside the given raw struct tag literal, preserving all other tags. The raw tag
// literal includes the surrounding backticks.
func rewriteXMLTag(raw string, rewrite tagRewriter) (string, bool) {
	inner, ok := unquoteTagLiteral(raw)
	if !ok {
		return raw, false
	}
	tags := parseTagList(inner)
	found := false
	for i, t := range tags {
		if t.key != "xml" {
			continue
		}
		name, opts, _ := strings.Cut(t.value, ",")
		tags[i].value = rewrite(name, opts)
		found = true
	}
	if !found {
		return raw, false
	}
	return "`" + formatTagList(tags) + "`", true
}

type tag struct {
	key, value string
}

func parseTagList(s string) []tag {
	var out []tag
	for len(s) > 0 {
		s = strings.TrimLeft(s, " ")
		if s == "" {
			break
		}
		key, rest, ok := strings.Cut(s, ":")
		if !ok {
			break
		}
		key = strings.TrimSpace(key)
		if !strings.HasPrefix(rest, `"`) {
			break
		}
		rest = rest[1:]
		end := strings.Index(rest, `"`)
		if end < 0 {
			break
		}
		out = append(out, tag{key: key, value: rest[:end]})
		s = rest[end+1:]
	}
	return out
}

func formatTagList(tags []tag) string {
	parts := make([]string, 0, len(tags))
	for _, t := range tags {
		parts = append(parts, fmt.Sprintf("%s:%q", t.key, t.value))
	}
	return strings.Join(parts, " ")
}

func unquoteTagLiteral(raw string) (string, bool) {
	if len(raw) < 2 || raw[0] != '`' || raw[len(raw)-1] != '`' {
		return "", false
	}
	return raw[1 : len(raw)-1], true
}

// ensureImport adds `import "encoding/xml"` to the file if it is not already imported. It creates a new single-line
// import declaration to match the style go-jsonschema emits.
func ensureImport(file *ast.File, path string) {
	quoted := fmt.Sprintf("%q", path)
	for _, imp := range file.Imports {
		if imp.Path.Value == quoted {
			return
		}
	}
	imp := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: quoted}},
		},
	}
	// Insert after the last existing import so imports stay grouped at the top.
	lastImport := -1
	for i, decl := range file.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok == token.IMPORT {
			lastImport = i
		}
	}
	insertAt := lastImport + 1
	file.Decls = append(file.Decls, nil)
	copy(file.Decls[insertAt+1:], file.Decls[insertAt:])
	file.Decls[insertAt] = imp
	file.Imports = append(file.Imports, imp.Specs[0].(*ast.ImportSpec))
}

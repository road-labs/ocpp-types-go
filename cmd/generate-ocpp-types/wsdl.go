package main

import (
	"embed"
	"encoding/xml"
	"errors"
	"fmt"
	"io/fs"
	"strings"
)

//go:embed wsdl
var wsdlFS embed.FS

// wsdlBinding records the SOAP XML wire-form of a single OCPP message.
//
// Most messages are unidirectional: they appear in only one of the two WSDLs, so exactly one of CSNamespace /
// CPNamespace is set. DataTransfer is the only bi-directional message and has both populated; for those, the generator
// leaves the namespace out of the tag and emits helper methods so callers can pin the direction at marshal time.
type wsdlBinding struct {
	ElementName string // e.g. "dataTransferRequest"
	CSNamespace string // target namespace when the message is bound to CentralSystemService
	CPNamespace string // target namespace when the message is bound to ChargePointService
}

// IsBidirectional reports whether this message is bound in both the CS and CP WSDLs.
func (b wsdlBinding) IsBidirectional() bool {
	return b.CSNamespace != "" && b.CPNamespace != ""
}

// DefaultNamespace is the namespace to bake into the XMLName struct tag. For bi-directional messages it returns "" so
// the caller (via the generated helper methods) picks the direction at runtime.
func (b wsdlBinding) DefaultNamespace() string {
	if b.IsBidirectional() {
		return ""
	}
	if b.CSNamespace != "" {
		return b.CSNamespace
	}
	return b.CPNamespace
}

// wsdlMessage keys the binding map by (base name, request-or-response). The base name matches the Go type name for the
// request side ("Authorize"); the response side has "Response" appended in Go.
type wsdlMessage struct {
	Name       string
	IsResponse bool
}

type wsdlBindings map[wsdlMessage]wsdlBinding

func (b wsdlBindings) Get(name string, isResponse bool) (wsdlBinding, bool) {
	binding, ok := b[wsdlMessage{Name: name, IsResponse: isResponse}]
	return binding, ok
}

// loadWSDLBindings reads the embedded WSDL files for the given version, extracts the SOAP message element bindings, and
// merges them into a single lookup keyed by (base name, isResponse). A message present in both WSDLs (DataTransfer)
// ends up with both CS and CP namespaces recorded.
func loadWSDLBindings(version string) (wsdlBindings, error) {
	dir := fmt.Sprintf("wsdl/%s", version)
	entries, err := wsdlFS.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading wsdl dir %s: %w", dir, err)
	}

	out := wsdlBindings{}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".wsdl") {
			continue
		}
		data, err := wsdlFS.ReadFile(dir + "/" + entry.Name())
		if err != nil {
			return nil, err
		}
		perFile, err := parseWSDLBindings(data)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}
		for k, v := range perFile {
			merged := out[k]
			if merged.ElementName == "" {
				merged.ElementName = v.ElementName
			}
			if v.CSNamespace != "" {
				merged.CSNamespace = v.CSNamespace
			}
			if v.CPNamespace != "" {
				merged.CPNamespace = v.CPNamespace
			}
			out[k] = merged
		}
	}
	return out, nil
}

// parseWSDLBindings extracts the top-level element declarations from a WSDL's inline schema.
//
// A WSDL types block looks like:
//
//	<wsdl:types>
//	  <s:schema targetNamespace="urn://Ocpp/Cs/2012/06/" ...>
//	    ...
//	    <s:element name="authorizeRequest"  type="tns:AuthorizeRequest"/>
//	    <s:element name="authorizeResponse" type="tns:AuthorizeResponse"/>
//	    ...
//	  </s:schema>
//	</wsdl:types>
//
// We record the targetNamespace and each top-level element that maps to a Request or Response type.
func parseWSDLBindings(data []byte) (wsdlBindings, error) {
	var doc struct {
		Types struct {
			Schemas []struct {
				TargetNamespace string `xml:"targetNamespace,attr"`
				Elements        []struct {
					Name string `xml:"name,attr"`
					Type string `xml:"type,attr"`
				} `xml:"element"`
			} `xml:"schema"`
		} `xml:"types"`
	}

	if err := xml.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	out := wsdlBindings{}
	for _, schema := range doc.Types.Schemas {
		if schema.TargetNamespace == "" {
			continue
		}
		// OCPP WSDLs use targetNamespace URNs of the form urn://Ocpp/Cs/YYYY/MM/ (CentralSystemService)
		// or urn://Ocpp/Cp/YYYY/MM/ (ChargePointService). We infer the direction from the URN.
		isCS := strings.Contains(schema.TargetNamespace, "/Cs/")
		isCP := strings.Contains(schema.TargetNamespace, "/Cp/")
		if !isCS && !isCP {
			continue
		}
		for _, el := range schema.Elements {
			// Skip elements whose type is a primitive (like chargeBoxIdentity). We only want message PDUs.
			typeName := strings.TrimPrefix(el.Type, "tns:")
			if typeName == el.Type {
				continue
			}
			msgName, isResponse, ok := splitRequestResponse(typeName)
			if !ok {
				continue
			}
			binding := wsdlBinding{ElementName: el.Name}
			if isCS {
				binding.CSNamespace = schema.TargetNamespace
			} else {
				binding.CPNamespace = schema.TargetNamespace
			}
			out[wsdlMessage{Name: msgName, IsResponse: isResponse}] = binding
		}
	}
	return out, nil
}

// splitRequestResponse peels the Request/Response suffix off a WSDL type name and returns the base message name plus
// whether it is the response side. Unrelated type names (IdToken, IdTagInfo, etc.) return ok=false.
func splitRequestResponse(typeName string) (base string, isResponse bool, ok bool) {
	switch {
	case strings.HasSuffix(typeName, "Request"):
		return strings.TrimSuffix(typeName, "Request"), false, true
	case strings.HasSuffix(typeName, "Response"):
		return strings.TrimSuffix(typeName, "Response"), true, true
	}
	return "", false, false
}

// wsdlSequence captures the ordered child-element names declared inside a WSDL <s:complexType><s:sequence>.
// SOAP peers that validate against the WSDL will reject XML whose children arrive out of order, so this drives the
// post-processor's field-reordering pass.
type wsdlSequence struct {
	ComplexTypeName string   // e.g. "IdTagInfo"
	ElementNames    []string // ordered child element names, e.g. ["status", "expiryDate", "parentIdTag"]
}

// loadWSDLSequences returns every complexType sequence from every WSDL file for the given version. A complexType
// declared identically in both the CS and CP files (the shared types like IdTagInfo) is folded to one entry.
func loadWSDLSequences(version string) ([]wsdlSequence, error) {
	dir := fmt.Sprintf("wsdl/%s", version)
	entries, err := wsdlFS.ReadDir(dir)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading wsdl dir %s: %w", dir, err)
	}

	seen := map[string]bool{}
	var out []wsdlSequence
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".wsdl") {
			continue
		}
		data, err := wsdlFS.ReadFile(dir + "/" + entry.Name())
		if err != nil {
			return nil, err
		}
		perFile, err := parseWSDLSequences(data)
		if err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}
		for _, seq := range perFile {
			if seen[seq.ComplexTypeName] {
				continue
			}
			seen[seq.ComplexTypeName] = true
			out = append(out, seq)
		}
	}
	return out, nil
}

// parseWSDLSequences walks every <s:complexType> in the WSDL's inline schema and records its <s:sequence> child
// element names in declaration order. complexTypes without a sequence (empty request/response bodies) are skipped.
func parseWSDLSequences(data []byte) ([]wsdlSequence, error) {
	var doc struct {
		Types struct {
			Schemas []struct {
				ComplexTypes []struct {
					Name     string `xml:"name,attr"`
					Sequence struct {
						Elements []struct {
							Name string `xml:"name,attr"`
						} `xml:"element"`
					} `xml:"sequence"`
				} `xml:"complexType"`
			} `xml:"schema"`
		} `xml:"types"`
	}

	if err := xml.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	var out []wsdlSequence
	for _, schema := range doc.Types.Schemas {
		for _, ct := range schema.ComplexTypes {
			if ct.Name == "" || len(ct.Sequence.Elements) == 0 {
				continue
			}
			seq := wsdlSequence{ComplexTypeName: ct.Name}
			for _, el := range ct.Sequence.Elements {
				seq.ElementNames = append(seq.ElementNames, el.Name)
			}
			out = append(out, seq)
		}
	}
	return out, nil
}

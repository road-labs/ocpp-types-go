package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/mitchellh/reflectwalk"
)

const patchesDir = "patches"

type schema struct {
	filePath string
	doc      map[string]any
}

func main() {
	var dirPath string
	flag.StringVar(&dirPath, "dir-path", "", "")
	flag.Parse()

	err := preProcess(dirPath)
	if err != nil {
		panic(err)
	}
}

// PreProcess looks to extract common definitions into a separate file. The JSON schema for OCPP 2.0.1 has a quirk in
// that many definitions are copy/pasted over and over. This causes issues during code generation, as the generator
// treats them all as different types - which they technically are according to the schema - even though they all share
// the same fields and types. By extracting out definitions, these can become shared - the $refs values are simply
// updated to reference the definition in a shared file, instead of the local one.
func preProcess(dirPath string) error {
	schemas := make([]schema, 0)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	patches, err := getPatches(filepath.Join(dirPath, patchesDir))
	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}

		// ignore directories (like patches)
		if info.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, info.Name())
		s, err := unmarshalSchema(filePath, patches[info.Name()])
		if err != nil {
			return err
		}
		schemas = append(schemas, s)
	}

	definitions := make(map[string]map[string]any)
	for _, s := range schemas {
		if schemaDefs, found := s.doc["definitions"]; found {
			// extract definitions
			for key, value := range schemaDefs.(map[string]any) {
				schemaDef := value.(map[string]any)
				delete(schemaDef, "description") // description is the only field that varies amongst the duplicate values

				if existing, seen := definitions[key]; seen {
					// ensure that we have an identical definition
					if !reflect.DeepEqual(existing, schemaDef) {
						return fmt.Errorf("types do not match: %+v vs %+v", existing, schemaDef)
					}
				} else {
					definitions[key] = schemaDef
				}
			}
			delete(s.doc, "definitions") // clear the definitions field, as all $refs values will point to the single file
		}
		if err = reflectwalk.Walk(s.doc, walker{}); err != nil {
			return err
		}
		if err = marshalSchema(s); err != nil {
			return err
		}
	}

	if len(definitions) > 0 {
		// create the single definitions file
		if err = os.Mkdir(filepath.Join(dirPath, "common"), 0o777); err != nil {
			return err
		}
		if err = marshalSchema(schema{
			filePath: filepath.Join(dirPath, "common", "Definitions.json"),
			doc: map[string]any{
				"$schema":     "http://json-schema.org/draft-06/schema#",
				"$id":         "urn:roadio:OCPP:Definitions",
				"definitions": definitions,
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

func unmarshalSchema(filePath string, patch jsonpatch.Patch) (schema, error) {
	f, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return schema{}, err
	}

	if patch != nil {
		patched, err := patch.Apply(f)
		if err != nil {
			return schema{}, err
		}

		f = patched
	}

	var doc map[string]any
	if err = json.Unmarshal(f, &doc); err != nil {
		return schema{}, err
	}
	return schema{
		filePath: filePath,
		doc:      doc,
	}, nil
}

func marshalSchema(s schema) error {
	f, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(s.doc)
}

// get and prepare patches
func getPatches(patchDir string) (map[string]jsonpatch.Patch, error) {
	patches := make(map[string]jsonpatch.Patch)

	// if dir does not exist, return empty patches
	if _, err := os.Stat(patchDir); os.IsNotExist(err) {
		return patches, nil
	}

	entries, err := os.ReadDir(patchDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filePath := filepath.Clean(filepath.Join(patchDir, entry.Name()))
		f, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		patchOps, err := jsonpatch.DecodePatch(f)
		if err != nil {
			return nil, err
		}

		patches[entry.Name()] = patchOps
	}

	return patches, nil
}

type walker struct{}

func (walker) Map(reflect.Value) error {
	return nil
}

//nolint:golint,unparam
func (walker) MapElem(m, k, v reflect.Value) error {
	if k.String() == "$ref" {
		// re-write the reference to point to the shared definitions file
		m.SetMapIndex(k, reflect.ValueOf("common/Definitions.json"+v.Interface().(string)))
	}
	return nil
}

package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/mitchellh/reflectwalk"
)

//go:embed schemas
var schemas embed.FS

// buildPatches loads patches for the given version from two sources:
//   - embedded: schemas/<version>/patches/ within the binary
//   - filesystem: a flat directory provided via --patches-dir
//
// The map key is the schema filename the patch applies to (e.g. "DeleteCertificate.json").
func buildPatches(version, patchesDir string) (map[string]jsonpatch.Patch, error) {
	patches := make(map[string]jsonpatch.Patch)

	embeddedPatchDir := fmt.Sprintf("schemas/%s/patches", version)
	entries, err := schemas.ReadDir(embeddedPatchDir)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, fmt.Errorf("reading embedded patches for %s: %w", version, err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, err := schemas.ReadFile(embeddedPatchDir + "/" + entry.Name())
		if err != nil {
			return nil, err
		}
		patch, err := jsonpatch.DecodePatch(data)
		if err != nil {
			return nil, fmt.Errorf("decoding embedded patch %s: %w", entry.Name(), err)
		}
		patches[entry.Name()] = patch
	}

	if patchesDir == "" {
		return patches, nil
	}

	entries, err = os.ReadDir(patchesDir)
	if err != nil {
		return nil, fmt.Errorf("reading --patches-dir: %w", err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(patchesDir, entry.Name()))
		if err != nil {
			return nil, err
		}
		patch, err := jsonpatch.DecodePatch(data)
		if err != nil {
			return nil, fmt.Errorf("decoding patch %s: %w", entry.Name(), err)
		}
		patches[entry.Name()] = append(patches[entry.Name()], patch...)
	}

	return patches, nil
}

// preprocessSchemas reads JSON schemas from srcFS, applies patches, extracts shared definitions into a common file,
// rewrites $ref pointers, and writes results to dstDir.
//
// The common definitions extraction addresses a quirk in OCPP 2.x schemas where many definitions are copy/pasted across
// files. Without extraction, the code generator treats them as distinct types, producing duplicates in the output.
func preprocessSchemas(srcFS fs.FS, dstDir string, patches map[string]jsonpatch.Patch) error {
	entries, err := fs.ReadDir(srcFS, ".")
	if err != nil {
		return err
	}

	type schemaEntry struct {
		name string
		doc  map[string]any
	}

	var schemaEntries []schemaEntry
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		data, err := fs.ReadFile(srcFS, entry.Name())
		if err != nil {
			return err
		}

		if patch, ok := patches[entry.Name()]; ok {
			if data, err = patch.Apply(data); err != nil {
				return fmt.Errorf("applying patch to %s: %w", entry.Name(), err)
			}
		}

		var doc map[string]any
		if err = json.Unmarshal(data, &doc); err != nil {
			return fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}
		schemaEntries = append(schemaEntries, schemaEntry{name: entry.Name(), doc: doc})
	}

	// Extract common definitions across schemas. "description" is the only field that legitimately varies across copies
	// of the same definition, so it is stripped before the equality check.
	definitions := make(map[string]map[string]any)
	for _, s := range schemaEntries {
		defs, ok := s.doc["definitions"]
		if !ok {
			continue
		}
		for key, val := range defs.(map[string]any) {
			def := val.(map[string]any)
			delete(def, "description")
			if existing, seen := definitions[key]; seen {
				if !reflect.DeepEqual(existing, def) {
					return fmt.Errorf("conflicting definitions for %q", key)
				}
			} else {
				definitions[key] = def
			}
		}
		delete(s.doc, "definitions")
	}

	// Rewrite all $ref pointers to reference the shared definitions file, then write
	// the processed schemas to the destination directory.
	for _, s := range schemaEntries {
		if err = reflectwalk.Walk(s.doc, refRewriter{}); err != nil {
			return err
		}
		if err = writeJSON(filepath.Join(dstDir, s.name), s.doc); err != nil {
			return err
		}
	}

	if len(definitions) == 0 {
		return nil
	}

	if err = os.Mkdir(filepath.Join(dstDir, "common"), 0o777); err != nil {
		return err
	}
	return writeJSON(filepath.Join(dstDir, "common", "Definitions.json"), map[string]any{
		"$schema":     "http://json-schema.org/draft-06/schema#",
		"$id":         "urn:ocpp-types-go:OCPP:Definitions",
		"definitions": definitions,
	})
}

func writeJSON(path string, v any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

type refRewriter struct{}

func (refRewriter) Map(reflect.Value) error { return nil }

func (refRewriter) MapElem(m, k, v reflect.Value) error {
	if k.String() == "$ref" {
		m.SetMapIndex(k, reflect.ValueOf("common/Definitions.json"+v.Interface().(string)))
	}
	return nil
}

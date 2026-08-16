package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atombender/go-jsonschema/pkg/generator"
	"github.com/urfave/cli/v3"
)

type config struct {
	version     string
	packagePath string
	patchesDir  string
	outputDir   string
	soap        bool
}

func main() {
	var cfg config

	app := &cli.Command{
		Name:  "generate-types",
		Usage: "Generate Go types for a specific OCPP version",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Usage:       "OCPP version to generate",
				Required:    true,
				Destination: &cfg.version,
			},
			&cli.StringFlag{
				Name:        "package",
				Usage:       "Go package path for the generated file (e.g. github.com/org/repo/gen/ocpp16)",
				Required:    true,
				Destination: &cfg.packagePath,
			},
			&cli.StringFlag{
				Name:        "patches-dir",
				Usage:       "Flat directory of extra patch files to apply (filename must match schema name)",
				Destination: &cfg.patchesDir,
			},
			&cli.StringFlag{
				Name:        "output-dir",
				Usage:       "Directory where schema.go will be written",
				Required:    true,
				Destination: &cfg.outputDir,
			},
			&cli.BoolFlag{
				Name:        "soap",
				Usage:       "Also generate xml struct tags and inject XMLName fields from the vendored WSDLs so the types work for OCPP-S (SOAP) as well as OCPP-J (JSON)",
				Destination: &cfg.soap,
			},
		},
		Action: func(ctx context.Context, _ *cli.Command) error {
			return run(ctx, cfg)
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context, cfg config) error {
	fmt.Printf("Generating OCPP %s...\n", cfg.version)

	patches, err := buildPatches(cfg.version, cfg.patchesDir)
	if err != nil {
		return fmt.Errorf("loading patches: %w", err)
	}

	schemasSubFS, err := fs.Sub(schemas, fmt.Sprintf("schemas/%s/schema", cfg.version))
	if err != nil {
		return fmt.Errorf("schemas not found for version %s: %w", cfg.version, err)
	}

	tmpDir, err := os.MkdirTemp("", "ocppgen-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	if err = preprocessSchemas(schemasSubFS, tmpDir, patches); err != nil {
		return fmt.Errorf("preprocessing: %w", err)
	}

	outputFile := filepath.Join(cfg.outputDir, "schema.go")

	if err = generate(tmpDir, cfg.packagePath, outputFile, cfg.soap); err != nil {
		return fmt.Errorf("generating: %w", err)
	}

	if cfg.soap {
		bindings, err := loadWSDLBindings(cfg.version)
		if err != nil {
			return fmt.Errorf("loading wsdl bindings: %w", err)
		}
		sequences, err := loadWSDLSequences(cfg.version)
		if err != nil {
			return fmt.Errorf("loading wsdl sequences: %w", err)
		}
		if err = postprocessForSOAP(outputFile, cfg.version, bindings, sequences); err != nil {
			return fmt.Errorf("postprocessing for soap: %w", err)
		}
	}

	fmt.Printf("  -> %s\n", outputFile)
	return nil
}

func generate(tmpDir, pkgPath, outputFile string, soap bool) error {
	tags := []string{"json"}
	if soap {
		tags = append(tags, "xml")
	}
	cfg := generator.Config{
		Warner: func(msg string) {
			_, _ = fmt.Fprintf(os.Stderr, "warning: %s\n", msg)
		},
		DefaultPackageName: pkgPath,
		DefaultOutputName:  outputFile,
		Capitalizations:    []string{"ID"},
		Tags:               tags,
		ResolveExtensions:  []string{"json"},
	}

	gen, err := generator.New(cfg)
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		if err = gen.DoFile(filepath.Join(tmpDir, entry.Name())); err != nil {
			return fmt.Errorf("processing %s: %w", entry.Name(), err)
		}
	}

	sources, err := gen.Sources()
	if err != nil {
		return err
	}

	for fileName, source := range sources {
		if err = os.MkdirAll(filepath.Dir(fileName), 0o755); err != nil {
			return err
		}
		if err = os.WriteFile(fileName, source, 0o644); err != nil {
			return err
		}
	}

	return nil
}

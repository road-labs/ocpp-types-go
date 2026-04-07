# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Install code generation tools (required before first generate)
task install-tools

# Regenerate all Go types from JSON schemas
task generate

# Build the module
go build ./...

# Run tests
go test ./...
```

## Architecture

This is a Go type definition library for the OCPP (Open Charge Point Protocol) used in EV charging infrastructure. It provides strongly-typed Go structs for all OCPP message types across four protocol versions (1.5, 1.6, 2.0.1, 2.1).

### Package Layout

- **Root package (`ocpp`)** — shared abstractions: version constants (`version.go`), error codes (`error.go`), and operation routing (`operation.go`)
- **`gen/ocpp15`, `gen/ocpp16`, `gen/ocpp201`, `gen/ocpp21`** — generated `schema.go` files containing all message structs for each protocol version; package names match the directory (e.g. `package ocpp15`)
- **`schemas/`** — source JSON schemas organized by version (`1.5/schema/`, `1.6/schema/`, etc.), including JSON patch files for schema customization
- **`tools/schemas/`** — schema preprocessing tool run before `go-jsonschema` code generation

### Code Generation Pipeline

The `schema.go` files are **auto-generated** — never edit them manually. The pipeline is:

1. Copy JSON schemas from `schemas/<version>/schema/` to `tmp/ocppgen/`
2. Run `tools/schemas` preprocessor — extracts duplicate definitions across schemas into a shared `common/Definitions.json`, rewrites `$ref` pointers to avoid duplicate type generation
3. Run `go-jsonschema` to generate Go structs with `--capitalization ID` and `--tags json`
4. Output lands in `gen/ocpp<version>/schema.go`

JSON patches in `schemas/<version>/schema/patches/` are applied during preprocessing to customize schemas before generation.

### Naming Conventions

OCPP 1.5/1.6 request types are named without suffix (e.g., `CancelReservation`, `Reset`). OCPP 2.0.1/2.1 request types use the `Request` suffix (e.g., `CancelReservationRequest`, `ResetRequest`). This reflects the upstream schema naming conventions.

### Operation Routing

`OperationToRequestStruct(operation, version)` in `operation.go` maps a `CentralSystemToChargerPointOperation` string and a `Version` to an empty struct pointer of the correct versioned type. When adding a new operation, update all four `operationToRequestStruct<version>` functions and the `IsValidCentralSystemToChargerPointOperation` switch.

### Adding a New OCPP Version or Operation

1. Add the JSON schema files under `schemas/<version>/schema/`
2. Create the output directory `gen/ocpp<version>/`
3. Run `task generate` to produce `schema.go`
4. Add the new version constant to `version.go`
5. Add operation constants to `operation.go` and update all routing switch statements
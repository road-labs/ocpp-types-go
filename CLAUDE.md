# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
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

- **Root package (`ocpp`)** — shared abstractions: version constants (`version.go`) and action routing (`action.go`)
- **`gen/ocpp15`, `gen/ocpp16`, `gen/ocpp201`, `gen/ocpp21`** — generated `schema.go` files containing all message structs for each protocol version; package names match the directory (e.g. `package ocpp15`)
- **`cmd/generate-ocpp-types/`** — the code generator. Holds `main.go` (schema preprocessor + `go-jsonschema` driver) and the embedded `schemas/` tree organised by version (`1.5/schema/`, `1.6/schema/`, etc.) along with optional JSON patch files per version.

### Code Generation Pipeline

The `schema.go` files are **auto-generated** — never edit them manually. The pipeline (implemented in `cmd/generate-ocpp-types/main.go`) is:

1. Copy JSON schemas from `cmd/generate-ocpp-types/schemas/<version>/schema/` into a temporary working directory
2. Apply any JSON patches for that version, then extract duplicate definitions across schemas into a shared `common/Definitions.json` and rewrite `$ref` pointers to avoid duplicate type generation
3. Run `go-jsonschema` to generate Go structs with `--capitalization ID` and `--tags json`
4. Output lands in `gen/ocpp<version>/schema.go`

JSON patches live at `cmd/generate-ocpp-types/schemas/<version>/patches/` (currently only OCPP 1.6 ships a patch) and are applied during preprocessing to customise schemas before generation.

### Naming Conventions

OCPP 1.5/1.6 request types are named without suffix (e.g., `CancelReservation`, `Reset`). OCPP 2.0.1/2.1 request types use the `Request` suffix (e.g., `CancelReservationRequest`, `ResetRequest`). This reflects the upstream schema naming conventions.

### Action Routing

`ActionToRequestStruct(action, version)` in `action.go` maps a `CentralSystemToChargingStationAction` string and a `Version` to an empty struct pointer of the correct versioned type. When adding a new action, update all four `ActionToRequestStruct<version>` functions and the `IsValidCentralSystemToChargingStationAction` switch.

### Adding a New OCPP Version or Action

1. Add the JSON schema files under `cmd/generate-ocpp-types/schemas/<version>/schema/`
2. Create the output directory `gen/ocpp<version>/`
3. Add the new version to the `for` list in `Taskfile.yml` and run `task generate` to produce `schema.go`
4. Add the new version constant to `version.go`
5. Add action constants to `action.go` and update all routing switch statements
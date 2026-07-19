# ocpp-types-go

Go type definitions for the [OCPP](https://www.openchargealliance.org/) (Open Charge Point Protocol), used in EV charging infrastructure. Provides strongly-typed structs for all OCPP message types across four protocol versions.

## Supported versions

| Version | Package |
|---------|---------|
| OCPP 1.5 | `github.com/e-flux-platform/ocpp-types-go/gen/ocpp15` |
| OCPP 1.6 | `github.com/e-flux-platform/ocpp-types-go/gen/ocpp16` |
| OCPP 2.0.1 | `github.com/e-flux-platform/ocpp-types-go/gen/ocpp201` |
| OCPP 2.1 | `github.com/e-flux-platform/ocpp-types-go/gen/ocpp21` |

## Installation

```bash
go get github.com/e-flux-platform/ocpp-types-go
```

## Usage

### Importing versioned types

```go
import (
    "github.com/e-flux-platform/ocpp-types-go"
    "github.com/e-flux-platform/ocpp-types-go/gen/ocpp16"
    "github.com/e-flux-platform/ocpp-types-go/gen/ocpp201"
)
```

### Working with message structs

Each versioned package exposes structs for all OCPP message types. OCPP 1.5/1.6 types have no suffix; OCPP 2.0.1/2.1 types use the `Request`/`Response` suffix:

```go
// OCPP 1.6
req := &ocpp16.RemoteStartTransaction{
    ConnectorId: 1,
    IdTag:       "ABC123",
}

// OCPP 2.0.1
req := &ocpp201.RequestStartTransactionRequest{
    RemoteStartId: 42,
    IdToken:       ocpp201.IdTokenType{IdToken: "ABC123", Type: "Central"},
}
```

### Action routing

The root `ocpp` package provides helpers for routing actions to the correct request struct by version:

```go
import "github.com/e-flux-platform/ocpp-types-go"

// Returns an empty struct pointer of the correct versioned type
s, err := ocpp.ActionToRequestStruct(ocpp.ResetAction, ocpp.Version16)
if err != nil {
    // handle ErrUnsupportedVersion or ErrUnknownAction
}

// Unmarshal the raw JSON payload into the struct
if err := json.Unmarshal(payload, s); err != nil {
    // handle error
}
```

### Validating actions

```go
ok := ocpp.IsValidCentralSystemToChargerPointAction(ocpp.ResetAction) // true
```

### Version constants

```go
ocpp.Version15   // "ocpp1.5"
ocpp.Version16   // "ocpp1.6"
ocpp.Version201  // "ocpp2.0.1"
ocpp.Version21   // "ocpp2.1"
```

### Action direction constants

Actions are typed by direction:

```go
// Charging station → Central system
ocpp.AuthorizeAction
ocpp.BootNotificationAction
ocpp.HeartbeatAction
ocpp.MeterValuesAction
ocpp.StartTransactionAction
ocpp.StopTransactionAction
ocpp.StatusNotificationAction
// ...

// Central system → Charging station
ocpp.ResetAction
ocpp.RemoteStartTransactionAction
ocpp.RemoteStopTransactionAction
ocpp.SetChargingProfileAction
// ...
```

## Code generation

The structs in `gen/` are auto-generated from JSON schemas — do not edit them manually.

**Prerequisites:** [Task](https://taskfile.dev) and Go must be installed.

```bash
# Regenerate all types from JSON schemas
task generate
```

Generation is handled by `cmd/generate-ocpp-types`, which is invoked once per protocol version. The tool preprocesses the embedded JSON schemas (deduplicating shared definitions across files), then uses [`go-jsonschema`](https://github.com/atombender/go-jsonschema) to produce the Go structs.

### Custom patches

JSON patch files ([RFC 6902](https://datatracker.ietf.org/doc/html/rfc6902)) can be applied to schemas before generation. Built-in patches live under `cmd/generate-ocpp-types/schemas/<version>/patches/`. To apply additional patches, pass a flat directory of patch files via `--patches-dir`; each file must be named after the schema it patches (e.g. `BootNotification.json`):

```bash
go run ./cmd/generate-ocpp-types \
  --version 1.6 \
  --package github.com/example/myrepo/gen/ocpp16 \
  --output-dir ./gen/ocpp16 \
  --patches-dir ./my-patches
```

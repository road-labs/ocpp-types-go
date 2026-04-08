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

### Operation routing

The root `ocpp` package provides helpers for routing operations to the correct request struct by version:

```go
import "github.com/e-flux-platform/ocpp-types-go"

// Returns an empty struct pointer of the correct versioned type
s, err := ocpp.OperationToRequestStruct(ocpp.ResetOperation, ocpp.Version16)
if err != nil {
    // handle ErrUnsupportedVersion or ErrUnableToValidateOperation
}

// Unmarshal the raw JSON payload into the struct
if err := json.Unmarshal(payload, s); err != nil {
    // handle error
}
```

### Validating operations

```go
ok := ocpp.IsValidCentralSystemToChargerPointOperation(ocpp.ResetOperation) // true
```

### Version constants

```go
ocpp.Version15   // "ocpp1.5"
ocpp.Version16   // "ocpp1.6"
ocpp.Version201  // "ocpp2.0.1"
ocpp.Version21   // "ocpp2.1"
```

### Operation direction constants

Operations are typed by direction:

```go
// Charging station → Central system
ocpp.AuthorizeOperation
ocpp.BootNotificationOperation
ocpp.HeartbeatOperation
ocpp.MeterValuesOperation
ocpp.StartTransactionOperation
ocpp.StopTransactionOperation
ocpp.StatusNotificationOperation
// ...

// Central system → Charging station
ocpp.ResetOperation
ocpp.RemoteStartTransactionOperation
ocpp.RemoteStopTransactionOperation
ocpp.SetChargingProfileOperation
// ...
```

## Code generation

The structs in `gen/` are auto-generated from the JSON schemas in `schemas/` — do not edit them manually.

**Prerequisites:** [Task](https://taskfile.dev) and Go must be installed.

```bash
# Install the code generation tool
task install-tools

# Regenerate all types from JSON schemas
task generate
```

The pipeline copies JSON schemas to a temp directory, runs a preprocessor (`tools/schemas`) to deduplicate shared definitions, then runs [`go-jsonschema`](https://github.com/atombender/go-jsonschema) to produce the Go structs.

JSON patch files under `schemas/<version>/schema/patches/` can be used to customize schemas before generation.
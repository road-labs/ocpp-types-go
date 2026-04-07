package ocpp

// Version describes the version of the OCPP protocol
type Version string

const (
	Version15  Version = "ocpp1.5"
	Version16  Version = "ocpp1.6"
	Version201 Version = "ocpp2.0.1"
	Version21  Version = "ocpp2.1"
)

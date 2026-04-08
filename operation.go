package ocpp

import (
	"errors"

	types15 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp15"
	types16 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp16"
	types201 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp201"
	types21 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp21"
)

var (
	ErrUnableToValidateOperation = errors.New("cannot validate operation")
	ErrUnsupportedVersion        = errors.New("unsupported version")
)

// ChargerPointToCentralSystemOperation is an operation sent from a charging station to a central system.
type ChargerPointToCentralSystemOperation string

func (o ChargerPointToCentralSystemOperation) String() string {
	return string(o)
}

// CentralSystemToChargerPointOperation is a operation sent from a central system to a charging station.
type CentralSystemToChargerPointOperation string

func (o CentralSystemToChargerPointOperation) String() string {
	return string(o)
}

const (
	AuthorizeOperation                        ChargerPointToCentralSystemOperation = "Authorize"
	BootNotificationOperation                 ChargerPointToCentralSystemOperation = "BootNotification"
	DataTransferCpToCsOperation               ChargerPointToCentralSystemOperation = "DataTransfer"
	DiagnosticsStatusNotificationOperation    ChargerPointToCentralSystemOperation = "DiagnosticsStatusNotification"
	FirmwareStatusNotificationOperation       ChargerPointToCentralSystemOperation = "FirmwareStatusNotification"
	SignedFirmwareStatusNotificationOperation ChargerPointToCentralSystemOperation = "SignedFirmwareStatusNotification"
	HeartbeatOperation                        ChargerPointToCentralSystemOperation = "Heartbeat"
	MeterValuesOperation                      ChargerPointToCentralSystemOperation = "MeterValues"
	NotifyReportOperation                     ChargerPointToCentralSystemOperation = "NotifyReport"
	StartTransactionOperation                 ChargerPointToCentralSystemOperation = "StartTransaction"
	StatusNotificationOperation               ChargerPointToCentralSystemOperation = "StatusNotification"
	StopTransactionOperation                  ChargerPointToCentralSystemOperation = "StopTransaction"
	TransactionEventOperation                 ChargerPointToCentralSystemOperation = "TransactionEvent"
	SecurityEventNotificationOperation        ChargerPointToCentralSystemOperation = "SecurityEventNotification"
	NotifyEventOperation                      ChargerPointToCentralSystemOperation = "NotifyEvent"
	NotifyEVChargingNeedsOperation            ChargerPointToCentralSystemOperation = "NotifyEVChargingNeeds"
	LogStatusNotificationOperation            ChargerPointToCentralSystemOperation = "LogStatusNotification"

	CancelReservationOperation          CentralSystemToChargerPointOperation = "CancelReservation"
	ChangeAvailabilityOperation         CentralSystemToChargerPointOperation = "ChangeAvailability"
	ChangeConfigurationOperation        CentralSystemToChargerPointOperation = "ChangeConfiguration"
	ClearCacheOperation                 CentralSystemToChargerPointOperation = "ClearCache"
	ClearChargingProfileOperation       CentralSystemToChargerPointOperation = "ClearChargingProfile"
	DataTransferCsToCpOperation         CentralSystemToChargerPointOperation = "DataTransfer"
	GetCompositeScheduleOperation       CentralSystemToChargerPointOperation = "GetCompositeSchedule"
	GetConfigurationOperation           CentralSystemToChargerPointOperation = "GetConfiguration"
	GetDiagnosticsOperation             CentralSystemToChargerPointOperation = "GetDiagnostics"
	GetLocalListVersionOperation        CentralSystemToChargerPointOperation = "GetLocalListVersion"
	GetLogOperation                     CentralSystemToChargerPointOperation = "GetLog"
	RemoteStartTransactionOperation     CentralSystemToChargerPointOperation = "RemoteStartTransaction"
	RequestStartTransactionOperation    CentralSystemToChargerPointOperation = "RequestStartTransaction"
	RemoteStopTransactionOperation      CentralSystemToChargerPointOperation = "RemoteStopTransaction"
	RequestStopTransactionOperation     CentralSystemToChargerPointOperation = "RequestStopTransaction"
	ReserveNowOperation                 CentralSystemToChargerPointOperation = "ReserveNow"
	ResetOperation                      CentralSystemToChargerPointOperation = "Reset"
	SendLocalListOperation              CentralSystemToChargerPointOperation = "SendLocalList"
	SetChargingProfileOperation         CentralSystemToChargerPointOperation = "SetChargingProfile"
	GetChargingProfilesOperation        CentralSystemToChargerPointOperation = "GetChargingProfiles"
	TriggerMessageOperation             CentralSystemToChargerPointOperation = "TriggerMessage"
	UnlockConnectorOperation            CentralSystemToChargerPointOperation = "UnlockConnector"
	UpdateFirmwareOperation             CentralSystemToChargerPointOperation = "UpdateFirmware"
	SetVariablesOperation               CentralSystemToChargerPointOperation = "SetVariables"
	GetVariablesOperation               CentralSystemToChargerPointOperation = "GetVariables"
	GetBaseReportOperation              CentralSystemToChargerPointOperation = "GetBaseReport"
	GetInstalledCertificateIdsOperation CentralSystemToChargerPointOperation = "GetInstalledCertificateIds"
	SetNetworkProfileOperation          CentralSystemToChargerPointOperation = "SetNetworkProfileOperation"
	CertificateSignedOperation          CentralSystemToChargerPointOperation = "CertificateSigned"
	DeleteCertificateOperation          CentralSystemToChargerPointOperation = "DeleteCertificate"
	ExtendedTriggerMessageOperation     CentralSystemToChargerPointOperation = "ExtendedTriggerMessage"
	InstallCertificateOperation         CentralSystemToChargerPointOperation = "InstallCertificate"
	SignedUpdateFirmwareOperation       CentralSystemToChargerPointOperation = "SignedUpdateFirmware"
	ClearDisplayMessageOperation        CentralSystemToChargerPointOperation = "ClearDisplayMessage"
	ClearVariableMonitoringOperation    CentralSystemToChargerPointOperation = "ClearVariableMonitoring"
	CostUpdatedOperation                CentralSystemToChargerPointOperation = "CostUpdated"
	CustomerInformationOperation        CentralSystemToChargerPointOperation = "CustomerInformation"
	GetDisplayMessagesOperation         CentralSystemToChargerPointOperation = "GetDisplayMessages"
	GetMonitoringReportOperation        CentralSystemToChargerPointOperation = "GetMonitoringReport"
	GetReportOperation                  CentralSystemToChargerPointOperation = "GetReport"
	GetTransactionStatusOperation       CentralSystemToChargerPointOperation = "GetTransactionStatus"
	PublishFirmwareOperation            CentralSystemToChargerPointOperation = "PublishFirmware"
	ReportChargingProfilesOperation     CentralSystemToChargerPointOperation = "ReportChargingProfiles"
	SetDisplayMessageOperation          CentralSystemToChargerPointOperation = "SetDisplayMessage"
	SetMonitoringBaseOperation          CentralSystemToChargerPointOperation = "SetMonitoringBase"
	SetMonitoringLevelOperation         CentralSystemToChargerPointOperation = "SetMonitoringLevel"
	SetVariableMonitoringOperation      CentralSystemToChargerPointOperation = "SetVariableMonitoring"
	UnpublishFirmwareOperation          CentralSystemToChargerPointOperation = "UnpublishFirmware"
)

func OperationToRequestStruct(operation CentralSystemToChargerPointOperation, version Version) (any, error) {
	switch version {
	case Version15:
		return operationToRequestStruct15(operation)
	case Version16:
		return operationToRequestStruct16(operation)
	case Version201:
		return operationToRequestStruct201(operation)
	case Version21:
		return operationToRequestStruct21(operation)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func operationToRequestStruct15(operation CentralSystemToChargerPointOperation) (any, error) {
	switch operation {
	case CancelReservationOperation:
		return &types15.CancelReservation{}, nil
	case ChangeAvailabilityOperation:
		return &types15.ChangeAvailability{}, nil
	case ChangeConfigurationOperation:
		return &types15.ChangeConfiguration{}, nil
	case ClearCacheOperation:
		return &types15.ClearCache{}, nil
	case DataTransferCsToCpOperation:
		return &types15.DataTransfer{}, nil
	case GetConfigurationOperation:
		return &types15.GetConfiguration{}, nil
	case GetDiagnosticsOperation:
		return &types15.GetDiagnostics{}, nil
	case GetLocalListVersionOperation:
		return &types15.GetLocalListVersion{}, nil
	case RemoteStartTransactionOperation:
		return &types15.RemoteStartTransaction{}, nil
	case RemoteStopTransactionOperation:
		return &types15.RemoteStopTransaction{}, nil
	case ReserveNowOperation:
		return &types15.ReserveNow{}, nil
	case ResetOperation:
		return &types15.Reset{}, nil
	case SendLocalListOperation:
		return &types15.SendLocalList{}, nil
	case UnlockConnectorOperation:
		return &types15.UnlockConnector{}, nil
	case UpdateFirmwareOperation:
		return &types15.UpdateFirmware{}, nil
	default:
		return nil, ErrUnableToValidateOperation
	}
}

func operationToRequestStruct16(operation CentralSystemToChargerPointOperation) (any, error) {
	switch operation {
	case CancelReservationOperation:
		return &types16.CancelReservation{}, nil
	case CertificateSignedOperation:
		return &types16.CertificateSigned{}, nil
	case ChangeAvailabilityOperation:
		return &types16.ChangeAvailability{}, nil
	case ChangeConfigurationOperation:
		return &types16.ChangeConfiguration{}, nil
	case ClearCacheOperation:
		return &types16.ClearCache{}, nil
	case ClearChargingProfileOperation:
		return &types16.ClearChargingProfile{}, nil
	case DataTransferCsToCpOperation:
		return &types16.DataTransfer{}, nil
	case DeleteCertificateOperation:
		return &types16.DeleteCertificate{}, nil
	case ExtendedTriggerMessageOperation:
		return &types16.ExtendedTriggerMessage{}, nil
	case GetCompositeScheduleOperation:
		return &types16.GetCompositeSchedule{}, nil
	case GetConfigurationOperation:
		return &types16.GetConfiguration{}, nil
	case GetDiagnosticsOperation:
		return &types16.GetDiagnostics{}, nil
	case GetInstalledCertificateIdsOperation:
		return &types16.GetInstalledCertificateIds{}, nil
	case GetLocalListVersionOperation:
		return &types16.GetLocalListVersion{}, nil
	case GetLogOperation:
		return &types16.GetLog{}, nil
	case InstallCertificateOperation:
		return &types16.InstallCertificate{}, nil
	case RemoteStartTransactionOperation:
		return &types16.RemoteStartTransaction{}, nil
	case RemoteStopTransactionOperation:
		return &types16.RemoteStopTransaction{}, nil
	case ReserveNowOperation:
		return &types16.ReserveNow{}, nil
	case ResetOperation:
		return &types16.Reset{}, nil
	case SendLocalListOperation:
		return &types16.SendLocalList{}, nil
	case SetChargingProfileOperation:
		return &types16.SetChargingProfile{}, nil
	case SignedUpdateFirmwareOperation:
		return &types16.SignedUpdateFirmware{}, nil
	case TriggerMessageOperation:
		return &types16.TriggerMessage{}, nil
	case UnlockConnectorOperation:
		return &types16.UnlockConnector{}, nil
	case UpdateFirmwareOperation:
		return &types16.UpdateFirmware{}, nil
	default:
		return nil, ErrUnableToValidateOperation
	}
}

func operationToRequestStruct201(operation CentralSystemToChargerPointOperation) (any, error) {
	switch operation {
	case CancelReservationOperation:
		return &types201.CancelReservationRequest{}, nil
	case CertificateSignedOperation:
		return &types201.CertificateSignedRequest{}, nil
	case ChangeAvailabilityOperation:
		return &types201.ChangeAvailabilityRequest{}, nil
	case ClearCacheOperation:
		return &types201.ClearCacheRequest{}, nil
	case ClearChargingProfileOperation:
		return &types201.ClearChargingProfileRequest{}, nil
	case ClearDisplayMessageOperation:
		return &types201.ClearDisplayMessageRequest{}, nil
	case ClearVariableMonitoringOperation:
		return &types201.ClearVariableMonitoringRequest{}, nil
	case CostUpdatedOperation:
		return &types201.CostUpdatedRequest{}, nil
	case CustomerInformationOperation:
		return &types201.CustomerInformationRequest{}, nil
	case DataTransferCsToCpOperation:
		return &types201.DataTransferRequest{}, nil
	case DeleteCertificateOperation:
		return &types201.DeleteCertificateRequest{}, nil
	case GetBaseReportOperation:
		return &types201.GetBaseReportRequest{}, nil
	case GetChargingProfilesOperation:
		return &types201.GetChargingProfilesRequest{}, nil
	case GetCompositeScheduleOperation:
		return &types201.GetCompositeScheduleRequest{}, nil
	case GetDisplayMessagesOperation:
		return &types201.GetDisplayMessagesRequest{}, nil
	case GetInstalledCertificateIdsOperation:
		return &types201.GetInstalledCertificateIdsRequest{}, nil
	case GetLocalListVersionOperation:
		return &types201.GetLocalListVersionRequest{}, nil
	case GetLogOperation:
		return &types201.GetLogRequest{}, nil
	case GetMonitoringReportOperation:
		return &types201.GetMonitoringReportRequest{}, nil
	case GetReportOperation:
		return &types201.GetReportRequest{}, nil
	case GetTransactionStatusOperation:
		return &types201.GetTransactionStatusRequest{}, nil
	case GetVariablesOperation:
		return &types201.GetVariablesRequest{}, nil
	case InstallCertificateOperation:
		return &types201.InstallCertificateRequest{}, nil
	case PublishFirmwareOperation:
		return &types201.PublishFirmwareRequest{}, nil
	case ReportChargingProfilesOperation:
		return &types201.ReportChargingProfilesRequest{}, nil
	case RequestStartTransactionOperation:
		return &types201.RequestStartTransactionRequest{}, nil
	case RequestStopTransactionOperation:
		return &types201.RequestStopTransactionRequest{}, nil
	case ReserveNowOperation:
		return &types201.ReserveNowRequest{}, nil
	case ResetOperation:
		return &types201.ResetRequest{}, nil
	case SendLocalListOperation:
		return &types201.SendLocalListRequest{}, nil
	case SetChargingProfileOperation:
		return &types201.SetChargingProfileRequest{}, nil
	case SetDisplayMessageOperation:
		return &types201.SetDisplayMessageRequest{}, nil
	case SetMonitoringBaseOperation:
		return &types201.SetMonitoringBaseRequest{}, nil
	case SetMonitoringLevelOperation:
		return &types201.SetMonitoringLevelRequest{}, nil
	case SetNetworkProfileOperation:
		return &types201.SetNetworkProfileRequest{}, nil
	case SetVariableMonitoringOperation:
		return &types201.SetVariableMonitoringRequest{}, nil
	case SetVariablesOperation:
		return &types201.SetVariablesRequest{}, nil
	case TriggerMessageOperation:
		return &types201.TriggerMessageRequest{}, nil
	case UnlockConnectorOperation:
		return &types201.UnlockConnectorRequest{}, nil
	case UnpublishFirmwareOperation:
		return &types201.UnpublishFirmwareRequest{}, nil
	case UpdateFirmwareOperation:
		return &types201.UpdateFirmwareRequest{}, nil
	default:
		return nil, ErrUnableToValidateOperation
	}
}

func operationToRequestStruct21(operation CentralSystemToChargerPointOperation) (any, error) {
	switch operation {
	case CancelReservationOperation:
		return &types21.CancelReservationRequest{}, nil
	case CertificateSignedOperation:
		return &types21.CertificateSignedRequest{}, nil
	case ChangeAvailabilityOperation:
		return &types21.ChangeAvailabilityRequest{}, nil
	case ClearCacheOperation:
		return &types21.ClearCacheRequest{}, nil
	case ClearChargingProfileOperation:
		return &types21.ClearChargingProfileRequest{}, nil
	case ClearDisplayMessageOperation:
		return &types21.ClearDisplayMessageRequest{}, nil
	case ClearVariableMonitoringOperation:
		return &types21.ClearVariableMonitoringRequest{}, nil
	case CostUpdatedOperation:
		return &types21.CostUpdatedRequest{}, nil
	case CustomerInformationOperation:
		return &types21.CustomerInformationRequest{}, nil
	case DataTransferCsToCpOperation:
		return &types21.DataTransferRequest{}, nil
	case DeleteCertificateOperation:
		return &types21.DeleteCertificateRequest{}, nil
	case GetBaseReportOperation:
		return &types21.GetBaseReportRequest{}, nil
	case GetChargingProfilesOperation:
		return &types21.GetChargingProfilesRequest{}, nil
	case GetCompositeScheduleOperation:
		return &types21.GetCompositeScheduleRequest{}, nil
	case GetDisplayMessagesOperation:
		return &types21.GetDisplayMessagesRequest{}, nil
	case GetInstalledCertificateIdsOperation:
		return &types21.GetInstalledCertificateIdsRequest{}, nil
	case GetLocalListVersionOperation:
		return &types21.GetLocalListVersionRequest{}, nil
	case GetLogOperation:
		return &types21.GetLogRequest{}, nil
	case GetMonitoringReportOperation:
		return &types21.GetMonitoringReportRequest{}, nil
	case GetReportOperation:
		return &types21.GetReportRequest{}, nil
	case GetTransactionStatusOperation:
		return &types21.GetTransactionStatusRequest{}, nil
	case GetVariablesOperation:
		return &types21.GetVariablesRequest{}, nil
	case InstallCertificateOperation:
		return &types21.InstallCertificateRequest{}, nil
	case PublishFirmwareOperation:
		return &types21.PublishFirmwareRequest{}, nil
	case ReportChargingProfilesOperation:
		return &types21.ReportChargingProfilesRequest{}, nil
	case RequestStartTransactionOperation:
		return &types21.RequestStartTransactionRequest{}, nil
	case RequestStopTransactionOperation:
		return &types21.RequestStopTransactionRequest{}, nil
	case ReserveNowOperation:
		return &types21.ReserveNowRequest{}, nil
	case ResetOperation:
		return &types21.ResetRequest{}, nil
	case SendLocalListOperation:
		return &types21.SendLocalListRequest{}, nil
	case SetChargingProfileOperation:
		return &types21.SetChargingProfileRequest{}, nil
	case SetDisplayMessageOperation:
		return &types21.SetDisplayMessageRequest{}, nil
	case SetMonitoringBaseOperation:
		return &types21.SetMonitoringBaseRequest{}, nil
	case SetMonitoringLevelOperation:
		return &types21.SetMonitoringLevelRequest{}, nil
	case SetNetworkProfileOperation:
		return &types21.SetNetworkProfileRequest{}, nil
	case SetVariableMonitoringOperation:
		return &types21.SetVariableMonitoringRequest{}, nil
	case SetVariablesOperation:
		return &types21.SetVariablesRequest{}, nil
	case TriggerMessageOperation:
		return &types21.TriggerMessageRequest{}, nil
	case UnlockConnectorOperation:
		return &types21.UnlockConnectorRequest{}, nil
	case UnpublishFirmwareOperation:
		return &types21.UnpublishFirmwareRequest{}, nil
	case UpdateFirmwareOperation:
		return &types21.UpdateFirmwareRequest{}, nil
	default:
		return nil, ErrUnableToValidateOperation
	}
}

func IsValidCentralSystemToChargerPointOperation(operation CentralSystemToChargerPointOperation) bool {
	switch operation {
	case CancelReservationOperation,
		ChangeAvailabilityOperation,
		ChangeConfigurationOperation,
		ClearCacheOperation,
		ClearChargingProfileOperation,
		DataTransferCsToCpOperation,
		GetCompositeScheduleOperation,
		GetConfigurationOperation,
		GetDiagnosticsOperation,
		GetLocalListVersionOperation,
		GetLogOperation,
		RemoteStartTransactionOperation,
		RequestStartTransactionOperation,
		RemoteStopTransactionOperation,
		RequestStopTransactionOperation,
		ReserveNowOperation,
		ResetOperation,
		SendLocalListOperation,
		SetChargingProfileOperation,
		GetChargingProfilesOperation,
		TriggerMessageOperation,
		UnlockConnectorOperation,
		UpdateFirmwareOperation,
		SetVariablesOperation,
		GetVariablesOperation,
		GetBaseReportOperation,
		GetInstalledCertificateIdsOperation,
		SetNetworkProfileOperation,
		CertificateSignedOperation,
		DeleteCertificateOperation,
		ExtendedTriggerMessageOperation,
		InstallCertificateOperation,
		SignedUpdateFirmwareOperation,
		ClearDisplayMessageOperation,
		ClearVariableMonitoringOperation,
		CostUpdatedOperation,
		CustomerInformationOperation,
		GetDisplayMessagesOperation,
		GetMonitoringReportOperation,
		GetReportOperation,
		GetTransactionStatusOperation,
		PublishFirmwareOperation,
		ReportChargingProfilesOperation,
		SetDisplayMessageOperation,
		SetMonitoringBaseOperation,
		SetMonitoringLevelOperation,
		SetVariableMonitoringOperation,
		UnpublishFirmwareOperation:
		return true
	default:
		return false
	}
}

package ocpp

import (
	"errors"

	types15 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp15"
	types16 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp16"
	types201 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp201"
	types21 "github.com/e-flux-platform/ocpp-types-go/gen/ocpp21"
)

var (
	ErrUnknownAction      = errors.New("unknown ocpp action")
	ErrUnsupportedVersion = errors.New("unsupported version")
)

// ChargingStationToCentralSystemAction is an action sent from a charging station to a central system.
type ChargingStationToCentralSystemAction string

func (o ChargingStationToCentralSystemAction) String() string {
	return string(o)
}

// CentralSystemToChargingStationAction is an action sent from a central system to a charging station.
type CentralSystemToChargingStationAction string

func (o CentralSystemToChargingStationAction) String() string {
	return string(o)
}

const (
	AuthorizeAction                        ChargingStationToCentralSystemAction = "Authorize"
	BootNotificationAction                 ChargingStationToCentralSystemAction = "BootNotification"
	DataTransferCpToCsAction               ChargingStationToCentralSystemAction = "DataTransfer"
	DiagnosticsStatusNotificationAction    ChargingStationToCentralSystemAction = "DiagnosticsStatusNotification"
	FirmwareStatusNotificationAction       ChargingStationToCentralSystemAction = "FirmwareStatusNotification"
	SignedFirmwareStatusNotificationAction ChargingStationToCentralSystemAction = "SignedFirmwareStatusNotification"
	HeartbeatAction                        ChargingStationToCentralSystemAction = "Heartbeat"
	MeterValuesAction                      ChargingStationToCentralSystemAction = "MeterValues"
	NotifyReportAction                     ChargingStationToCentralSystemAction = "NotifyReport"
	StartTransactionAction                 ChargingStationToCentralSystemAction = "StartTransaction"
	StatusNotificationAction               ChargingStationToCentralSystemAction = "StatusNotification"
	StopTransactionAction                  ChargingStationToCentralSystemAction = "StopTransaction"
	TransactionEventAction                 ChargingStationToCentralSystemAction = "TransactionEvent"
	SecurityEventNotificationAction        ChargingStationToCentralSystemAction = "SecurityEventNotification"
	NotifyEventAction                      ChargingStationToCentralSystemAction = "NotifyEvent"
	NotifyEVChargingNeedsAction            ChargingStationToCentralSystemAction = "NotifyEVChargingNeeds"
	LogStatusNotificationAction            ChargingStationToCentralSystemAction = "LogStatusNotification"

	CancelReservationAction          CentralSystemToChargingStationAction = "CancelReservation"
	ChangeAvailabilityAction         CentralSystemToChargingStationAction = "ChangeAvailability"
	ChangeConfigurationAction        CentralSystemToChargingStationAction = "ChangeConfiguration"
	ClearCacheAction                 CentralSystemToChargingStationAction = "ClearCache"
	ClearChargingProfileAction       CentralSystemToChargingStationAction = "ClearChargingProfile"
	DataTransferCsToCpAction         CentralSystemToChargingStationAction = "DataTransfer"
	GetCompositeScheduleAction       CentralSystemToChargingStationAction = "GetCompositeSchedule"
	GetConfigurationAction           CentralSystemToChargingStationAction = "GetConfiguration"
	GetDiagnosticsAction             CentralSystemToChargingStationAction = "GetDiagnostics"
	GetLocalListVersionAction        CentralSystemToChargingStationAction = "GetLocalListVersion"
	GetLogAction                     CentralSystemToChargingStationAction = "GetLog"
	RemoteStartTransactionAction     CentralSystemToChargingStationAction = "RemoteStartTransaction"
	RequestStartTransactionAction    CentralSystemToChargingStationAction = "RequestStartTransaction"
	RemoteStopTransactionAction      CentralSystemToChargingStationAction = "RemoteStopTransaction"
	RequestStopTransactionAction     CentralSystemToChargingStationAction = "RequestStopTransaction"
	ReserveNowAction                 CentralSystemToChargingStationAction = "ReserveNow"
	ResetAction                      CentralSystemToChargingStationAction = "Reset"
	SendLocalListAction              CentralSystemToChargingStationAction = "SendLocalList"
	SetChargingProfileAction         CentralSystemToChargingStationAction = "SetChargingProfile"
	GetChargingProfilesAction        CentralSystemToChargingStationAction = "GetChargingProfiles"
	TriggerMessageAction             CentralSystemToChargingStationAction = "TriggerMessage"
	UnlockConnectorAction            CentralSystemToChargingStationAction = "UnlockConnector"
	UpdateFirmwareAction             CentralSystemToChargingStationAction = "UpdateFirmware"
	SetVariablesAction               CentralSystemToChargingStationAction = "SetVariables"
	GetVariablesAction               CentralSystemToChargingStationAction = "GetVariables"
	GetBaseReportAction              CentralSystemToChargingStationAction = "GetBaseReport"
	GetInstalledCertificateIdsAction CentralSystemToChargingStationAction = "GetInstalledCertificateIds"
	SetNetworkProfileAction          CentralSystemToChargingStationAction = "SetNetworkProfile"
	CertificateSignedAction          CentralSystemToChargingStationAction = "CertificateSigned"
	DeleteCertificateAction          CentralSystemToChargingStationAction = "DeleteCertificate"
	ExtendedTriggerMessageAction     CentralSystemToChargingStationAction = "ExtendedTriggerMessage"
	InstallCertificateAction         CentralSystemToChargingStationAction = "InstallCertificate"
	SignedUpdateFirmwareAction       CentralSystemToChargingStationAction = "SignedUpdateFirmware"
	ClearDisplayMessageAction        CentralSystemToChargingStationAction = "ClearDisplayMessage"
	ClearVariableMonitoringAction    CentralSystemToChargingStationAction = "ClearVariableMonitoring"
	CostUpdatedAction                CentralSystemToChargingStationAction = "CostUpdated"
	CustomerInformationAction        CentralSystemToChargingStationAction = "CustomerInformation"
	GetDisplayMessagesAction         CentralSystemToChargingStationAction = "GetDisplayMessages"
	GetMonitoringReportAction        CentralSystemToChargingStationAction = "GetMonitoringReport"
	GetReportAction                  CentralSystemToChargingStationAction = "GetReport"
	GetTransactionStatusAction       CentralSystemToChargingStationAction = "GetTransactionStatus"
	PublishFirmwareAction            CentralSystemToChargingStationAction = "PublishFirmware"
	ReportChargingProfilesAction     CentralSystemToChargingStationAction = "ReportChargingProfiles"
	SetDisplayMessageAction          CentralSystemToChargingStationAction = "SetDisplayMessage"
	SetMonitoringBaseAction          CentralSystemToChargingStationAction = "SetMonitoringBase"
	SetMonitoringLevelAction         CentralSystemToChargingStationAction = "SetMonitoringLevel"
	SetVariableMonitoringAction      CentralSystemToChargingStationAction = "SetVariableMonitoring"
	UnpublishFirmwareAction          CentralSystemToChargingStationAction = "UnpublishFirmware"
)

func ActionToRequestStruct(action CentralSystemToChargingStationAction, version Version) (any, error) {
	switch version {
	case Version15:
		return actionToRequestStruct15(action)
	case Version16:
		return actionToRequestStruct16(action)
	case Version201:
		return actionToRequestStruct201(action)
	case Version21:
		return actionToRequestStruct21(action)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func actionToRequestStruct15(action CentralSystemToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types15.CancelReservation{}, nil
	case ChangeAvailabilityAction:
		return &types15.ChangeAvailability{}, nil
	case ChangeConfigurationAction:
		return &types15.ChangeConfiguration{}, nil
	case ClearCacheAction:
		return &types15.ClearCache{}, nil
	case DataTransferCsToCpAction:
		return &types15.DataTransfer{}, nil
	case GetConfigurationAction:
		return &types15.GetConfiguration{}, nil
	case GetDiagnosticsAction:
		return &types15.GetDiagnostics{}, nil
	case GetLocalListVersionAction:
		return &types15.GetLocalListVersion{}, nil
	case RemoteStartTransactionAction:
		return &types15.RemoteStartTransaction{}, nil
	case RemoteStopTransactionAction:
		return &types15.RemoteStopTransaction{}, nil
	case ReserveNowAction:
		return &types15.ReserveNow{}, nil
	case ResetAction:
		return &types15.Reset{}, nil
	case SendLocalListAction:
		return &types15.SendLocalList{}, nil
	case UnlockConnectorAction:
		return &types15.UnlockConnector{}, nil
	case UpdateFirmwareAction:
		return &types15.UpdateFirmware{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func actionToRequestStruct16(action CentralSystemToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types16.CancelReservation{}, nil
	case CertificateSignedAction:
		return &types16.CertificateSigned{}, nil
	case ChangeAvailabilityAction:
		return &types16.ChangeAvailability{}, nil
	case ChangeConfigurationAction:
		return &types16.ChangeConfiguration{}, nil
	case ClearCacheAction:
		return &types16.ClearCache{}, nil
	case ClearChargingProfileAction:
		return &types16.ClearChargingProfile{}, nil
	case DataTransferCsToCpAction:
		return &types16.DataTransfer{}, nil
	case DeleteCertificateAction:
		return &types16.DeleteCertificate{}, nil
	case ExtendedTriggerMessageAction:
		return &types16.ExtendedTriggerMessage{}, nil
	case GetCompositeScheduleAction:
		return &types16.GetCompositeSchedule{}, nil
	case GetConfigurationAction:
		return &types16.GetConfiguration{}, nil
	case GetDiagnosticsAction:
		return &types16.GetDiagnostics{}, nil
	case GetInstalledCertificateIdsAction:
		return &types16.GetInstalledCertificateIds{}, nil
	case GetLocalListVersionAction:
		return &types16.GetLocalListVersion{}, nil
	case GetLogAction:
		return &types16.GetLog{}, nil
	case InstallCertificateAction:
		return &types16.InstallCertificate{}, nil
	case RemoteStartTransactionAction:
		return &types16.RemoteStartTransaction{}, nil
	case RemoteStopTransactionAction:
		return &types16.RemoteStopTransaction{}, nil
	case ReserveNowAction:
		return &types16.ReserveNow{}, nil
	case ResetAction:
		return &types16.Reset{}, nil
	case SendLocalListAction:
		return &types16.SendLocalList{}, nil
	case SetChargingProfileAction:
		return &types16.SetChargingProfile{}, nil
	case SignedUpdateFirmwareAction:
		return &types16.SignedUpdateFirmware{}, nil
	case TriggerMessageAction:
		return &types16.TriggerMessage{}, nil
	case UnlockConnectorAction:
		return &types16.UnlockConnector{}, nil
	case UpdateFirmwareAction:
		return &types16.UpdateFirmware{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func actionToRequestStruct201(action CentralSystemToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types201.CancelReservationRequest{}, nil
	case CertificateSignedAction:
		return &types201.CertificateSignedRequest{}, nil
	case ChangeAvailabilityAction:
		return &types201.ChangeAvailabilityRequest{}, nil
	case ClearCacheAction:
		return &types201.ClearCacheRequest{}, nil
	case ClearChargingProfileAction:
		return &types201.ClearChargingProfileRequest{}, nil
	case ClearDisplayMessageAction:
		return &types201.ClearDisplayMessageRequest{}, nil
	case ClearVariableMonitoringAction:
		return &types201.ClearVariableMonitoringRequest{}, nil
	case CostUpdatedAction:
		return &types201.CostUpdatedRequest{}, nil
	case CustomerInformationAction:
		return &types201.CustomerInformationRequest{}, nil
	case DataTransferCsToCpAction:
		return &types201.DataTransferRequest{}, nil
	case DeleteCertificateAction:
		return &types201.DeleteCertificateRequest{}, nil
	case GetBaseReportAction:
		return &types201.GetBaseReportRequest{}, nil
	case GetChargingProfilesAction:
		return &types201.GetChargingProfilesRequest{}, nil
	case GetCompositeScheduleAction:
		return &types201.GetCompositeScheduleRequest{}, nil
	case GetDisplayMessagesAction:
		return &types201.GetDisplayMessagesRequest{}, nil
	case GetInstalledCertificateIdsAction:
		return &types201.GetInstalledCertificateIdsRequest{}, nil
	case GetLocalListVersionAction:
		return &types201.GetLocalListVersionRequest{}, nil
	case GetLogAction:
		return &types201.GetLogRequest{}, nil
	case GetMonitoringReportAction:
		return &types201.GetMonitoringReportRequest{}, nil
	case GetReportAction:
		return &types201.GetReportRequest{}, nil
	case GetTransactionStatusAction:
		return &types201.GetTransactionStatusRequest{}, nil
	case GetVariablesAction:
		return &types201.GetVariablesRequest{}, nil
	case InstallCertificateAction:
		return &types201.InstallCertificateRequest{}, nil
	case PublishFirmwareAction:
		return &types201.PublishFirmwareRequest{}, nil
	case ReportChargingProfilesAction:
		return &types201.ReportChargingProfilesRequest{}, nil
	case RequestStartTransactionAction:
		return &types201.RequestStartTransactionRequest{}, nil
	case RequestStopTransactionAction:
		return &types201.RequestStopTransactionRequest{}, nil
	case ReserveNowAction:
		return &types201.ReserveNowRequest{}, nil
	case ResetAction:
		return &types201.ResetRequest{}, nil
	case SendLocalListAction:
		return &types201.SendLocalListRequest{}, nil
	case SetChargingProfileAction:
		return &types201.SetChargingProfileRequest{}, nil
	case SetDisplayMessageAction:
		return &types201.SetDisplayMessageRequest{}, nil
	case SetMonitoringBaseAction:
		return &types201.SetMonitoringBaseRequest{}, nil
	case SetMonitoringLevelAction:
		return &types201.SetMonitoringLevelRequest{}, nil
	case SetNetworkProfileAction:
		return &types201.SetNetworkProfileRequest{}, nil
	case SetVariableMonitoringAction:
		return &types201.SetVariableMonitoringRequest{}, nil
	case SetVariablesAction:
		return &types201.SetVariablesRequest{}, nil
	case TriggerMessageAction:
		return &types201.TriggerMessageRequest{}, nil
	case UnlockConnectorAction:
		return &types201.UnlockConnectorRequest{}, nil
	case UnpublishFirmwareAction:
		return &types201.UnpublishFirmwareRequest{}, nil
	case UpdateFirmwareAction:
		return &types201.UpdateFirmwareRequest{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func actionToRequestStruct21(action CentralSystemToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types21.CancelReservationRequest{}, nil
	case CertificateSignedAction:
		return &types21.CertificateSignedRequest{}, nil
	case ChangeAvailabilityAction:
		return &types21.ChangeAvailabilityRequest{}, nil
	case ClearCacheAction:
		return &types21.ClearCacheRequest{}, nil
	case ClearChargingProfileAction:
		return &types21.ClearChargingProfileRequest{}, nil
	case ClearDisplayMessageAction:
		return &types21.ClearDisplayMessageRequest{}, nil
	case ClearVariableMonitoringAction:
		return &types21.ClearVariableMonitoringRequest{}, nil
	case CostUpdatedAction:
		return &types21.CostUpdatedRequest{}, nil
	case CustomerInformationAction:
		return &types21.CustomerInformationRequest{}, nil
	case DataTransferCsToCpAction:
		return &types21.DataTransferRequest{}, nil
	case DeleteCertificateAction:
		return &types21.DeleteCertificateRequest{}, nil
	case GetBaseReportAction:
		return &types21.GetBaseReportRequest{}, nil
	case GetChargingProfilesAction:
		return &types21.GetChargingProfilesRequest{}, nil
	case GetCompositeScheduleAction:
		return &types21.GetCompositeScheduleRequest{}, nil
	case GetDisplayMessagesAction:
		return &types21.GetDisplayMessagesRequest{}, nil
	case GetInstalledCertificateIdsAction:
		return &types21.GetInstalledCertificateIdsRequest{}, nil
	case GetLocalListVersionAction:
		return &types21.GetLocalListVersionRequest{}, nil
	case GetLogAction:
		return &types21.GetLogRequest{}, nil
	case GetMonitoringReportAction:
		return &types21.GetMonitoringReportRequest{}, nil
	case GetReportAction:
		return &types21.GetReportRequest{}, nil
	case GetTransactionStatusAction:
		return &types21.GetTransactionStatusRequest{}, nil
	case GetVariablesAction:
		return &types21.GetVariablesRequest{}, nil
	case InstallCertificateAction:
		return &types21.InstallCertificateRequest{}, nil
	case PublishFirmwareAction:
		return &types21.PublishFirmwareRequest{}, nil
	case ReportChargingProfilesAction:
		return &types21.ReportChargingProfilesRequest{}, nil
	case RequestStartTransactionAction:
		return &types21.RequestStartTransactionRequest{}, nil
	case RequestStopTransactionAction:
		return &types21.RequestStopTransactionRequest{}, nil
	case ReserveNowAction:
		return &types21.ReserveNowRequest{}, nil
	case ResetAction:
		return &types21.ResetRequest{}, nil
	case SendLocalListAction:
		return &types21.SendLocalListRequest{}, nil
	case SetChargingProfileAction:
		return &types21.SetChargingProfileRequest{}, nil
	case SetDisplayMessageAction:
		return &types21.SetDisplayMessageRequest{}, nil
	case SetMonitoringBaseAction:
		return &types21.SetMonitoringBaseRequest{}, nil
	case SetMonitoringLevelAction:
		return &types21.SetMonitoringLevelRequest{}, nil
	case SetNetworkProfileAction:
		return &types21.SetNetworkProfileRequest{}, nil
	case SetVariableMonitoringAction:
		return &types21.SetVariableMonitoringRequest{}, nil
	case SetVariablesAction:
		return &types21.SetVariablesRequest{}, nil
	case TriggerMessageAction:
		return &types21.TriggerMessageRequest{}, nil
	case UnlockConnectorAction:
		return &types21.UnlockConnectorRequest{}, nil
	case UnpublishFirmwareAction:
		return &types21.UnpublishFirmwareRequest{}, nil
	case UpdateFirmwareAction:
		return &types21.UpdateFirmwareRequest{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func IsValidCentralSystemToChargingStationAction(action CentralSystemToChargingStationAction) bool {
	switch action {
	case CancelReservationAction,
		ChangeAvailabilityAction,
		ChangeConfigurationAction,
		ClearCacheAction,
		ClearChargingProfileAction,
		DataTransferCsToCpAction,
		GetCompositeScheduleAction,
		GetConfigurationAction,
		GetDiagnosticsAction,
		GetLocalListVersionAction,
		GetLogAction,
		RemoteStartTransactionAction,
		RequestStartTransactionAction,
		RemoteStopTransactionAction,
		RequestStopTransactionAction,
		ReserveNowAction,
		ResetAction,
		SendLocalListAction,
		SetChargingProfileAction,
		GetChargingProfilesAction,
		TriggerMessageAction,
		UnlockConnectorAction,
		UpdateFirmwareAction,
		SetVariablesAction,
		GetVariablesAction,
		GetBaseReportAction,
		GetInstalledCertificateIdsAction,
		SetNetworkProfileAction,
		CertificateSignedAction,
		DeleteCertificateAction,
		ExtendedTriggerMessageAction,
		InstallCertificateAction,
		SignedUpdateFirmwareAction,
		ClearDisplayMessageAction,
		ClearVariableMonitoringAction,
		CostUpdatedAction,
		CustomerInformationAction,
		GetDisplayMessagesAction,
		GetMonitoringReportAction,
		GetReportAction,
		GetTransactionStatusAction,
		PublishFirmwareAction,
		ReportChargingProfilesAction,
		SetDisplayMessageAction,
		SetMonitoringBaseAction,
		SetMonitoringLevelAction,
		SetVariableMonitoringAction,
		UnpublishFirmwareAction:
		return true
	default:
		return false
	}
}

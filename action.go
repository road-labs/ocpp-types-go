package ocpp

import "errors"

var (
	ErrUnknownAction      = errors.New("unknown ocpp action")
	ErrUnsupportedVersion = errors.New("unsupported version")
)

// ChargingStationToCSMSAction is an action sent from a charging station to a CSMS.
type ChargingStationToCSMSAction string

func (o ChargingStationToCSMSAction) String() string {
	return string(o)
}

// CSMSToChargingStationAction is an action sent from a CSMS to a charging station.
type CSMSToChargingStationAction string

func (o CSMSToChargingStationAction) String() string {
	return string(o)
}

const (
	AuthorizeAction                         ChargingStationToCSMSAction = "Authorize"
	BootNotificationAction                  ChargingStationToCSMSAction = "BootNotification"
	DataTransferCSToCSMSAction              ChargingStationToCSMSAction = "DataTransfer"
	DiagnosticsStatusNotificationAction     ChargingStationToCSMSAction = "DiagnosticsStatusNotification"
	FirmwareStatusNotificationAction        ChargingStationToCSMSAction = "FirmwareStatusNotification"
	SignedFirmwareStatusNotificationAction  ChargingStationToCSMSAction = "SignedFirmwareStatusNotification"
	HeartbeatAction                         ChargingStationToCSMSAction = "Heartbeat"
	MeterValuesAction                       ChargingStationToCSMSAction = "MeterValues"
	NotifyReportAction                      ChargingStationToCSMSAction = "NotifyReport"
	StartTransactionAction                  ChargingStationToCSMSAction = "StartTransaction"
	StatusNotificationAction                ChargingStationToCSMSAction = "StatusNotification"
	StopTransactionAction                   ChargingStationToCSMSAction = "StopTransaction"
	TransactionEventAction                  ChargingStationToCSMSAction = "TransactionEvent"
	SecurityEventNotificationAction         ChargingStationToCSMSAction = "SecurityEventNotification"
	NotifyEventAction                       ChargingStationToCSMSAction = "NotifyEvent"
	NotifyEVChargingNeedsAction             ChargingStationToCSMSAction = "NotifyEVChargingNeeds"
	LogStatusNotificationAction             ChargingStationToCSMSAction = "LogStatusNotification"
	SignCertificateAction                   ChargingStationToCSMSAction = "SignCertificate"
	Get15118EVCertificateAction             ChargingStationToCSMSAction = "Get15118EVCertificate"
	GetCertificateStatusAction              ChargingStationToCSMSAction = "GetCertificateStatus"
	ClearedChargingLimitAction              ChargingStationToCSMSAction = "ClearedChargingLimit"
	NotifyChargingLimitAction               ChargingStationToCSMSAction = "NotifyChargingLimit"
	NotifyCustomerInformationAction         ChargingStationToCSMSAction = "NotifyCustomerInformation"
	NotifyDisplayMessagesAction             ChargingStationToCSMSAction = "NotifyDisplayMessages"
	NotifyEVChargingScheduleAction          ChargingStationToCSMSAction = "NotifyEVChargingSchedule"
	NotifyMonitoringReportAction            ChargingStationToCSMSAction = "NotifyMonitoringReport"
	PublishFirmwareStatusNotificationAction ChargingStationToCSMSAction = "PublishFirmwareStatusNotification"
	ReportChargingProfilesAction            ChargingStationToCSMSAction = "ReportChargingProfiles"
	ReservationStatusUpdateAction           ChargingStationToCSMSAction = "ReservationStatusUpdate"
	BatterySwapAction                       ChargingStationToCSMSAction = "BatterySwap"
	GetCertificateChainStatusAction         ChargingStationToCSMSAction = "GetCertificateChainStatus"
	NotifyDERAlarmAction                    ChargingStationToCSMSAction = "NotifyDERAlarm"
	NotifyDERStartStopAction                ChargingStationToCSMSAction = "NotifyDERStartStop"
	NotifyPeriodicEventStreamAction         ChargingStationToCSMSAction = "NotifyPeriodicEventStream"
	NotifyPriorityChargingAction            ChargingStationToCSMSAction = "NotifyPriorityCharging"
	NotifySettlementAction                  ChargingStationToCSMSAction = "NotifySettlement"
	OpenPeriodicEventStreamAction           ChargingStationToCSMSAction = "OpenPeriodicEventStream"
	ClosePeriodicEventStreamAction          ChargingStationToCSMSAction = "ClosePeriodicEventStream"
	PullDynamicScheduleUpdateAction         ChargingStationToCSMSAction = "PullDynamicScheduleUpdate"
	ReportDERControlAction                  ChargingStationToCSMSAction = "ReportDERControl"
	VatNumberValidationAction               ChargingStationToCSMSAction = "VatNumberValidation"

	CancelReservationAction           CSMSToChargingStationAction = "CancelReservation"
	ChangeAvailabilityAction          CSMSToChargingStationAction = "ChangeAvailability"
	ChangeConfigurationAction         CSMSToChargingStationAction = "ChangeConfiguration"
	ClearCacheAction                  CSMSToChargingStationAction = "ClearCache"
	ClearChargingProfileAction        CSMSToChargingStationAction = "ClearChargingProfile"
	DataTransferCSMSToCSAction        CSMSToChargingStationAction = "DataTransfer"
	GetCompositeScheduleAction        CSMSToChargingStationAction = "GetCompositeSchedule"
	GetConfigurationAction            CSMSToChargingStationAction = "GetConfiguration"
	GetDiagnosticsAction              CSMSToChargingStationAction = "GetDiagnostics"
	GetLocalListVersionAction         CSMSToChargingStationAction = "GetLocalListVersion"
	GetLogAction                      CSMSToChargingStationAction = "GetLog"
	RemoteStartTransactionAction      CSMSToChargingStationAction = "RemoteStartTransaction"
	RequestStartTransactionAction     CSMSToChargingStationAction = "RequestStartTransaction"
	RemoteStopTransactionAction       CSMSToChargingStationAction = "RemoteStopTransaction"
	RequestStopTransactionAction      CSMSToChargingStationAction = "RequestStopTransaction"
	ReserveNowAction                  CSMSToChargingStationAction = "ReserveNow"
	ResetAction                       CSMSToChargingStationAction = "Reset"
	SendLocalListAction               CSMSToChargingStationAction = "SendLocalList"
	SetChargingProfileAction          CSMSToChargingStationAction = "SetChargingProfile"
	GetChargingProfilesAction         CSMSToChargingStationAction = "GetChargingProfiles"
	TriggerMessageAction              CSMSToChargingStationAction = "TriggerMessage"
	UnlockConnectorAction             CSMSToChargingStationAction = "UnlockConnector"
	UpdateFirmwareAction              CSMSToChargingStationAction = "UpdateFirmware"
	SetVariablesAction                CSMSToChargingStationAction = "SetVariables"
	GetVariablesAction                CSMSToChargingStationAction = "GetVariables"
	GetBaseReportAction               CSMSToChargingStationAction = "GetBaseReport"
	GetInstalledCertificateIdsAction  CSMSToChargingStationAction = "GetInstalledCertificateIds"
	SetNetworkProfileAction           CSMSToChargingStationAction = "SetNetworkProfile"
	CertificateSignedAction           CSMSToChargingStationAction = "CertificateSigned"
	DeleteCertificateAction           CSMSToChargingStationAction = "DeleteCertificate"
	ExtendedTriggerMessageAction      CSMSToChargingStationAction = "ExtendedTriggerMessage"
	InstallCertificateAction          CSMSToChargingStationAction = "InstallCertificate"
	SignedUpdateFirmwareAction        CSMSToChargingStationAction = "SignedUpdateFirmware"
	ClearDisplayMessageAction         CSMSToChargingStationAction = "ClearDisplayMessage"
	ClearVariableMonitoringAction     CSMSToChargingStationAction = "ClearVariableMonitoring"
	CostUpdatedAction                 CSMSToChargingStationAction = "CostUpdated"
	CustomerInformationAction         CSMSToChargingStationAction = "CustomerInformation"
	GetDisplayMessagesAction          CSMSToChargingStationAction = "GetDisplayMessages"
	GetMonitoringReportAction         CSMSToChargingStationAction = "GetMonitoringReport"
	GetReportAction                   CSMSToChargingStationAction = "GetReport"
	GetTransactionStatusAction        CSMSToChargingStationAction = "GetTransactionStatus"
	PublishFirmwareAction             CSMSToChargingStationAction = "PublishFirmware"
	SetDisplayMessageAction           CSMSToChargingStationAction = "SetDisplayMessage"
	SetMonitoringBaseAction           CSMSToChargingStationAction = "SetMonitoringBase"
	SetMonitoringLevelAction          CSMSToChargingStationAction = "SetMonitoringLevel"
	SetVariableMonitoringAction       CSMSToChargingStationAction = "SetVariableMonitoring"
	UnpublishFirmwareAction           CSMSToChargingStationAction = "UnpublishFirmware"
	AFRRSignalAction                  CSMSToChargingStationAction = "AFRRSignal"
	AdjustPeriodicEventStreamAction   CSMSToChargingStationAction = "AdjustPeriodicEventStream"
	ChangeTransactionTariffAction     CSMSToChargingStationAction = "ChangeTransactionTariff"
	ClearDERControlAction             CSMSToChargingStationAction = "ClearDERControl"
	ClearTariffsAction                CSMSToChargingStationAction = "ClearTariffs"
	GetDERControlAction               CSMSToChargingStationAction = "GetDERControl"
	GetPeriodicEventStreamAction      CSMSToChargingStationAction = "GetPeriodicEventStream"
	GetTariffsAction                  CSMSToChargingStationAction = "GetTariffs"
	RequestBatterySwapAction          CSMSToChargingStationAction = "RequestBatterySwap"
	SetDERControlAction               CSMSToChargingStationAction = "SetDERControl"
	SetDefaultTariffAction            CSMSToChargingStationAction = "SetDefaultTariff"
	UpdateDynamicScheduleAction       CSMSToChargingStationAction = "UpdateDynamicSchedule"
	NotifyAllowedEnergyTransferAction CSMSToChargingStationAction = "NotifyAllowedEnergyTransfer"
	NotifyWebPaymentStartedAction     CSMSToChargingStationAction = "NotifyWebPaymentStarted"
	UsePriorityChargingAction         CSMSToChargingStationAction = "UsePriorityCharging"
)

func IsValidChargingStationToCSMSAction(action ChargingStationToCSMSAction) bool {
	switch action {
	case AuthorizeAction,
		BootNotificationAction,
		DataTransferCSToCSMSAction,
		DiagnosticsStatusNotificationAction,
		FirmwareStatusNotificationAction,
		SignedFirmwareStatusNotificationAction,
		HeartbeatAction,
		MeterValuesAction,
		NotifyReportAction,
		StartTransactionAction,
		StatusNotificationAction,
		StopTransactionAction,
		TransactionEventAction,
		SecurityEventNotificationAction,
		NotifyEventAction,
		NotifyEVChargingNeedsAction,
		LogStatusNotificationAction,
		SignCertificateAction,
		Get15118EVCertificateAction,
		GetCertificateStatusAction,
		ClearedChargingLimitAction,
		NotifyChargingLimitAction,
		NotifyCustomerInformationAction,
		NotifyDisplayMessagesAction,
		NotifyEVChargingScheduleAction,
		NotifyMonitoringReportAction,
		PublishFirmwareStatusNotificationAction,
		ReportChargingProfilesAction,
		ReservationStatusUpdateAction,
		BatterySwapAction,
		GetCertificateChainStatusAction,
		NotifyDERAlarmAction,
		NotifyDERStartStopAction,
		NotifyPeriodicEventStreamAction,
		NotifyPriorityChargingAction,
		NotifySettlementAction,
		OpenPeriodicEventStreamAction,
		ClosePeriodicEventStreamAction,
		PullDynamicScheduleUpdateAction,
		ReportDERControlAction,
		VatNumberValidationAction:
		return true
	default:
		return false
	}
}

func IsValidCSMSToChargingStationAction(action CSMSToChargingStationAction) bool {
	switch action {
	case CancelReservationAction,
		ChangeAvailabilityAction,
		ChangeConfigurationAction,
		ClearCacheAction,
		ClearChargingProfileAction,
		DataTransferCSMSToCSAction,
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
		SetDisplayMessageAction,
		SetMonitoringBaseAction,
		SetMonitoringLevelAction,
		SetVariableMonitoringAction,
		UnpublishFirmwareAction,
		AFRRSignalAction,
		AdjustPeriodicEventStreamAction,
		ChangeTransactionTariffAction,
		ClearDERControlAction,
		ClearTariffsAction,
		GetDERControlAction,
		GetPeriodicEventStreamAction,
		GetTariffsAction,
		RequestBatterySwapAction,
		SetDERControlAction,
		SetDefaultTariffAction,
		UpdateDynamicScheduleAction,
		NotifyAllowedEnergyTransferAction,
		NotifyWebPaymentStartedAction,
		UsePriorityChargingAction:
		return true
	default:
		return false
	}
}

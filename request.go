package ocpp

import (
	types15 "github.com/road-labs/ocpp-types-go/gen/ocpp15"
	types16 "github.com/road-labs/ocpp-types-go/gen/ocpp16"
	types201 "github.com/road-labs/ocpp-types-go/gen/ocpp201"
	types21 "github.com/road-labs/ocpp-types-go/gen/ocpp21"
)

func CSMSActionToRequestStruct(action CSMSToChargingStationAction, version Version) (any, error) {
	switch version {
	case Version15:
		return csmsActionToRequestStruct15(action)
	case Version16:
		return csmsActionToRequestStruct16(action)
	case Version201:
		return csmsActionToRequestStruct201(action)
	case Version21:
		return csmsActionToRequestStruct21(action)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func csmsActionToRequestStruct15(action CSMSToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types15.CancelReservation{}, nil
	case ChangeAvailabilityAction:
		return &types15.ChangeAvailability{}, nil
	case ChangeConfigurationAction:
		return &types15.ChangeConfiguration{}, nil
	case ClearCacheAction:
		return &types15.ClearCache{}, nil
	case DataTransferCSMSToCSAction:
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

func csmsActionToRequestStruct16(action CSMSToChargingStationAction) (any, error) {
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
	case DataTransferCSMSToCSAction:
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

func csmsActionToRequestStruct201(action CSMSToChargingStationAction) (any, error) {
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
	case DataTransferCSMSToCSAction:
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

func csmsActionToRequestStruct21(action CSMSToChargingStationAction) (any, error) {
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
	case DataTransferCSMSToCSAction:
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
	case AFRRSignalAction:
		return &types21.AFRRSignalRequest{}, nil
	case AdjustPeriodicEventStreamAction:
		return &types21.AdjustPeriodicEventStreamRequest{}, nil
	case ChangeTransactionTariffAction:
		return &types21.ChangeTransactionTariffRequest{}, nil
	case ClearDERControlAction:
		return &types21.ClearDERControlRequest{}, nil
	case ClearTariffsAction:
		return &types21.ClearTariffsRequest{}, nil
	case GetDERControlAction:
		return &types21.GetDERControlRequest{}, nil
	case GetPeriodicEventStreamAction:
		return &types21.GetPeriodicEventStreamRequest{}, nil
	case GetTariffsAction:
		return &types21.GetTariffsRequest{}, nil
	case RequestBatterySwapAction:
		return &types21.RequestBatterySwapRequest{}, nil
	case SetDERControlAction:
		return &types21.SetDERControlRequest{}, nil
	case SetDefaultTariffAction:
		return &types21.SetDefaultTariffRequest{}, nil
	case UpdateDynamicScheduleAction:
		return &types21.UpdateDynamicScheduleRequest{}, nil
	case NotifyAllowedEnergyTransferAction:
		return &types21.NotifyAllowedEnergyTransferRequest{}, nil
	case NotifyWebPaymentStartedAction:
		return &types21.NotifyWebPaymentStartedRequest{}, nil
	case UsePriorityChargingAction:
		return &types21.UsePriorityChargingRequest{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func ChargingStationActionToRequestStruct(action ChargingStationToCSMSAction, version Version) (any, error) {
	switch version {
	case Version15:
		return chargingStationActionToRequestStruct15(action)
	case Version16:
		return chargingStationActionToRequestStruct16(action)
	case Version201:
		return chargingStationActionToRequestStruct201(action)
	case Version21:
		return chargingStationActionToRequestStruct21(action)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func chargingStationActionToRequestStruct15(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types15.Authorize{}, nil
	case BootNotificationAction:
		return &types15.BootNotification{}, nil
	case DataTransferCSToCSMSAction:
		return &types15.DataTransfer{}, nil
	case DiagnosticsStatusNotificationAction:
		return &types15.DiagnosticsStatusNotification{}, nil
	case FirmwareStatusNotificationAction:
		return &types15.FirmwareStatusNotification{}, nil
	case HeartbeatAction:
		return &types15.Heartbeat{}, nil
	case MeterValuesAction:
		return &types15.MeterValues{}, nil
	case StartTransactionAction:
		return &types15.StartTransaction{}, nil
	case StatusNotificationAction:
		return &types15.StatusNotification{}, nil
	case StopTransactionAction:
		return &types15.StopTransaction{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToRequestStruct16(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types16.Authorize{}, nil
	case BootNotificationAction:
		return &types16.BootNotification{}, nil
	case DataTransferCSToCSMSAction:
		return &types16.DataTransfer{}, nil
	case DiagnosticsStatusNotificationAction:
		return &types16.DiagnosticsStatusNotification{}, nil
	case FirmwareStatusNotificationAction:
		return &types16.FirmwareStatusNotification{}, nil
	case HeartbeatAction:
		return &types16.Heartbeat{}, nil
	case LogStatusNotificationAction:
		return &types16.LogStatusNotification{}, nil
	case MeterValuesAction:
		return &types16.MeterValues{}, nil
	case SecurityEventNotificationAction:
		return &types16.SecurityEventNotification{}, nil
	case SignCertificateAction:
		return &types16.SignCertificate{}, nil
	case SignedFirmwareStatusNotificationAction:
		return &types16.SignedFirmwareStatusNotification{}, nil
	case StartTransactionAction:
		return &types16.StartTransaction{}, nil
	case StatusNotificationAction:
		return &types16.StatusNotification{}, nil
	case StopTransactionAction:
		return &types16.StopTransaction{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToRequestStruct201(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types201.AuthorizeRequest{}, nil
	case BootNotificationAction:
		return &types201.BootNotificationRequest{}, nil
	case ClearedChargingLimitAction:
		return &types201.ClearedChargingLimitRequest{}, nil
	case DataTransferCSToCSMSAction:
		return &types201.DataTransferRequest{}, nil
	case FirmwareStatusNotificationAction:
		return &types201.FirmwareStatusNotificationRequest{}, nil
	case Get15118EVCertificateAction:
		return &types201.Get15118EVCertificateRequest{}, nil
	case GetCertificateStatusAction:
		return &types201.GetCertificateStatusRequest{}, nil
	case HeartbeatAction:
		return &types201.HeartbeatRequest{}, nil
	case LogStatusNotificationAction:
		return &types201.LogStatusNotificationRequest{}, nil
	case MeterValuesAction:
		return &types201.MeterValuesRequest{}, nil
	case NotifyChargingLimitAction:
		return &types201.NotifyChargingLimitRequest{}, nil
	case NotifyCustomerInformationAction:
		return &types201.NotifyCustomerInformationRequest{}, nil
	case NotifyDisplayMessagesAction:
		return &types201.NotifyDisplayMessagesRequest{}, nil
	case NotifyEVChargingNeedsAction:
		return &types201.NotifyEVChargingNeedsRequest{}, nil
	case NotifyEVChargingScheduleAction:
		return &types201.NotifyEVChargingScheduleRequest{}, nil
	case NotifyEventAction:
		return &types201.NotifyEventRequest{}, nil
	case NotifyMonitoringReportAction:
		return &types201.NotifyMonitoringReportRequest{}, nil
	case NotifyReportAction:
		return &types201.NotifyReportRequest{}, nil
	case PublishFirmwareStatusNotificationAction:
		return &types201.PublishFirmwareStatusNotificationRequest{}, nil
	case ReportChargingProfilesAction:
		return &types201.ReportChargingProfilesRequest{}, nil
	case ReservationStatusUpdateAction:
		return &types201.ReservationStatusUpdateRequest{}, nil
	case SecurityEventNotificationAction:
		return &types201.SecurityEventNotificationRequest{}, nil
	case SignCertificateAction:
		return &types201.SignCertificateRequest{}, nil
	case StatusNotificationAction:
		return &types201.StatusNotificationRequest{}, nil
	case TransactionEventAction:
		return &types201.TransactionEventRequest{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToRequestStruct21(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types21.AuthorizeRequest{}, nil
	case BatterySwapAction:
		return &types21.BatterySwapRequest{}, nil
	case BootNotificationAction:
		return &types21.BootNotificationRequest{}, nil
	case ClearedChargingLimitAction:
		return &types21.ClearedChargingLimitRequest{}, nil
	case ClosePeriodicEventStreamAction:
		return &types21.ClosePeriodicEventStreamRequest{}, nil
	case DataTransferCSToCSMSAction:
		return &types21.DataTransferRequest{}, nil
	case FirmwareStatusNotificationAction:
		return &types21.FirmwareStatusNotificationRequest{}, nil
	case Get15118EVCertificateAction:
		return &types21.Get15118EVCertificateRequest{}, nil
	case GetCertificateChainStatusAction:
		return &types21.GetCertificateChainStatusRequest{}, nil
	case GetCertificateStatusAction:
		return &types21.GetCertificateStatusRequest{}, nil
	case HeartbeatAction:
		return &types21.HeartbeatRequest{}, nil
	case LogStatusNotificationAction:
		return &types21.LogStatusNotificationRequest{}, nil
	case MeterValuesAction:
		return &types21.MeterValuesRequest{}, nil
	case NotifyChargingLimitAction:
		return &types21.NotifyChargingLimitRequest{}, nil
	case NotifyCustomerInformationAction:
		return &types21.NotifyCustomerInformationRequest{}, nil
	case NotifyDERAlarmAction:
		return &types21.NotifyDERAlarmRequest{}, nil
	case NotifyDERStartStopAction:
		return &types21.NotifyDERStartStopRequest{}, nil
	case NotifyDisplayMessagesAction:
		return &types21.NotifyDisplayMessagesRequest{}, nil
	case NotifyEVChargingNeedsAction:
		return &types21.NotifyEVChargingNeedsRequest{}, nil
	case NotifyEVChargingScheduleAction:
		return &types21.NotifyEVChargingScheduleRequest{}, nil
	case NotifyEventAction:
		return &types21.NotifyEventRequest{}, nil
	case NotifyMonitoringReportAction:
		return &types21.NotifyMonitoringReportRequest{}, nil
	case NotifyPeriodicEventStreamAction:
		return &types21.NotifyPeriodicEventStream{}, nil
	case NotifyPriorityChargingAction:
		return &types21.NotifyPriorityChargingRequest{}, nil
	case NotifyReportAction:
		return &types21.NotifyReportRequest{}, nil
	case NotifySettlementAction:
		return &types21.NotifySettlementRequest{}, nil
	case OpenPeriodicEventStreamAction:
		return &types21.OpenPeriodicEventStreamRequest{}, nil
	case PublishFirmwareStatusNotificationAction:
		return &types21.PublishFirmwareStatusNotificationRequest{}, nil
	case PullDynamicScheduleUpdateAction:
		return &types21.PullDynamicScheduleUpdateRequest{}, nil
	case ReportChargingProfilesAction:
		return &types21.ReportChargingProfilesRequest{}, nil
	case ReportDERControlAction:
		return &types21.ReportDERControlRequest{}, nil
	case ReservationStatusUpdateAction:
		return &types21.ReservationStatusUpdateRequest{}, nil
	case SecurityEventNotificationAction:
		return &types21.SecurityEventNotificationRequest{}, nil
	case SignCertificateAction:
		return &types21.SignCertificateRequest{}, nil
	case StatusNotificationAction:
		return &types21.StatusNotificationRequest{}, nil
	case TransactionEventAction:
		return &types21.TransactionEventRequest{}, nil
	case VatNumberValidationAction:
		return &types21.VatNumberValidationRequest{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

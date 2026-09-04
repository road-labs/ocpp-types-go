package ocpp

import (
	types15 "github.com/road-labs/ocpp-types-go/gen/ocpp15"
	types16 "github.com/road-labs/ocpp-types-go/gen/ocpp16"
	types201 "github.com/road-labs/ocpp-types-go/gen/ocpp201"
	types21 "github.com/road-labs/ocpp-types-go/gen/ocpp21"
)

func CSMSActionToResponseStruct(action CSMSToChargingStationAction, version Version) (any, error) {
	switch version {
	case Version15:
		return csmsActionToResponseStruct15(action)
	case Version16:
		return csmsActionToResponseStruct16(action)
	case Version201:
		return csmsActionToResponseStruct201(action)
	case Version21:
		return csmsActionToResponseStruct21(action)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func csmsActionToResponseStruct15(action CSMSToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types15.CancelReservationResponse{}, nil
	case ChangeAvailabilityAction:
		return &types15.ChangeAvailabilityResponse{}, nil
	case ChangeConfigurationAction:
		return &types15.ChangeConfigurationResponse{}, nil
	case ClearCacheAction:
		return &types15.ClearCacheResponse{}, nil
	case DataTransferCSMSToCSAction:
		return &types15.DataTransferResponse{}, nil
	case GetConfigurationAction:
		return &types15.GetConfigurationResponse{}, nil
	case GetDiagnosticsAction:
		return &types15.GetDiagnosticsResponse{}, nil
	case GetLocalListVersionAction:
		return &types15.GetLocalListVersionResponse{}, nil
	case RemoteStartTransactionAction:
		return &types15.RemoteStartTransactionResponse{}, nil
	case RemoteStopTransactionAction:
		return &types15.RemoteStopTransactionResponse{}, nil
	case ReserveNowAction:
		return &types15.ReserveNowResponse{}, nil
	case ResetAction:
		return &types15.ResetResponse{}, nil
	case SendLocalListAction:
		return &types15.SendLocalListResponse{}, nil
	case UnlockConnectorAction:
		return &types15.UnlockConnectorResponse{}, nil
	case UpdateFirmwareAction:
		return &types15.UpdateFirmwareResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func csmsActionToResponseStruct16(action CSMSToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types16.CancelReservationResponse{}, nil
	case CertificateSignedAction:
		return &types16.CertificateSignedResponse{}, nil
	case ChangeAvailabilityAction:
		return &types16.ChangeAvailabilityResponse{}, nil
	case ChangeConfigurationAction:
		return &types16.ChangeConfigurationResponse{}, nil
	case ClearCacheAction:
		return &types16.ClearCacheResponse{}, nil
	case ClearChargingProfileAction:
		return &types16.ClearChargingProfileResponse{}, nil
	case DataTransferCSMSToCSAction:
		return &types16.DataTransferResponse{}, nil
	case DeleteCertificateAction:
		return &types16.DeleteCertificateResponse{}, nil
	case ExtendedTriggerMessageAction:
		return &types16.ExtendedTriggerMessageResponse{}, nil
	case GetCompositeScheduleAction:
		return &types16.GetCompositeScheduleResponse{}, nil
	case GetConfigurationAction:
		return &types16.GetConfigurationResponse{}, nil
	case GetDiagnosticsAction:
		return &types16.GetDiagnosticsResponse{}, nil
	case GetInstalledCertificateIdsAction:
		return &types16.GetInstalledCertificateIdsResponse{}, nil
	case GetLocalListVersionAction:
		return &types16.GetLocalListVersionResponse{}, nil
	case GetLogAction:
		return &types16.GetLogResponse{}, nil
	case InstallCertificateAction:
		return &types16.InstallCertificateResponse{}, nil
	case RemoteStartTransactionAction:
		return &types16.RemoteStartTransactionResponse{}, nil
	case RemoteStopTransactionAction:
		return &types16.RemoteStopTransactionResponse{}, nil
	case ReserveNowAction:
		return &types16.ReserveNowResponse{}, nil
	case ResetAction:
		return &types16.ResetResponse{}, nil
	case SendLocalListAction:
		return &types16.SendLocalListResponse{}, nil
	case SetChargingProfileAction:
		return &types16.SetChargingProfileResponse{}, nil
	case SignedUpdateFirmwareAction:
		return &types16.SignedUpdateFirmwareResponse{}, nil
	case TriggerMessageAction:
		return &types16.TriggerMessageResponse{}, nil
	case UnlockConnectorAction:
		return &types16.UnlockConnectorResponse{}, nil
	case UpdateFirmwareAction:
		return &types16.UpdateFirmwareResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func csmsActionToResponseStruct201(action CSMSToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types201.CancelReservationResponse{}, nil
	case CertificateSignedAction:
		return &types201.CertificateSignedResponse{}, nil
	case ChangeAvailabilityAction:
		return &types201.ChangeAvailabilityResponse{}, nil
	case ClearCacheAction:
		return &types201.ClearCacheResponse{}, nil
	case ClearChargingProfileAction:
		return &types201.ClearChargingProfileResponse{}, nil
	case ClearDisplayMessageAction:
		return &types201.ClearDisplayMessageResponse{}, nil
	case ClearVariableMonitoringAction:
		return &types201.ClearVariableMonitoringResponse{}, nil
	case CostUpdatedAction:
		return &types201.CostUpdatedResponse{}, nil
	case CustomerInformationAction:
		return &types201.CustomerInformationResponse{}, nil
	case DataTransferCSMSToCSAction:
		return &types201.DataTransferResponse{}, nil
	case DeleteCertificateAction:
		return &types201.DeleteCertificateResponse{}, nil
	case GetBaseReportAction:
		return &types201.GetBaseReportResponse{}, nil
	case GetChargingProfilesAction:
		return &types201.GetChargingProfilesResponse{}, nil
	case GetCompositeScheduleAction:
		return &types201.GetCompositeScheduleResponse{}, nil
	case GetDisplayMessagesAction:
		return &types201.GetDisplayMessagesResponse{}, nil
	case GetInstalledCertificateIdsAction:
		return &types201.GetInstalledCertificateIdsResponse{}, nil
	case GetLocalListVersionAction:
		return &types201.GetLocalListVersionResponse{}, nil
	case GetLogAction:
		return &types201.GetLogResponse{}, nil
	case GetMonitoringReportAction:
		return &types201.GetMonitoringReportResponse{}, nil
	case GetReportAction:
		return &types201.GetReportResponse{}, nil
	case GetTransactionStatusAction:
		return &types201.GetTransactionStatusResponse{}, nil
	case GetVariablesAction:
		return &types201.GetVariablesResponse{}, nil
	case InstallCertificateAction:
		return &types201.InstallCertificateResponse{}, nil
	case PublishFirmwareAction:
		return &types201.PublishFirmwareResponse{}, nil
	case RequestStartTransactionAction:
		return &types201.RequestStartTransactionResponse{}, nil
	case RequestStopTransactionAction:
		return &types201.RequestStopTransactionResponse{}, nil
	case ReserveNowAction:
		return &types201.ReserveNowResponse{}, nil
	case ResetAction:
		return &types201.ResetResponse{}, nil
	case SendLocalListAction:
		return &types201.SendLocalListResponse{}, nil
	case SetChargingProfileAction:
		return &types201.SetChargingProfileResponse{}, nil
	case SetDisplayMessageAction:
		return &types201.SetDisplayMessageResponse{}, nil
	case SetMonitoringBaseAction:
		return &types201.SetMonitoringBaseResponse{}, nil
	case SetMonitoringLevelAction:
		return &types201.SetMonitoringLevelResponse{}, nil
	case SetNetworkProfileAction:
		return &types201.SetNetworkProfileResponse{}, nil
	case SetVariableMonitoringAction:
		return &types201.SetVariableMonitoringResponse{}, nil
	case SetVariablesAction:
		return &types201.SetVariablesResponse{}, nil
	case TriggerMessageAction:
		return &types201.TriggerMessageResponse{}, nil
	case UnlockConnectorAction:
		return &types201.UnlockConnectorResponse{}, nil
	case UnpublishFirmwareAction:
		return &types201.UnpublishFirmwareResponse{}, nil
	case UpdateFirmwareAction:
		return &types201.UpdateFirmwareResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func csmsActionToResponseStruct21(action CSMSToChargingStationAction) (any, error) {
	switch action {
	case CancelReservationAction:
		return &types21.CancelReservationResponse{}, nil
	case CertificateSignedAction:
		return &types21.CertificateSignedResponse{}, nil
	case ChangeAvailabilityAction:
		return &types21.ChangeAvailabilityResponse{}, nil
	case ClearCacheAction:
		return &types21.ClearCacheResponse{}, nil
	case ClearChargingProfileAction:
		return &types21.ClearChargingProfileResponse{}, nil
	case ClearDisplayMessageAction:
		return &types21.ClearDisplayMessageResponse{}, nil
	case ClearVariableMonitoringAction:
		return &types21.ClearVariableMonitoringResponse{}, nil
	case CostUpdatedAction:
		return &types21.CostUpdatedResponse{}, nil
	case CustomerInformationAction:
		return &types21.CustomerInformationResponse{}, nil
	case DataTransferCSMSToCSAction:
		return &types21.DataTransferResponse{}, nil
	case DeleteCertificateAction:
		return &types21.DeleteCertificateResponse{}, nil
	case GetBaseReportAction:
		return &types21.GetBaseReportResponse{}, nil
	case GetChargingProfilesAction:
		return &types21.GetChargingProfilesResponse{}, nil
	case GetCompositeScheduleAction:
		return &types21.GetCompositeScheduleResponse{}, nil
	case GetDisplayMessagesAction:
		return &types21.GetDisplayMessagesResponse{}, nil
	case GetInstalledCertificateIdsAction:
		return &types21.GetInstalledCertificateIdsResponse{}, nil
	case GetLocalListVersionAction:
		return &types21.GetLocalListVersionResponse{}, nil
	case GetLogAction:
		return &types21.GetLogResponse{}, nil
	case GetMonitoringReportAction:
		return &types21.GetMonitoringReportResponse{}, nil
	case GetReportAction:
		return &types21.GetReportResponse{}, nil
	case GetTransactionStatusAction:
		return &types21.GetTransactionStatusResponse{}, nil
	case GetVariablesAction:
		return &types21.GetVariablesResponse{}, nil
	case InstallCertificateAction:
		return &types21.InstallCertificateResponse{}, nil
	case PublishFirmwareAction:
		return &types21.PublishFirmwareResponse{}, nil
	case RequestStartTransactionAction:
		return &types21.RequestStartTransactionResponse{}, nil
	case RequestStopTransactionAction:
		return &types21.RequestStopTransactionResponse{}, nil
	case ReserveNowAction:
		return &types21.ReserveNowResponse{}, nil
	case ResetAction:
		return &types21.ResetResponse{}, nil
	case SendLocalListAction:
		return &types21.SendLocalListResponse{}, nil
	case SetChargingProfileAction:
		return &types21.SetChargingProfileResponse{}, nil
	case SetDisplayMessageAction:
		return &types21.SetDisplayMessageResponse{}, nil
	case SetMonitoringBaseAction:
		return &types21.SetMonitoringBaseResponse{}, nil
	case SetMonitoringLevelAction:
		return &types21.SetMonitoringLevelResponse{}, nil
	case SetNetworkProfileAction:
		return &types21.SetNetworkProfileResponse{}, nil
	case SetVariableMonitoringAction:
		return &types21.SetVariableMonitoringResponse{}, nil
	case SetVariablesAction:
		return &types21.SetVariablesResponse{}, nil
	case TriggerMessageAction:
		return &types21.TriggerMessageResponse{}, nil
	case UnlockConnectorAction:
		return &types21.UnlockConnectorResponse{}, nil
	case UnpublishFirmwareAction:
		return &types21.UnpublishFirmwareResponse{}, nil
	case UpdateFirmwareAction:
		return &types21.UpdateFirmwareResponse{}, nil
	case AFRRSignalAction:
		return &types21.AFRRSignalResponse{}, nil
	case AdjustPeriodicEventStreamAction:
		return &types21.AdjustPeriodicEventStreamResponse{}, nil
	case ChangeTransactionTariffAction:
		return &types21.ChangeTransactionTariffResponse{}, nil
	case ClearDERControlAction:
		return &types21.ClearDERControlResponse{}, nil
	case ClearTariffsAction:
		return &types21.ClearTariffsResponse{}, nil
	case GetDERControlAction:
		return &types21.GetDERControlResponse{}, nil
	case GetPeriodicEventStreamAction:
		return &types21.GetPeriodicEventStreamResponse{}, nil
	case GetTariffsAction:
		return &types21.GetTariffsResponse{}, nil
	case RequestBatterySwapAction:
		return &types21.RequestBatterySwapResponse{}, nil
	case SetDERControlAction:
		return &types21.SetDERControlResponse{}, nil
	case SetDefaultTariffAction:
		return &types21.SetDefaultTariffResponse{}, nil
	case UpdateDynamicScheduleAction:
		return &types21.UpdateDynamicScheduleResponse{}, nil
	case NotifyAllowedEnergyTransferAction:
		return &types21.NotifyAllowedEnergyTransferResponse{}, nil
	case NotifyWebPaymentStartedAction:
		return &types21.NotifyWebPaymentStartedResponse{}, nil
	case UsePriorityChargingAction:
		return &types21.UsePriorityChargingResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func ChargingStationActionToResponseStruct(action ChargingStationToCSMSAction, version Version) (any, error) {
	switch version {
	case Version15:
		return chargingStationActionToResponseStruct15(action)
	case Version16:
		return chargingStationActionToResponseStruct16(action)
	case Version201:
		return chargingStationActionToResponseStruct201(action)
	case Version21:
		return chargingStationActionToResponseStruct21(action)
	default:
		return nil, ErrUnsupportedVersion
	}
}

func chargingStationActionToResponseStruct15(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types15.AuthorizeResponse{}, nil
	case BootNotificationAction:
		return &types15.BootNotificationResponse{}, nil
	case DataTransferCSToCSMSAction:
		return &types15.DataTransferResponse{}, nil
	case DiagnosticsStatusNotificationAction:
		return &types15.DiagnosticsStatusNotificationResponse{}, nil
	case FirmwareStatusNotificationAction:
		return &types15.FirmwareStatusNotificationResponse{}, nil
	case HeartbeatAction:
		return &types15.HeartbeatResponse{}, nil
	case MeterValuesAction:
		return &types15.MeterValuesResponse{}, nil
	case StartTransactionAction:
		return &types15.StartTransactionResponse{}, nil
	case StatusNotificationAction:
		return &types15.StatusNotificationResponse{}, nil
	case StopTransactionAction:
		return &types15.StopTransactionResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToResponseStruct16(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types16.AuthorizeResponse{}, nil
	case BootNotificationAction:
		return &types16.BootNotificationResponse{}, nil
	case DataTransferCSToCSMSAction:
		return &types16.DataTransferResponse{}, nil
	case DiagnosticsStatusNotificationAction:
		return &types16.DiagnosticsStatusNotificationResponse{}, nil
	case FirmwareStatusNotificationAction:
		return &types16.FirmwareStatusNotificationResponse{}, nil
	case HeartbeatAction:
		return &types16.HeartbeatResponse{}, nil
	case LogStatusNotificationAction:
		return &types16.LogStatusNotificationResponse{}, nil
	case MeterValuesAction:
		return &types16.MeterValuesResponse{}, nil
	case SecurityEventNotificationAction:
		return &types16.SecurityEventNotificationResponse{}, nil
	case SignCertificateAction:
		return &types16.SignCertificateResponse{}, nil
	case SignedFirmwareStatusNotificationAction:
		return &types16.SignedFirmwareStatusNotificationResponse{}, nil
	case StartTransactionAction:
		return &types16.StartTransactionResponse{}, nil
	case StatusNotificationAction:
		return &types16.StatusNotificationResponse{}, nil
	case StopTransactionAction:
		return &types16.StopTransactionResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToResponseStruct201(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types201.AuthorizeResponse{}, nil
	case BootNotificationAction:
		return &types201.BootNotificationResponse{}, nil
	case ClearedChargingLimitAction:
		return &types201.ClearedChargingLimitResponse{}, nil
	case DataTransferCSToCSMSAction:
		return &types201.DataTransferResponse{}, nil
	case FirmwareStatusNotificationAction:
		return &types201.FirmwareStatusNotificationResponse{}, nil
	case Get15118EVCertificateAction:
		return &types201.Get15118EVCertificateResponse{}, nil
	case GetCertificateStatusAction:
		return &types201.GetCertificateStatusResponse{}, nil
	case HeartbeatAction:
		return &types201.HeartbeatResponse{}, nil
	case LogStatusNotificationAction:
		return &types201.LogStatusNotificationResponse{}, nil
	case MeterValuesAction:
		return &types201.MeterValuesResponse{}, nil
	case NotifyChargingLimitAction:
		return &types201.NotifyChargingLimitResponse{}, nil
	case NotifyCustomerInformationAction:
		return &types201.NotifyCustomerInformationResponse{}, nil
	case NotifyDisplayMessagesAction:
		return &types201.NotifyDisplayMessagesResponse{}, nil
	case NotifyEVChargingNeedsAction:
		return &types201.NotifyEVChargingNeedsResponse{}, nil
	case NotifyEVChargingScheduleAction:
		return &types201.NotifyEVChargingScheduleResponse{}, nil
	case NotifyEventAction:
		return &types201.NotifyEventResponse{}, nil
	case NotifyMonitoringReportAction:
		return &types201.NotifyMonitoringReportResponse{}, nil
	case NotifyReportAction:
		return &types201.NotifyReportResponse{}, nil
	case PublishFirmwareStatusNotificationAction:
		return &types201.PublishFirmwareStatusNotificationResponse{}, nil
	case ReportChargingProfilesAction:
		return &types201.ReportChargingProfilesResponse{}, nil
	case ReservationStatusUpdateAction:
		return &types201.ReservationStatusUpdateResponse{}, nil
	case SecurityEventNotificationAction:
		return &types201.SecurityEventNotificationResponse{}, nil
	case SignCertificateAction:
		return &types201.SignCertificateResponse{}, nil
	case StatusNotificationAction:
		return &types201.StatusNotificationResponse{}, nil
	case TransactionEventAction:
		return &types201.TransactionEventResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

func chargingStationActionToResponseStruct21(action ChargingStationToCSMSAction) (any, error) {
	switch action {
	case AuthorizeAction:
		return &types21.AuthorizeResponse{}, nil
	case BatterySwapAction:
		return &types21.BatterySwapResponse{}, nil
	case BootNotificationAction:
		return &types21.BootNotificationResponse{}, nil
	case ClearedChargingLimitAction:
		return &types21.ClearedChargingLimitResponse{}, nil
	case ClosePeriodicEventStreamAction:
		return &types21.ClosePeriodicEventStreamResponse{}, nil
	case DataTransferCSToCSMSAction:
		return &types21.DataTransferResponse{}, nil
	case FirmwareStatusNotificationAction:
		return &types21.FirmwareStatusNotificationResponse{}, nil
	case Get15118EVCertificateAction:
		return &types21.Get15118EVCertificateResponse{}, nil
	case GetCertificateChainStatusAction:
		return &types21.GetCertificateChainStatusResponse{}, nil
	case GetCertificateStatusAction:
		return &types21.GetCertificateStatusResponse{}, nil
	case HeartbeatAction:
		return &types21.HeartbeatResponse{}, nil
	case LogStatusNotificationAction:
		return &types21.LogStatusNotificationResponse{}, nil
	case MeterValuesAction:
		return &types21.MeterValuesResponse{}, nil
	case NotifyChargingLimitAction:
		return &types21.NotifyChargingLimitResponse{}, nil
	case NotifyCustomerInformationAction:
		return &types21.NotifyCustomerInformationResponse{}, nil
	case NotifyDERAlarmAction:
		return &types21.NotifyDERAlarmResponse{}, nil
	case NotifyDERStartStopAction:
		return &types21.NotifyDERStartStopResponse{}, nil
	case NotifyDisplayMessagesAction:
		return &types21.NotifyDisplayMessagesResponse{}, nil
	case NotifyEVChargingNeedsAction:
		return &types21.NotifyEVChargingNeedsResponse{}, nil
	case NotifyEVChargingScheduleAction:
		return &types21.NotifyEVChargingScheduleResponse{}, nil
	case NotifyEventAction:
		return &types21.NotifyEventResponse{}, nil
	case NotifyMonitoringReportAction:
		return &types21.NotifyMonitoringReportResponse{}, nil
	case NotifyPriorityChargingAction:
		return &types21.NotifyPriorityChargingResponse{}, nil
	case NotifyReportAction:
		return &types21.NotifyReportResponse{}, nil
	case NotifySettlementAction:
		return &types21.NotifySettlementResponse{}, nil
	case OpenPeriodicEventStreamAction:
		return &types21.OpenPeriodicEventStreamResponse{}, nil
	case PublishFirmwareStatusNotificationAction:
		return &types21.PublishFirmwareStatusNotificationResponse{}, nil
	case PullDynamicScheduleUpdateAction:
		return &types21.PullDynamicScheduleUpdateResponse{}, nil
	case ReportChargingProfilesAction:
		return &types21.ReportChargingProfilesResponse{}, nil
	case ReportDERControlAction:
		return &types21.ReportDERControlResponse{}, nil
	case ReservationStatusUpdateAction:
		return &types21.ReservationStatusUpdateResponse{}, nil
	case SecurityEventNotificationAction:
		return &types21.SecurityEventNotificationResponse{}, nil
	case SignCertificateAction:
		return &types21.SignCertificateResponse{}, nil
	case StatusNotificationAction:
		return &types21.StatusNotificationResponse{}, nil
	case TransactionEventAction:
		return &types21.TransactionEventResponse{}, nil
	case VatNumberValidationAction:
		return &types21.VatNumberValidationResponse{}, nil
	default:
		return nil, ErrUnknownAction
	}
}

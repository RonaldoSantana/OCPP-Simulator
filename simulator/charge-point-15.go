package simulator

type ChargePointInterface interface {
	// Methods Initiated by the charge point
	Authorize() int
	BootNotification() int
	DataTransfer()
	DiagnosticsStatusNotification()
	FirmwareStatusNotification()
	Heartbeat()
	MeterValues()
	StartTransaction()
	StatusNotification()
	StopTransaction()

	// start communication
	request()
	confirm()

	// Methods Initiated by a Central System that the charge point needs to be able to respond
	HandleCancelReservation()
	HandleChangeAvailability()
	HandleChangeConfiguration()
	HandleClearCache()
	HandleDataTransfer()
	HandleGetConfiguration()
	HandleGetDiagnostics()
	HandleGetLocalListVersion()
	HandleRemoteStartTransaction()
	HandleRemoteStopTransaction()
	HandleReserveNow()
	HandleReset()
	HandleSendLocalList()
	HandleUnlockConnector()
	HandleUpdateFirmware()

}


type ChargePoint struct {

}
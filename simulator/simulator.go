package simulator

import (
	"encoding/xml"
	/*"text/template"*/
	/*"bytes"
	"log"*/
	/*"github.com/RonaldoSantana/ocpp-simulator/soap"*/
)

const StatusAccepted = "Accepted"
const StatusBlocked = "Blocked"
const StatusExpired = "Expired"
const StatusInvalid = "Invalid"
const StatusConcurrenTx = "ConcurrentTx"

// Basic request parameters
type RequestData struct {
	ChargeBoxID string
	AuthID string
}

/*// XML Body
type XMLBody struct {
	XMLName	xml.Name `xml:"Body"`
}

// XML Envelope
type XMLEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody	XMLBody
}*/

// empty interface so every request has it' own definition

// Interface that all charge point Request types needs to implement
type ChargePoint interface {
	ParseRequestBody(data []string) string // the request XML to be posted to central system
	ParseResponseBody(responseData []byte) // the parsed response, according to request
	ResponseStatus() string
}

func request(requestMethod ChargePoint) {

	/*var buffer bytes.Buffer
	tplData := AuthTemplateData{
		ChargeBoxID: "veefil-21159",
		AuthID: "B4F62CEF",
	}

	err := Tpl.ExecuteTemplate(buffer, requestMethod.Template(), tplData)
	if err != nil {
		log.Fatalln(err)
	}*/

	/*soap := soap.Request{
		Url : "https://ocpp.ron.testcharge.net.nz",
	}
*/


	//soap.Call()

}

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
package simulator

import (
	"encoding/xml"
	"text/template"
	"bytes"
	"os"
	"log"
	"github.com/RonaldoSantana/ocpp-simulator/soap"
)

// global package template variable
var Tpl *template.Template

const statusAccepted = "Accepted"
const statusBlocked = "Blocked"
const statusExpired = "Expired"
const statusInvalid = "Invalid"
const statusConcurrenTx = "ConcurrentTx"

// XML Body
type XMLBody struct {
	XMLName	xml.Name `xml:"Body"`
}

// XML Envelope
type XMLEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody	XMLBody
}

// empty interface so every request has it' own definition
type RequestData interface {}

// Interface that all charge point Request types needs to implement
type ChargePoint interface {
	ParseRequestBody() string // the request XML to be posted to central system
	ParseResponseBody(requestData RequestData) interface{} // the parsed response, according to request
	Template() string // the parsed response, according to request
}

type ChargePointRequest struct {
	Template string // the request XML to be posted to central system
}

func request(requestMethod ChargePoint) {

	var buffer bytes.Buffer
	tplData := AuthTemplateData{
		ChargeBoxID: "veefil-21159",
		AuthID: "B4F62CEF",
	}

	err := Tpl.ExecuteTemplate(buffer, requestMethod.Template(), tplData)
	if err != nil {
		log.Fatalln(err)
	}

	soap := soap.Request{
		Url : "https://ocpp.ron.testcharge.net.nz",
	}



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


type ChargePoint struct {

}
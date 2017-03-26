package charger

import (
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
	"encoding/json"
)

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

type Simulator struct {
	Method ChargePointMethod
	config Config
}

type Config struct {
	CentralSystemUrl string
}

var config *Config

func init() {
	file, e := ioutil.ReadFile("./config/env.json")
	if e != nil {
		log.Fatalf("File error: %v\n", e)
	}

	json.Unmarshal(file, &config)
}

func (simulator *Simulator) Call(method string, args ...string) string {

	switch method {
	case "Authorize":
		simulator.Method = NewAuthorize()
	case "StartTransaction":
		simulator.Method = NewStartTransaction()
	case "StopTransaction":
		simulator.Method = NewStopTransaction()
	default:
		log.Fatalf(`Invalid method "%v" called`, method )
	}

	simulator.soapCall(args)

	return simulator.Method.ResponseStatus()
}

func (simulator *Simulator) soapCall(args []string) {

	soapRequestContent := simulator.Method.ParseRequestBody(args)

	httpClient := new(http.Client)

	// make request to central system
	response, err := httpClient.Post(config.CentralSystemUrl, "text/xml; charset=utf-8", bytes.NewBufferString(soapRequestContent))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	simulator.Method.ParseResponseBody(responseBody)
}
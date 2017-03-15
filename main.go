package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"github.com/RonaldoSantana/ocpp-simulator/simulator"
)

// The URL of the SOAP server
const MH_SOAP_URL = "https://ocpp.ron.testcharge.net.nz"

func soapCall(request simulator.ChargePoint) {

	// TODO: this will be the command line arguments
	args := []string {
		"Authorize",
		"veefil-21159",
		"B4F62CEF",
	}
	soapRequestContent := request.ParseRequestBody(args);

	httpClient := new(http.Client)
	// make request to central system
	response, err := httpClient.Post(MH_SOAP_URL, "text/xml; charset=utf-8", bytes.NewBufferString(soapRequestContent))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	request.ParseResponseBody(responseBody)
}

func main() {
	var request simulator.ChargePoint
	method := "Authorize"

	switch ( method ) {
	case "Authorize" :
		request = simulator.NewAuthorize()
	case "StartTransaction" :
		request = simulator.NewStartTransaction()
	default:
		// TODO: invalid request, stop everything
		request = simulator.NewAuthorize()
	}
	soapCall(request)

	fmt.Println(request.ResponseStatus())
}
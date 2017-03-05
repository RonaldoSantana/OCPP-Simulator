package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

// The URL of the SOAP server
const MH_SOAP_URL = "https://ocpp.ron.testcharge.net.nz"

// this is just the message I'll send for interrogation, with placeholders
//  for my parameters
const SOAP_VUE_QUERY_FORMAT = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<x:Envelope
		xmlns:x="http://schemas.xmlsoap.org/soap/envelope/"
	 	xmlns:ns="urn://Ocpp/Cs/2012/06/">
	 	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:tns="urn:SP_WebService"
        xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/"
        xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/"
        xmlns:x="http://schemas.xmlsoap.org/soap/encoding/" >
    <x:Header>
        <ns:chargeBoxIdentity>%s</ns:chargeBoxIdentity>
    </x:Header>
    <x:Body>
        <ns:authorizeRequest>
            <ns:idTag>%s</ns:idTag>
        </ns:authorizeRequest>
    </x:Body>
</x:Envelope>
`

const test2 = `<soap:Envelope
		xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
		xmlns:tns="urn://Ocpp/Cs/2012/06/"
		xmlns:wsp="http://schemas.xmlsoapf.org/ws/2004/09/policy"
		xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
		xmlns:wsa5="http://www.w3.org/2005/08/addressing">
		<soap:Header>
			<tns:chargeBoxIdentity>%s</tns:chargeBoxIdentity>
			<wsa5:To>https://ocpp.ron.testcharge.net.nz</wsa5:To>
			<wsa5:From>
				<wsa5:Address>http://10.1.2.62:40173/</wsa5:Address>
			</wsa5:From>
			<wsa5:Action>/Authorize</wsa5:Action>
		</soap:Header>
		<soap:Body>
			<tns:authorizeRequest>
				<tns:idTag>%s</tns:idTag>
			</tns:authorizeRequest>
		</soap:Body>
	</soap:Envelope>`

const test = `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="urn://Ocpp/Cs/2012/06/">
    <x:Header>
        <ns:chargeBoxIdentity>%s</ns:chargeBoxIdentity>
    </x:Header>
    <x:Body>
        <ns:authorizeRequest>
            <ns:idTag>%s</ns:idTag>
        </ns:authorizeRequest>
    </x:Body>
</x:Envelope>`

// Here I define Go structures, almost identical to the structure of the
// XML message we'll fetch
// Note that annotations (the string "return>item") allow to have a slightly
//  different structure or different namings

type SoapItem struct {
	Numero    int
	Nom       string
	Type      string
	PositionX int
	PositionY int
	PositionN int
	Monde     int
}
type SoapVue struct {
	Items []SoapItem "return>item"
}
type SoapProfil struct { // un peu incomplet, certes
	Numero int
}
type SoapFault struct {
	Faultstring string
	Detail      string
}
type SoapBody struct {
	Fault          SoapFault
	ProfilResponse SoapProfil
	VueResponse    SoapVue
}
type SoapEnvelope struct {
	XMLName xml.Name
	Body    SoapBody
}

// Here is the function querying the SOAP server
// It returns the whole answer as a Go structure (a SoapEnvelope)
// You could also return an error in a second returned parameter
func GetSoapEnvelope(query string, numero string, mdp string) (envelope *SoapEnvelope) {
	soapRequestContent := fmt.Sprintf(query, numero, mdp)
	httpClient := new(http.Client)
	// make request to central system
	response, err := httpClient.Post(MH_SOAP_URL, "text/xml; charset=utf-8", bytes.NewBufferString(soapRequestContent))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body) // probably not efficient, done because the stream isn't always a pure XML stream and I have to fix things (not shown here)
	if err != nil {
		log.Fatalln(err)
	}
	in := string(b)
	fmt.Println(in)
	parser := xml.NewDecoder(bytes.NewBufferString(in))
	envelope = new(SoapEnvelope) // this allocates the structure in which we'll decode the XML
	err = parser.DecodeElement(&envelope, nil)
	if err != nil {
		// handle error
		log.Fatalln(err)
	}

	return
}

func main() {
	GetSoapEnvelope(test, "veefil-21159", "B4F62CEF")
}

/*
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Post("http://www.soapui.org", " text/xml;charset=UTF-8", strings.NewReader("the body"))

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}*/

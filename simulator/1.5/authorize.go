package simulator

import (
	"encoding/xml"
	"io/ioutil"
	"github.com/RonaldoSantana/ocpp-simulator/soap"
	"text/template"
	"os"
	"log"
	"bytes"
)

// TODO: create an interface that declares the method to parse the response and the request body
// TODO: create a base class that implements the parse of the request body and the parse of the template with a base struct type?
// TODO: the method that makes the call on request.go receives the interface

const statusAccepted = "Accepted"
const statusBlocked = "Blocked"
const statusExpired = "Expired"
const statusInvalid = "Invalid"
const statusConcurrenTx = "ConcurrentTx"

var tpl *template.Template

type XMLParentIdTag struct {
	XMLName xml.Name `xml:"parentIdTag"`
	Value string `xml:",chardata"`
}

type XMLExpiryDate struct {
	XMLName xml.Name `xml:"expiryDate"`
	Value string `xml:",chardata"`
}

type XMLStatus struct {
	XMLName xml.Name `xml:"status"`
	Value string `xml:",chardata"`
}

type XMLIdTagInfo struct {
	XMLName xml.Name `xml:"idTagInfo"`
	Status XMLStatus
	ExpiryDate XMLExpiryDate
	ParentIdTag XMLParentIdTag
}

type XMLAuthorizeResponse struct {
	XMLName xml.Name  `xml:"authorizeResponse"`
	IdTagInfo XMLIdTagInfo
}

type XMLBody struct {
	XMLName  xml.Name `xml:"Body"`
	AuthorizeResponse XMLAuthorizeResponse
}

type Envelope struct {
	XMLName  xml.Name    `xml:"Envelope"`
	Body   XMLBody
}

type Authorize struct {
	Response Envelope
}

// Defines structure to render XML for Authorize request
type AuthTemplateData struct {
	ChargeBoxID string
	AuthID string
}

func (auth Authorize) ResponseStatus() string {
	d := Envelope{}
	xmlContent, _ := ioutil.ReadFile("example.xml")
	err := xml.Unmarshal(xmlContent, &d)
	if err != nil { panic(err) }
	return d.Body.AuthorizeResponse.IdTagInfo.Status.Value
}

// Check if the authorize call to the central system has been accepted
func (auth Authorize) Accepted() bool {
	return auth.ResponseStatus() == statusAccepted
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (auth Authorize) request() {

	var buffer bytes.Buffer
	file, err := os.Open("xml/1.5/Authorize.xml")
	if err != nil {
		log.Fatalln(err)
	}
	buffer.ReadFrom(file)

	err = tpl.Execute(buffer, buddha)
	if err != nil {
		log.Fatalln(err)
	}

	soap := soap.Request{
		Url : "https://ocpp.ron.testcharge.net.nz",
	}



	soap.Call()

}

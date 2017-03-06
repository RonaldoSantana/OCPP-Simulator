package simulator

import (
	"encoding/xml"
	"io/ioutil"
	"github.com/RonaldoSantana/ocpp-simulator/soap"
	"text/template"
	"os"
	"log"
	"bytes"
	"fmt"
)

// TODO: create an interface that declares the method to parse the response and the request body
// TODO: create a base class that implements the parse of the request body and the parse of the template with a base struct type?
// TODO: the method that makes the call on request.go receives the interface

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
	Tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func (auth Authorize) request() {

	var buffer bytes.Buffer
	tplData := AuthTemplateData{
		ChargeBoxID: "veefil-21159",
		AuthID: "B4F62CEF",
	}

	file, err := os.Open("xml/1.5/Authorize.xml")
	if err != nil {
		log.Fatalln(err)
	}
	buffer.ReadFrom(file)

	err = Tpl.ExecuteTemplate(buffer, "xml/1.5/Authorize.xml", tplData)
	if err != nil {
		log.Fatalln(err)
	}

	/*soap := soap.Request{
		Url : "https://ocpp.ron.testcharge.net.nz",
	}*/



	//soap.Call()

}

func (auth Authorize) ParseResponseBody() {
	response := Envelope{}
	xmlContent, _ := ioutil.ReadFile("example.xml")
	err := xml.Unmarshal(xmlContent, &response)
	if err != nil { panic(err) }
	fmt.Println("XMLName:", response.XMLName)
	fmt.Println("Status:", response.Body.AuthorizeResponse.IdTagInfo.Status.Value)
}

func (auth Authorize) ParseRequestBody(requestData AuthTemplateData) {

	var buffer bytes.Buffer
	tplData := AuthTemplateData{
		ChargeBoxID: requestData.ChargeBoxID,
		AuthID: requestData.AuthID,
	}

	err := Tpl.ExecuteTemplate(buffer, auth.Template(), tplData)
	if err != nil {
		log.Fatalln(err)
	}
}

func (auth Authorize) Template() {
	return "xml/Authorize.xml"
}

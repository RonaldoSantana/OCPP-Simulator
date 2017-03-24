package charger

import (
	"encoding/xml"
	"text/template"
	"log"
	"bytes"
)

// TODO: create an interface that declares the method to parse the response and the request body
// TODO: create a base class that implements the parse of the request body and the parse of the template with a base struct type?
// TODO: the method that makes the call on request.go receives the interface

type XMLAuthorizeResponse struct {
	XMLName xml.Name  `xml:"authorizeResponse"`
	IdTagInfo XMLIdTagInfo
}

type XMLAuthorizeBody struct {
	XMLName  xml.Name `xml:"Body"`
	AuthorizeResponse XMLAuthorizeResponse
}

type EnvelopeAuthorize struct {
	XMLName  xml.Name    `xml:"Envelope"`
	Body   XMLAuthorizeBody
}

type Authorize struct {
	Response *EnvelopeAuthorize
}

// parses the XML - adding values to parameters, etc.
func (auth *Authorize) ParseRequestBody(data []string) string {

	// TODO: validate number of arguments

	var buffer bytes.Buffer
	tpl := template.Must(template.ParseFiles(auth.Template()))

	// template data
	tplData := RequestData{
		ChargeBoxID: data[1],
		AuthID: data[2],
	}

	err := tpl.Execute(&buffer, tplData)
	if err != nil {
		log.Fatalln(err)
	}

	return buffer.String()
}

// Parse response
func (auth *Authorize) ParseResponseBody(responseData []byte) {
	err := xml.Unmarshal(responseData, &auth.Response)
	if err != nil {
		log.Fatalln(err)
	}

	//simulator.Response = *response;
}

// Gets the XML to be used for this request
func (auth *Authorize) Template() string {
	return "xml/Authorize.xml"
}

// Gets the response status for the Authorize request
func (auth *Authorize) ResponseStatus() string {
	return auth.Response.Body.AuthorizeResponse.IdTagInfo.Status.Value
}

// Check if the authorize call to the central system has been accepted
func (auth *Authorize) Accepted() bool {
	return auth.ResponseStatus() == StatusAccepted
}


// Gets the response status for the Authorize request
func (auth *Authorize) ValidateArguments(data []string) {
	// TODO
}

func NewAuthorize() *Authorize {
	return &Authorize{}
}
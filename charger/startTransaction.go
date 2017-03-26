package charger

import (
	"encoding/xml"
	"text/template"
	"log"
	"bytes"
	"time"
	"github.com/satori/go.uuid"
)

type XMLStartTransactionResponse struct {
	XMLName xml.Name  `xml:"startTransactionResponse"`
	TransactionId XMLTransactionId
	IdTagInfo XMLIdTagInfo
}

type XMLStartTransactionBody struct {
	XMLName  xml.Name `xml:"Body"`
	StartTransactionResponse XMLStartTransactionResponse
}

type EnvelopeStartTransaction struct {
	XMLName  xml.Name    `xml:"Envelope"`
	Body   XMLStartTransactionBody
}

// Defines structure to render XML for Authorize request
type StartTransactionData struct {
	RequestData
	MessageID string
	DateTimeStart string
}

type StartTransaction struct {
	Response *EnvelopeStartTransaction
}

// parses the XML - adding values to parameters, etc.
func (auth *StartTransaction) ParseRequestBody(data []string) string {

	// TODO: validate number of arguments

	var buffer bytes.Buffer
	tpl := template.Must(template.ParseFiles(auth.Template()))

	// date and time we are starting the transaction
	t1 := time.Now()

	// template data
	tplData := StartTransactionData{
		RequestData{
			data[0],
			data[1],
		},
		uuid.NewV4().String(),
		t1.Format(time.RFC3339),
	}

	err := tpl.Execute(&buffer, tplData)
	if err != nil {
		log.Fatalln(err)
	}

	return buffer.String()
}

// Parse response
func (auth *StartTransaction) ParseResponseBody(responseData []byte) {
	err := xml.Unmarshal(responseData, &auth.Response)
	if err != nil {
		log.Fatalln(err)
	}

	//simulator.Response = *response;
}

// Gets the XML to be used for this request
func (auth *StartTransaction) Template() string {
	return "xml/StartTransaction.xml"
}

// Gets the response status for the Authorize request
func (auth *StartTransaction) ResponseStatus() string {
	return auth.Response.Body.StartTransactionResponse.IdTagInfo.Status.Value + " : TransactionID: " + auth.Response.Body.StartTransactionResponse.TransactionId.Value
}

// Check if the authorize call to the central system has been accepted
func (auth *StartTransaction) Accepted() bool {
	return auth.ResponseStatus() == StatusAccepted
}

// Gets the response status for the Authorize request
func (auth *StartTransaction) ValidateArguments(data []string) {
	// TODO
}

func NewStartTransaction() *StartTransaction {
	return &StartTransaction{}
}
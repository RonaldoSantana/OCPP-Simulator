package soap

import (
	"net/http"
	"bytes"
	"log"
	"io/ioutil"
)

type Request struct {
	Url string
}

// Here is the function querying the SOAP server
// It returns the whole response as a string - up to the caller to create the structure
func (request Request) Call(soapRequestContent string) string {

	httpClient := new(http.Client)

	// make request to central system
	response, err := httpClient.Post(request.Url, "text/xml; charset=utf-8", bytes.NewBufferString(soapRequestContent))
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// the response XML are small, we can just use ReadAll
	xmlBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(xmlBody)
}


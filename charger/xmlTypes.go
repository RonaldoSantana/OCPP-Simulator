package charger

import (
	"encoding/xml"
)

// IdTagInfo complex type on wsdl
type XMLIdTagInfo struct {
	XMLName xml.Name `xml:"idTagInfo"`
	Status XMLStatus
	ExpiryDate XMLExpiryDate
	ParentIdTag XMLParentIdTag
}

type XMLStatus struct {
	XMLName xml.Name `xml:"status"`
	Value string `xml:",chardata"`
}

type XMLExpiryDate struct {
	XMLName xml.Name `xml:"expiryDate"`
	Value string `xml:",chardata"`
}

type XMLParentIdTag struct {
	XMLName xml.Name `xml:"parentIdTag"`
	Value string `xml:",chardata"`
}

// transactionId element on diverse types on wsdl
type XMLTransactionId struct {
	XMLName xml.Name `xml:"transactionId"`
	Value string `xml:",chardata"`
}
package charger

const (
	StatusAccepted = "Accepted"
	StatusBlocked = "Blocked"
	StatusExpired = "Expired"
	StatusInvalid = "Invalid"
	StatusConcurrenTx = "ConcurrentTx"
)

// Basic request parameters
type RequestData struct {
	ChargeBoxID string
	AuthID string
}

// empty interface so every request has it' own definition

// Interface that all charge point Request types needs to implement
type ChargePointMethod interface {
	ParseRequestBody(data []string) string // the request XML to be posted to central system
	ParseResponseBody(responseData []byte) // the parsed response, according to request
	ResponseStatus() string
}
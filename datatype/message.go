package datatype

//Message implements the sentry interface that a message is in a event request
type Message struct {
	Message   *string   `json:"message,omitempty"`
	Formatted *string   `json:"formatted,omitempty"`
	Params    *[]string `json:"params,omitempty"`
}

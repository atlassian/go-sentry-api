package datatype

//Query implements the sentry interface for a query
type Query struct {
	Query  *string `json:"query,omitempty"`
	Engine *string `json:"engine,omitempty"`
}

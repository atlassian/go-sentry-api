package datatype

//Request implements the sentry interface for a request error
type Request struct {
	Cookies  *[][2]string       `json:"cookies,omitempty"`
	Fragment *string            `json:"fragment,omitempty"`
	Headers  *[][2]string       `json:"headers,omitempty"`
	URL      *string            `json:"uRL,omitempty"`
	Env      *map[string]string `json:"env,omitempty"`
	Query    *string            `json:"query,omitempty"`
	Data     *string            `json:"data,omitempty"`
	Method   *string            `json:"method,omitempty"`
}

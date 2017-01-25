package sentry

// QueryReq is a simple internal interface
type QueryReq interface {
	ToQueryString() string
}

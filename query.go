package sentry

type SentryQueryReq interface {
	ToQueryString() string
}

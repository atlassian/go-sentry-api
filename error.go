package sentry

import (
	"fmt"
)

type SentryApiError struct {
	Detail     string `json:detail`
	StatusCode int    `json:-`
}

func (s SentryApiError) Error() string {
	if s.StatusCode == 404 {
		return "404: Endpoint/Resource not found"
	}
	return fmt.Sprintf("%d: %s", s.StatusCode, s.Detail)
}

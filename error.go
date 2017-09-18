package sentry

import (
	"fmt"
)

// APIError is used when the api returns back a non 200 response
type APIError struct {
	Detail     string `json:"detail,omitempty"`
	StatusCode int    `json:"-"`
}

// Error is the interface needed to transfor to a error type
func (s APIError) Error() string {
	if s.StatusCode == 404 {
		return "404: Endpoint/Resource not found"
	}
	if s.StatusCode == 400 && s.Detail == "" {
		return fmt.Sprintf("400: Bad Request. Invalid data is probably the issue")
	}
	return fmt.Sprintf("%d: %s", s.StatusCode, s.Detail)
}

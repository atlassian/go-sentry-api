package sentry

import (
	"fmt"
)

// APIError is used when the api returns back a non 200 response
type APIError struct {
	Detail     string `json:"detail"`
	StatusCode int    `json:"-"`
}

// Error is the interface needed to transfor to a error type
func (s APIError) Error() string {
	if s.StatusCode == 404 {
		return "404: Endpoint/Resource not found"
	}
	return fmt.Sprintf("%d: %s", s.StatusCode, s.Detail)
}

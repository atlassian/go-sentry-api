package sentry

import (
	"os"
	"testing"
)

var authtoken = os.Getenv("SENTRY_AUTH_TOKEN")
var endpoint = os.Getenv("SENTRY_ENDPOINT")
var client, clienterr = NewClient(authtoken, &endpoint, nil)

func TestClientBadEndpoint(t *testing.T) {
	t.Parallel()
	badendpoint := ""
	_, berr := NewClient(authtoken, &badendpoint, nil)
	if berr == nil {
		t.Error("Should have gotten an error for an empty endpoint")
	}
}

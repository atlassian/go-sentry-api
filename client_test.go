package sentry

import (
	"os"
	"testing"
)

var authtoken = os.Getenv("SENTRY_AUTH_TOKEN")
var endpoint = os.Getenv("SENTRY_ENDPOINT")
var defaultorg = os.Getenv("SENTRY_DEFAULT_ORG")
var client, clienterr = NewClient(authtoken, &endpoint, nil)

func getDefaultOrg() string {
	if defaultorg == "" {
		return "sentry"
	}

	return defaultorg
}

func TestClientBadEndpoint(t *testing.T) {
	t.Parallel()
	badendpoint := ""
	_, berr := NewClient(authtoken, &badendpoint, nil)
	if berr == nil {
		t.Error("Should have gotten an error for an empty endpoint")
	}
}

func TestClientKnownGoodEndpoint(t *testing.T) {
	bclient, berr := NewClient("testauthclient", nil, nil)
	if berr != nil {
		t.Error(berr)
	}
	if bclient.Endpoint != "https://sentry.io/api/0/" {
		t.Errorf("Endpoint is not https://sentry.io/api/0 got %s", bclient.Endpoint)
	}
}

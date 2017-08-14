package sentry

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
)

var (
	authtoken         = os.Getenv("SENTRY_AUTH_TOKEN")
	endpoint          = os.Getenv("SENTRY_ENDPOINT")
	defaultorg        = os.Getenv("SENTRY_DEFAULT_ORG")
	client, clienterr = NewClient(authtoken, &endpoint, nil)
)

func getDefaultOrg() string {
	if defaultorg == "" {
		return "sentry"
	}

	if authtoken == "" && endpoint == "" {
		log.Fatalf("Failed to setup tests. Please have SENTRY_AUTH_TOKEN and SENTRY_ENDPOINT set")
	}

	log.Printf("Using sentry slug organization %s", defaultorg)

	return defaultorg
}

func generateIdentifier(prefix string) string {
	return fmt.Sprintf("Test %s for go-sentry-api-%d", prefix, rand.Int())
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

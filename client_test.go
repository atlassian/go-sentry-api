package sentry

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getDefaultOrg() string {
	defaultorg := os.Getenv("SENTRY_DEFAULT_ORG")

	if defaultorg == "" {
		return "sentry"
	}

	return defaultorg
}

func newTestClient(t *testing.T) *Client {
	endpoint := os.Getenv("SENTRY_ENDPOINT")
	if endpoint == "" {
		endpoint = "https://sentry.io/api/0/"
	}

	authtoken := os.Getenv("SENTRY_AUTH_TOKEN")
	require.NotEmpty(t, authtoken, "Need SENTRY_AUTH_TOKEN env var to continue")

	client, clienterr := NewClient(authtoken, &endpoint, nil)
	require.NoError(t, clienterr)
	require.NotNil(t, client)

	return client
}

func generateIdentifier(prefix string) string {
	return fmt.Sprintf("%s %d", prefix, rand.Int())
}

func TestClientBadEndpoint(t *testing.T) {
	t.Parallel()
	badendpoint := ""

	authtoken := os.Getenv("SENTRY_AUTH_TOKEN")

	_, berr := NewClient(authtoken, &badendpoint, nil)
	assert.Error(t, berr)
}

func TestClientKnownGoodEndpoint(t *testing.T) {
	bclient, berr := NewClient("testauthclient", nil, nil)
	if berr != nil {
		t.Error(berr)
	}
	if bclient.endPoint != "https://sentry.io/api/0/" {
		t.Errorf("Endpoint is not https://sentry.io/api/0 got %s", bclient.endPoint)
	}
}

func TestNewRequestWillNotAddExtraTrailingSlashToEndpoint(t *testing.T) {
	endpoint := "some-endpoint/"
	bclient, berr := NewClient("testauthclient", nil, nil)
	if berr != nil {
		t.Error(berr)
	}
	req, err := bclient.newRequest("get", endpoint, nil)
	if req == nil || err != nil {
		t.Errorf("can't generate request: %v", err)
	}

	if req.URL.String() != "https://sentry.io/api/0/some-endpoint/" {
		t.Errorf("Endpoint is not https://sentry.io/api/0/some-endpoint/ got %s", req.URL.String())
	}
}

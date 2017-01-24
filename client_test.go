package sentry

import (
	"os"
)

var authtoken = os.Getenv("SENTRY_AUTH_TOKEN")
var client = NewClient(authtoken, nil, nil)

module github.com/atlassian/go-sentry-api

go 1.13

require (
	github.com/getsentry/sentry-go v0.9.0
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace (
	// The following dependencies contain vulnerabilities in the versions that
	// dependencies are trying to import them as. The alternative to this is to
	// add these libraries as direct dependencies, but to stop `go mod tidy`
	// removing them we would need to import them somewhere in the code.
	// See: https://github.com/golang/go/issues/37352
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/labstack/echo/v4 => github.com/labstack/echo/v4 v4.2.0
	golang.org/x/text => golang.org/x/text v0.3.5
)

.PHONY: test

test:
	go fmt *.go
	SENTRY_ENDPOINT=http://localhost:8080/api/0/ go test -race ./...

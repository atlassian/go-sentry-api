STUPIDSECRET=thisisadumbsecretyo

.PHONY: all

all:
	test

test:
	go fmt *.go
	SENTRY_ENDPOINT=http://localhost:8080/api/0/ go test $$(glide novendor) -v

devenv:
	docker run -d --name sentry-redis redis
	docker run -d --name sentry-postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=sentry postgres
	sleep 5 # Wait for postgres to bootup
	docker run -it --rm -e SENTRY_SECRET_KEY='${STUPIDSECRET}' --link sentry-postgres:postgres --link sentry-redis:redis sentry:latest upgrade
	docker run -d --name my-sentry -e SENTRY_SECRET_KEY='${STUPIDSECRET}' --link sentry-redis:redis --link sentry-postgres:postgres -p 8080:9000 sentry:latest
	docker run -d --name sentry-cron -e SENTRY_SECRET_KEY='${STUPIDSECRET}' --link sentry-postgres:postgres --link sentry-redis:redis sentry:latest run cron
	docker run -d --name sentry-worker-1 -e SENTRY_SECRET_KEY='${STUPIDSECRET}' --link sentry-postgres:postgres --link sentry-redis:redis sentry:latest run worker

devclean:
	docker kill $$(docker ps -q -a --no-trunc --filter name=^sentry)
	docker rm $$(docker ps -q -a --no-trunc --filter name=^sentry)

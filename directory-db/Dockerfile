# Using stretch here since build artifact should work on ubuntu image
FROM golang:1-stretch AS build

RUN apt-get update && \
	apt-get -y install --no-install-recommends git && \
	rm -rf /var/lib/apt/lists/*

# Build migrate with postgres support
# TODO: Consider moving to single-stage docker build and obtain migrate from ppa.
# 	See https://github.com/golang-migrate/migrate/tree/master/cli
RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq && \
	go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

FROM ubuntu:artful

# Install psql
RUN apt-get update && \
	apt-get -y install --no-install-recommends wget ca-certificates && \
	wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
	echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" > /etc/apt/sources.list.d/postgres.list && \
	apt update && \
	apt -y install --no-install-recommends postgresql-client-9.6 && \
	rm -rf /var/lib/apt/lists/*

COPY --from=build /usr/local/bin/migrate	/usr/local/bin/migrate
COPY directory-db/docker/reset-db.sh		/usr/local/bin/reset-db.sh
COPY directory-db/docker/upgrade-db.sh		/usr/local/bin/upgrade-db.sh
COPY directory-db/migrations	/db-migrations
COPY directory-db/testdata		/db-testdata

FROM alpine:3.16.0 AS build
ARG MIGRATE_VERSION="4.15.1"
ARG MIGATE_CHECKSUM="17e997ed5fe08d54b53ec0d7f364715e14bfb90566aa6455b51ba8f88a039bda"

RUN wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz && \
    echo "${MIGATE_CHECKSUM}  /tmp/migrate.linux-amd64.tar.gz" | sha256sum -c && \
    tar zxf /tmp/migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    chmod 0755 /usr/local/bin/migrate && \
    /usr/local/bin/migrate -version

FROM alpine:3.16.0
RUN apk add --update bash postgresql-client

COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate

COPY txlog-db/docker/reset-db.sh /usr/local/bin/reset-db.sh
COPY txlog-db/docker/upgrade-db.sh /usr/local/bin/upgrade-db.sh
COPY txlog-db/migrations /db-migrations
COPY txlog-db/testdata /db-testdata

RUN adduser -D -u 1001 appuser

RUN chown appuser /db-migrations
RUN chown appuser /db-testdata

USER appuser

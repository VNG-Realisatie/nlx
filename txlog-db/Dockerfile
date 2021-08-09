FROM alpine:20210804 AS build
ARG MIGRATE_VERSION="4.7.0"
ARG MIGATE_CHECKSUM="976e351a7f5afb5f13d2020123c450cb56bb66e6a0542efe563a89dbabc3fbdc"

RUN wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz && \
    tar zxf /tmp/migrate.linux-amd64.tar.gz && \
    echo "${MIGATE_CHECKSUM}  migrate.linux-amd64" | sha256sum -c && \
    mv migrate.linux-amd64 /usr/local/bin/migrate && \
    chmod 0755 /usr/local/bin/migrate && \
    /usr/local/bin/migrate -version


FROM alpine:20210804
COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate

RUN apk --no-cache add bash postgresql-client

COPY txlog-db/docker/reset-db.sh /usr/local/bin/reset-db.sh
COPY txlog-db/docker/upgrade-db.sh /usr/local/bin/upgrade-db.sh
COPY txlog-db/migrations /db-migrations
COPY txlog-db/testdata /db-testdata

RUN adduser -D -u 1001 appuser

RUN chown appuser /db-migrations
RUN chown appuser /db-testdata

USER appuser

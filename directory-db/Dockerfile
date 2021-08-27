FROM alpine:3.14.2 AS build
ARG MIGRATE_VERSION="4.14.1"
ARG MIGATE_CHECKSUM="faaf3c4c96b6a5fb55a1b29888deb473aacbb47a858dafa7e586e1587c968265"

RUN wget -O /tmp/migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz && \
    tar zxf /tmp/migrate.linux-amd64.tar.gz && \
    echo "${MIGATE_CHECKSUM}  migrate.linux-amd64" | sha256sum -c && \
    mv migrate.linux-amd64 /usr/local/bin/migrate && \
    chmod 0755 /usr/local/bin/migrate && \
    /usr/local/bin/migrate -version


FROM alpine:3.14.2
COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate

RUN apk --no-cache add bash postgresql-client

COPY directory-db/docker/reset-db.sh /usr/local/bin/reset-db.sh
COPY directory-db/docker/upgrade-db.sh /usr/local/bin/upgrade-db.sh
COPY directory-db/migrations /db-migrations
COPY directory-db/testdata /db-testdata

RUN adduser -D -u 1001 appuser

RUN chown appuser /db-migrations
RUN chown appuser /db-testdata

USER appuser

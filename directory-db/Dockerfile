FROM ubuntu:bionic

# Install psql
RUN apt-get update && \
	apt-get -y install --no-install-recommends wget ca-certificates gnupg && \
	wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - && \
	echo "deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main" > /etc/apt/sources.list.d/postgres.list && \
	apt update && \
	apt -y install --no-install-recommends postgresql-client-9.6 && \
	rm -rf /var/lib/apt/lists/*

# Install migrate
RUN wget --quiet -O - https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && \
    apt-get install -y migrate && \
    rm -rf /var/lib/apt/lists/*

COPY directory-db/docker/reset-db.sh		/usr/local/bin/reset-db.sh
COPY directory-db/docker/upgrade-db.sh		/usr/local/bin/upgrade-db.sh
COPY directory-db/migrations	/db-migrations
COPY directory-db/testdata		/db-testdata

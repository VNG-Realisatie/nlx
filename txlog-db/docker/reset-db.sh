#!/bin/bash
#
# Reset the database, run migrations and load test data
#
# Usage: ./reset-db.sh
#
# The following environment variables should be set:
#
#   PGHOST
#     Hostname of the database cluster/server. Unix socket directories are supported
#     as well. Defaults to /tmp for unix sockets.
#
#   PGPORT
#     Port number of the database cluster/server. Defaults to 5432.
#
#   PGDATABASE
#     Name of the transaction log database. Will be created if it doesn't exist yet.
#
#   PGUSER:
#     The administrative database user managing the schema of the transaction
#     log database
#
#   PGPASSWORD:
#     Password of the administrative database user account.
#
#   PGSSLMODE (optional):
#     This option determines whether or with what priority a secure SSL TCP/IP
#     connection will be negotiated with the server.
#
#		PGCONNECT_TIMEOUT (optional):
#			Maximum wait for connection, in seconds (write as a decimal integer string).
#			Zero or not specified means wait indefinitely. It is not recommended to
#			use a` timeout of less than 2 seconds.

set -e
set -o pipefail
set -u  # fail on undefined (env) vars

export PGCONNECT_TIMEOUT=${PGCONNECT_TIMEOUT:-5}

echo "Drop existing database ${PGDATABASE}"
# We specify a low connect_timeout so a db pod/container doesn't hang a long time when the postgres is not created yet. Failing and restarting is faster.
psql --echo-errors --variable "ON_ERROR_STOP=1" "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/postgres" <<EOF
UPDATE pg_database SET datallowconn = false WHERE datname = '${PGDATABASE}';
SELECT pid, datname, usename, application_name, pg_terminate_backend(pid) AS terminated
	FROM pg_stat_activity
	WHERE datname = '${PGDATABASE}'
		AND pid <> pg_backend_pid();
DROP DATABASE IF EXISTS "${PGDATABASE}";
CREATE DATABASE "${PGDATABASE}";
EOF

echo "Creating database structure from migrations and adding testdata"
dbVersion=0
for _ in $(find /db-migrations -name "*.up.sql" -print0 | sort -z | xargs -r0 echo); do
	migrate --path /db-migrations/ --database "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/${PGDATABASE}" up 1
	(( dbVersion=dbVersion+1 ))
	dbVersionZerofill=$(printf "%03d" "${dbVersion}")
	# Add test data
	for dataFile in $(find /db-testdata -name "${dbVersionZerofill}_*.sql" -print0 | sort -z | xargs -r0 echo); do
		psql --echo-errors --variable "ON_ERROR_STOP=1" "${PGDATABASE}" < "${dataFile}" | awk "\$0=\"[${dataFile/.sql/}] \"\$0"
	done
done

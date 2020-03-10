#!/bin/bash
#
# Migrate the database for the directory to its latest state.
#
# Usage: ./upgrade-db.sh
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

set -e
set -o pipefail
set -u  # fail on undefined (env) vars

# Export database connection variables, and set sane defaults (same defaults as psql)
export PGHOST=${PGHOST:-/tmp}
export PGPORT=${PGPORT:-5432}

# 1. Ensure the database exists. Either the database must be provisioned up-front,
#    or $PGUSER must have CREATEDB privileges.
psql --echo-errors --variable "ON_ERROR_STOP=1" "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/postgres?sslmode=disable&connect_timeout=5" <<EOF
    SELECT 'CREATE DATABASE "${PGDATABASE}"' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '${PGDATABASE}')\gexec
EOF

# 2. Perform the actual migrations with the 'templated out' SQL.
# escape slashes in case we're dealing with unix sockets
host=$(echo $PGHOST | sed "s:/:%2F:g")
/usr/local/bin/migrate \
    --database "postgres://${PGUSER}:${PGPASSWORD}@:${PGPORT}/${PGDATABASE}?host=${host}&sslmode=disable&connect_timeout=5" \
    --lock-timeout 600 \
    --path /db-migrations/ \
    up

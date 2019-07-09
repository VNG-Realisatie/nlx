#!/bin/bash
set -e
set -o pipefail

psql --echo-errors --variable "ON_ERROR_STOP=1" "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/postgres?sslmode=disable&connect_timeout=5" <<EOF
    SELECT 'CREATE DATABASE "${PGDATABASE}"' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '${PGDATABASE}')\gexec
EOF

sed -i "s/nlx-config/${PGDATABASE}/g" /db-migrations/*.up.sql

/usr/local/bin/migrate \
    --database "postgres://${PGUSER}:${PGPASSWORD}@${PGHOST}:5432/${PGDATABASE}?sslmode=disable&connect_timeout=5" \
    --lock-timeout 600 \
    --path /db-migrations/ \
    up

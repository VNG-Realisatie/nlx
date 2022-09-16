#!/bin/bash
# Copyright © VNG Realisatie 2022
# Licensed under the EUPL

# Run this script from the root folder of the git repository with the following command: ./scripts/start-development.sh

# Make sure permissions for pki files are ok
sh ./pki/fix-permissions.sh

# Start docker-compose
# Note: the --build flag is needed to rebuild the auth-opa containers,
# since changes to the content of the files do not trigger a rebuild automatically
if ! docker-compose -f docker-compose.dev.yml up -d --remove-orphans --build; then
    echo "Error while starting docker-compose, exiting now"
    exit
fi

# Wait for postgres to accept connections
until docker-compose -f docker-compose.dev.yml exec postgres pg_isready
do
    sleep 1;
done;

# Migrate txlog databases
go run ./txlog-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_a?sslmode=disable"
go run ./txlog-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_b?sslmode=disable"
go run ./txlog-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_c?sslmode=disable"

# Migrate directory database
go run ./directory-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx?sslmode=disable"

# Migrate nlx_management databases
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_c?sslmode=disable"

# Create admin users
go run ./management-api create-user --email admin@nlx.local --role admin \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api create-user --email admin@nlx.local --password development --role admin \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
go run ./management-api create-user --email admin@nlx.local --password development --role admin \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_c?sslmode=disable"

# Create readonly users
go run ./management-api create-user --email readonly@nlx.local --role readonly \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api create-user --email readonly@nlx.local --password development --role readonly \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
go run ./management-api create-user --email readonly@nlx.local --password development --role readonly \
    --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_c?sslmode=disable"

# start services
TXLOG_A=1 TXLOG_B=1 TXLOG_C=1 modd

function finish {
  docker-compose -f docker-compose.dev.yml down --remove-orphans
}

trap finish EXIT

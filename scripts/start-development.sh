#!/bin/bash
# Run this script from the root folder of the git repository with the following command: ./scripts/start-development.sh

# Start docker-compose
docker-compose -f docker-compose.dev.yml up -d

# Wait for postgres to accept connections
until docker-compose -f docker-compose.dev.yml exec postgres pg_isready
do
    sleep 1;
done;

# Migrate databases
docker run --rm -v $(pwd)/txlog-db/migrations:/migrations     --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_a?sslmode=disable up
docker run --rm -v $(pwd)/txlog-db/migrations:/migrations     --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_b?sslmode=disable up
docker run --rm -v $(pwd)/directory-db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@127.0.0.1:5432/nlx?sslmode=disable up

# Migrate nlx_management databases
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"

# Create admin users
go run ./management-api create-user --email admin@nlx.local \
    --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api create-user --email admin@nlx.local --password development \
    --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"

# start services
TXLOG_A=1 TXLOG_B=1 modd

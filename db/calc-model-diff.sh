#!/bin/bash

set -e # exit on error

tmpdir=`mktemp -d -t nlx_db_diff-XXXXXXXX`

if docker inspect $(hostname) &> /dev/null; then
    dockerNetwork="--network $(docker inspect -f '{{range .NetworkSettings.Networks}}{{.NetworkID}}{{end}}' $(hostname))"
fi

echo "Starting postgres containers"
function cleanupDockerContainers {
  docker kill nlx_diff_migrate &> /dev/null || true
  docker rm nlx_diff_migrate &> /dev/null || true
  docker kill nlx_diff_target &> /dev/null || true
  docker rm nlx_diff_target &> /dev/null || true
}
cleanupDockerContainers
pgTargetContainer=`docker run --name nlx_diff_target ${dockerNetwork} -e POSTGRES_PASSWORD=postgres -d postgres:9.6`
pgMigrateContainer=`docker run --name nlx_diff_migrate ${dockerNetwork} -e POSTGRES_PASSWORD=postgres -d postgres:9.6`
trap cleanupDockerContainers EXIT
export PGPASSWORD=postgres

echo "Creating database strucutre from pgmodeler generated sql"
pgTargetIP=`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${pgTargetContainer}`
until (psql --user postgres --host ${pgTargetIP} --command '\l'&>/dev/null); do
  echo "Waiting for postgres in container nlx_diff_target"
  sleep 1
done
psql --echo-errors --variable "ON_ERROR_STOP=1" --user postgres --host ${pgTargetIP} --command "CREATE DATABASE nlx;" postgres >/dev/null
psql --echo-errors --variable "ON_ERROR_STOP=1" --user postgres --host ${pgTargetIP} --file "model/nlx.sql" nlx >/dev/null
pg_dump --user postgres --host ${pgTargetIP} --schema-only --file ${tmpdir}/target.sql nlx

echo "Creating database structure from migration files"
pgMigrateIP=`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${pgMigrateContainer}`
until (psql --user postgres --host ${pgMigrateIP} --command '\l' &>/dev/null); do
  echo "Waiting for postgres in container nlx_diff_migrate"
  sleep 1
done
# Create database and apply all migrations
psql --echo-errors --variable "ON_ERROR_STOP=1" --user postgres --host ${pgMigrateIP} --command "CREATE DATABASE nlx;" postgres >/dev/null
migrateDSN="postgres://postgres:postgres@${pgMigrateIP}:5432/nlx?sslmode=disable"
migrate --database ${migrateDSN} --path migrations/ up
migrateLastVersion=$(migrate --database ${migrateDSN} --path migrations/ version 2>&1)
# Dump migrated version
pg_dump --user postgres --host ${pgMigrateIP} --schema-only --exclude-table schema_migrations --file ${tmpdir}/migrate.sql nlx
# Reset database and apply migrations except the last migration
psql --echo-errors --variable "ON_ERROR_STOP=1" --user postgres --host ${pgMigrateIP} --command "DROP DATABASE IF EXISTS nlx;" postgres >/dev/null
psql --echo-errors --variable "ON_ERROR_STOP=1" --user postgres --host ${pgMigrateIP} --command "CREATE DATABASE nlx;" postgres >/dev/null
migrate --database ${migrateDSN} --path migrations/ up $(expr ${migrateLastVersion} - 1)
# Dump migrate-previous version
pg_dump --user postgres --host ${pgMigrateIP} --schema-only --exclude-table schema_migrations --file ${tmpdir}/migrate-previous.sql nlx
# Apply last migration up and down
migrate --database ${migrateDSN} --path migrations/ up 1
migrate --database ${migrateDSN} --path migrations/ down 1
# Dump migrate-rollback
pg_dump --user postgres --host ${pgMigrateIP} --schema-only --exclude-table schema_migrations --file ${tmpdir}/migrate-rollback.sql nlx

echo "Finding diffs for up"
diff=$(apgdiff --ignore-function-whitespace ${tmpdir}/migrate.sql ${tmpdir}/target.sql)
if [[ $? != 0 ]]; then
  echo "Diff up failed"
elif [[ $diff ]]; then
  echo "${diff}"
  exit 64
else
  echo "No differences found for up"
fi

echo "Finding diffs for down"
diff=$(apgdiff --ignore-function-whitespace ${tmpdir}/migrate-rollback.sql ${tmpdir}/migrate-previous.sql)
if [[ $? != 0 ]]; then
  echo "Diff down failed"
elif [[ $diff ]]; then
  echo "${diff}"
  exit 64
else
  echo "No differences found for down"
fi

if ! grep -q "^const LatestVersion = ${migrateLastVersion}\$" dbversion/version.go; then
    echo "dbversion/version.go is invalid, expected LatestVersion to equal ${migrateLastVersion}, got different value.";
    exit 64
fi

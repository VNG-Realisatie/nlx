#!/bin/bash
set -e
set -o pipefail

echo "Drop existing database ${PGDATABASE}"
# We specify a low connect_timeout so a db pod/container doesn't hang a long time when the postgres is not created yet. Failing and restarting is faster.
psql --echo-errors --variable "ON_ERROR_STOP=1" "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/postgres?sslmode=disable&connect_timeout=5" <<EOF
UPDATE pg_database SET datallowconn = false WHERE datname = '${PGDATABASE}';
SELECT pid, datname, usename, application_name, pg_terminate_backend(pid) AS terminated
	FROM pg_stat_activity
	WHERE datname = '${PGDATABASE}'
		AND pid <> pg_backend_pid();
DROP DATABASE IF EXISTS "${PGDATABASE}";
CREATE DATABASE "${PGDATABASE}";
EOF

sed -i "s/nlx-org-txlog-/nlx-${PGDATABASE}-/g" /db-migrations/*.up.sql

echo "Creating database structure from migrations and adding testdata"
dbVersion=0
for _unused in $(find /db-migrations -name "*.up.sql" -print0 | sort -z | xargs -r0 echo); do
	migrate --path /db-migrations/ --database "postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}/${PGDATABASE}?sslmode=disable&connect_timeout=5" up 1
	let dbVersion=dbVersion+1
	dbVersionZerofill=$(printf "%03d" ${dbVersion})
	# Add test data
	for dataFile in $(find /db-testdata -name "${dbVersionZerofill}_*.sql" -print0 | sort -z | xargs -r0 echo); do
		psql --echo-errors --variable "ON_ERROR_STOP=1" ${PGDATABASE} < ${dataFile} | awk "\$0=\"[${dataFile/.sql/}] \"\$0"
	done
done

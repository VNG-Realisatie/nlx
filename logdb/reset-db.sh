#!/bin/bash
set -e
set -o pipefail

echo "Drop existing database"
psql --echo-errors --variable "ON_ERROR_STOP=1" postgres <<EOF
SELECT pid, datname, usename, pg_terminate_backend(pg_stat_activity.pid) AS terminated
	FROM pg_stat_activity
	WHERE pg_stat_activity.datname = 'nlx_logdb'
		AND pid <> pg_backend_pid();
DROP DATABASE IF EXISTS nlx_logdb;
CREATE DATABASE nlx_logdb;
EOF

echo "Creating database structure from migrations and adding testdata"
dbVersion=0
while migrate --path migrations/ --database "postgresql://postgres:postgres@postgres/nlx_logdb?sslmode=disable" up 1;
do
	let dbVersion=dbVersion+1
	dbVersionZerofill=$(printf "%03d" ${dbVersion})
	# Add test data
	for dataFile in $(find testdata -name "${dbVersionZerofill}_*.sql" -print0 | sort -z | xargs -r0 echo); do
		psql --echo-errors --variable "ON_ERROR_STOP=1" nlx_logdb < ${dataFile} | awk "\$0=\"[${dataFile/.sql/}] \"\$0"
	done
done

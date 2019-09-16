#!/usr/bin/env bash

set -u
set -e
set -x

# wait for postgres
while ! nc -z postgres 5432
do
	echo "Waiting for postgres..."
	sleep 2
done

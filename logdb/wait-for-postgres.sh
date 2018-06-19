#!/bin/bash
set -e

cmd="$*"

until (psql -c '\l' &>/dev/null); do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd

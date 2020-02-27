#!/bin/sh

set -e

(( c=1 ))

until /usr/bin/curl -o /dev/null --fail -s $1; do
  >&2 echo "$1 unavailable - sleeping"
  sleep 2
  ((c++)) && ((c==100)) && exit 1
done

>&2 echo "$1 is up!"

exit 0

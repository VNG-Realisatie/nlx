#!/usr/bin/env bash
#shellcheck disable=SC2034 disable=SC1091

set -u
set -e
set -x

NLXROOT=$(git rev-parse --show-toplevel)

source dc.sh

echo "Starting postgres containers"
./docker-cleanup.sh
dc build
dc up -d postgres

./docker-wait.sh

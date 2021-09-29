#!/usr/bin/env bash
#shellcheck disable=SC2034 disable=SC1091

set -u
set -e
set -x

NLXROOT=$(git rev-parse --show-toplevel)

source dc.sh

function cleanupDockerContainers {
  dc down || true
  dc rm --force || true
}
cleanupDockerContainers

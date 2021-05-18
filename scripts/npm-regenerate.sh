#!/usr/bin/env bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
ROOT_DIR="${BASE_DIR}/.."

DIRS="
directory-ui
directory-ui-e2e-tests
docs/website
management-ui
management-ui-e2e-tests"

clean-install() {
  cd "${1}"

  echo "Cleaning..."
  rm -rf "node_modules/" "package-lock.json"

  echo "Installing..."
  npm install

  echo ""
}

echo "Regenerate for root..."
clean-install "${ROOT_DIR}"

for DIRECTORY in $DIRS; do
  echo "Regenerate for '${DIRECTORY}'..."
  clean-install "${ROOT_DIR}/${DIRECTORY}"
done

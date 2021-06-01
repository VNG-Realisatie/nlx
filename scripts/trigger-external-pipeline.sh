#!/usr/bin/env bash

PROJECT_ID="${1}"
VERSION="${2}"
VERSION_VAR_NAME="${3}"
TOKEN="${4}"

URL="https://gitlab.com/api/v4/projects/${PROJECT_ID}/trigger/pipeline"
OUTPUT_FILE=$(mktemp -t curl.XXXXX)

echo "Version: ${VERSION}"
echo "Sending request to ${URL}"

STATUS=$(curl \
  --silent \
  --output "${OUTPUT_FILE}" \
  --write-out "%{http_code}" \
  --request POST \
  --form "ref=master" \
  --form "variables[${VERSION_VAR_NAME}]=${VERSION}" \
  --form "token=${TOKEN}" \
  "${URL}"
)
EXIT="${?}"

if [[ ${STATUS} -ne 201 ]]; then
  EXIT="1"
  echo "Trigger failed. Request output:"
  cat "${OUTPUT_FILE}"
  echo
fi

rm -r "${OUTPUT_FILE}"

exit ${EXIT}

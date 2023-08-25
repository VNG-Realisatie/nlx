#!/usr/bin/env bash
# Copyright © VNG Realisatie 2022
# Licensed under the EUPL


PROJECT_ID="${1}"
VERSION="${2}"
VERSION_VAR_NAME="${3}"
TOKEN="${4}"

URL="https://gitlab.com/api/v4/projects/${PROJECT_ID}/trigger/pipeline"

echo "Version: ${VERSION}"
echo "Sending request to ${URL}"

STATUS=$(curl \
  --write-out "%{http_code}" \
  --request POST \
  --form "ref=main" \
  --form "variables[${VERSION_VAR_NAME}]=${VERSION}" \
  --form "token=${TOKEN}" \
  "${URL}"
)
EXIT="${?}"

if [[ ${STATUS} -ne 201 ]]; then
  EXIT="1"
  echo "Trigger failed. See request output."
  echo
fi

exit ${EXIT}

#!/usr/bin/env bash
# Copyright © VNG Realisatie 2022
# Licensed under the EUPL


BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

FORCE=0
if [ "${1}" = "-f" ]; then
  FORCE=1
fi

PKIS="external internal/organization-a internal/organization-b internal/organization-c"

for PKI in ${PKIS}; do
  PKI_DIR="${BASE_DIR}/${PKI}"

  for CERT_DIR in $(find "${PKI_DIR}/certs"/* -maxdepth 1 -type d -print); do
    CERT="$(basename "${CERT_DIR}")"

    # Skip when no certs found in directory
    if [ ! -f "${CERT_DIR}/csr.json" ]; then
      continue
    fi

    if [ -f "${CERT_DIR}/cert.pem" ] && [ ${FORCE} == 0 ]; then
      echo "Skipping ${CERT_DIR} because -f (force) parameter is not specified"
      continue
    fi

    echo "Generating certificate for ${CERT}..."

    cfssl gencert \
      -config "${PKI_DIR}/config.json" \
      -ca "${PKI_DIR}/ca/intermediate.pem" \
      -ca-key "${PKI_DIR}/ca/intermediate-key.pem" \
      -profile peer \
      "${CERT_DIR}/csr.json" \
    | cfssljson -bare "${CERT_DIR}/cert"

    cat "${PKI_DIR}/ca/intermediate.pem" >> "${CERT_DIR}/cert.pem"

    mv "${CERT_DIR}/cert-key.pem" "${CERT_DIR}/key.pem"
    rm "${CERT_DIR}/cert.csr"

  done
done

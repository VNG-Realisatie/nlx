#!/usr/bin/env bash
BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

PKIS="shared organization-a organization-b"

for PKI in ${PKIS}; do
  PKI_DIR="${BASE_DIR}/${PKI}"
  CONFIG="${PKI_DIR}/config.json"
  CA_DIR="${PKI_DIR}/ca"

  cfssl genkey \
    -config "${CONFIG}" \
    -initca \
    "${CA_DIR}/root.json" \
  | cfssljson -bare "${CA_DIR}/root"

  cfssl genkey \
    -config "${CONFIG}" \
    -initca \
    "${CA_DIR}/intermediate.json" \
  | cfssljson -bare "${CA_DIR}/intermediate"

  cfssl sign \
    -config "${CONFIG}" \
    -ca "${CA_DIR}/root.pem" \
    -ca-key "${CA_DIR}/root-key.pem" \
    -profile intermediate \
    "${CA_DIR}/intermediate.csr" \
  | cfssljson -bare "${CA_DIR}/intermediate"

  rm "${CA_DIR}/root.csr"
  rm "${CA_DIR}/intermediate.csr"
done

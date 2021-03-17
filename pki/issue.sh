#!/usr/bin/env bash
BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

FORCE=0
if [ "${1}" = "-f" ]; then
  FORCE=1
fi

PKIS="shared organization-a organization-b organization-c"

for PKI in ${PKIS}; do
  PKI_DIR="${BASE_DIR}/${PKI}"

  for CERT_DIR in $(find ${PKI_DIR}/certs/* -type d -maxdepth 0 -print); do
    CERT="$(basename ${CERT_DIR})"
    if [ -f "${CERT_DIR}/cert.pem" ] && [ ${FORCE} == 0 ]; then
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

    cat ${PKI_DIR}/ca/intermediate.pem >> "${CERT_DIR}/cert.pem"

    mv "${CERT_DIR}/cert-key.pem" "${CERT_DIR}/key.pem"
    rm "${CERT_DIR}/cert.csr"

  done
done

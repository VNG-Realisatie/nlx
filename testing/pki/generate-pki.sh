#!/bin/sh
set -e

# Create root CA certificates
cfssl gencert -initca ca-root-csr.json | cfssljson -bare ca-root
cfssl gencert -initca ca-root-second-csr.json | cfssljson -bare ca-root-second

# Create intermediate CA certificate
cfssl gencert -initca ca-intermediate-csr.json | cfssljson -bare ca-intermediate
cfssl sign -ca ca-root.pem -ca-key ca-root-key.pem -config ca-config.json -profile intermediate ca-intermediate.csr | cfssljson -bare ca-intermediate

# Create organisation certificiates
cfssl gencert -ca=ca-intermediate.pem -ca-key=ca-intermediate-key.pem -config=ca-config.json -profile=peer org-nlx-test-csr.json | cfssljson -bare org-nlx-test
cfssl gencert -ca=ca-intermediate.pem -ca-key=ca-intermediate-key.pem -config=ca-config.json -profile=peer org-without-name-csr.json | cfssljson -bare org-without-name

# Combine certificates
cat org-nlx-test.pem ca-intermediate.pem > org-nlx-test-chain.pem
cat org-without-name.pem ca-intermediate.pem > org-without-name-chain.pem

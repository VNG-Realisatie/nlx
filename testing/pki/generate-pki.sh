#!/bin/sh
set -e

cfssl gencert -initca ca-csr.json | cfssljson -bare ca
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=peer org-nlx-test-csr.json | cfssljson -bare org-nlx-test
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=peer org-without-name-csr.json | cfssljson -bare org-without-name

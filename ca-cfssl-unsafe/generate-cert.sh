#!/bin/ash

set -e # exit on error

## This script generates a cert signed with a remote cfssl server.
## It takes two arguments:
## - The first argument must be the domain name of the cert to generate.
## - The second argument must be the address on which the cfssl server can be reached.

certDomain=$1
certOrganization=$2
remoteCA=$3

csrFilename="${certDomain}-csr.json"
echo '{"hosts": ["'${certDomain}'"], "key": {"algo": "rsa", "size": 3072}, "CN": "'${certDomain}'", "names": [{"O": "'${certOrganization}'", "OU": "NLX"}]}' > "${csrFilename}"

## Wait for remote CA (cfssl server) to be online
while ! nc -z "${remoteCA}" 8888 </dev/null; do echo "waiting for ca" && sleep 1; done
## Fetch root cert from remote CA (cfssl server)
cfssl info -remote "${remoteCA}" | cfssljson -bare nlx_root
## Generate and sign cert using remote CA (cfssl server)
cfssl gencert -remote="${remoteCA}" "${csrFilename}" | cfssljson -bare "${certDomain}"

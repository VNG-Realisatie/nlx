#!/bin/ash

set -e # exit on error

## This scripts generates a new self-signed CA cert and starts the cfssl server.
## It takes one argument: the domain name for the CA.

cadomain=$1
if [ -z "${cadomain}" ]; then
	echo "Missing first argument: CA domain name";
	exit 1;
fi;

echo '{"hosts": ["'${cadomain}'"], "key": {"algo": "rsa", "size": 4096}, "names": [{"C": "NL", "ST": "Noord-Holland", "L": "Amsterdam", "O": "Common Ground", "OU": "NLX"}]}' | 
	cfssl genkey -initca /dev/stdin | 
	cfssljson -bare ca

cfssl serve --address 0.0.0.0 \
	--ca ca.pem --ca-key ca-key.pem

# pki

This document describes how the initial NLX pki is configured. This component does not contain any secrets.

## Introduction

pki-nlx describes the PKI that is used for NLX components. The PKI is setup using [cfssl](https://github.com/cloudflare/cfssl).
This document currently describes the initial version of the pki-nlx, which is not at all intended to be final and supersecure. The main goal is to have a prod and preprod environment that developers can use, and to discover what features pki-nlx should provide. This means that **future pki-nlx iterations will break the current iteration.**

### Create a veracrypt container

We store CA credentials in a veracrypt container. [Download and install](https://www.veracrypt.fr/en/Downloads.html) veracrpyt for console.

Create the container by executing the following commands. In this example we create a container for preprod. Substitute preprod with any environment you like.

```bash
cd ~/some-folder/pki-nlx/preprod

veracrypt --create

Volume type:
 1) Normal
 2) Hidden
Select [1]: < leave empty >

Enter volume path: ./preprod-ca.hc

Enter volume size (sizeK/size[M]/sizeG): 128M

Encryption Algorithm:
 1) AES
 2) Serpent
 3) Twofish
 4) Camellia
 5) Kuznyechik
 6) AES(Twofish)
 7) AES(Twofish(Serpent))
 8) Camellia(Kuznyechik)
 9) Camellia(Serpent)
 10) Kuznyechik(AES)
 11) Kuznyechik(Serpent(Camellia))
 12) Kuznyechik(Twofish)
 13) Serpent(AES)
 14) Serpent(Twofish(AES))
 15) Twofish(Serpent)
Select [1]: 14

Hash algorithm:
 1) SHA-512
 2) Whirlpool
 3) SHA-256
 4) Streebog
Select [1]: < leave empty >

Filesystem:
 1) None
 2) FAT
 3) Linux Ext2
 4) Linux Ext3
 5) Linux Ext4
 6) NTFS
 7) exFAT
Select [2]: < leave empty >

Enter password: < enter an actual password >
Re-enter password: < enter an actual password >

Enter PIM: < leave empty >

Enter keyfile path [none]: < leave empty >

Please type at least 320 randomly chosen characters and then press Enter:
< go crazy on the keyboard, but make sure you\'re not predictable ;) >

Done: 100.000%  Speed:   68 MB/s  Left: 0 s

The VeraCrypt volume has been successfully created.
```

Then mount the container:

```bash
cd ~/some-folder/pki-nlx/preprod

veracrypt --mount preprod-ca.hc /mnt/nlx-pki/preprod
Enter password for ~/some-folder/pki-nlx/preprod/preprod-ca.hc: < enter the password >
Enter PIM for ~/some-folder/pki-nlx/preprod/preprod-ca.hc: < leave empty >
Enter keyfile [none]: < leave empty >
Protect hidden volume (if any)? (y=Yes/n=No) [No]: < leave empty >
```

Create a ca

```bash
env=preprod echo '{"hosts": ["'${env}'.nlx.io"], "key": {"algo": "rsa", "size": 4096}, "names": [{"O": "Common Ground NLX CA", "OU": "NLX"}]}' | 
	cfssl genkey -initca /dev/stdin | 
	cfssljson -bare ca
```

Sign a certificate for the directory components, run for all components that need a cert.

```bash
env=preprod
component=directory-monitor
certDomain=${component}.${env}.nlx.io
certOrganization=NLX

csrFilename="${certDomain}-csr.json"
echo '{"hosts": ["'${certDomain}'"], "key": {"algo": "rsa", "size": 4096}, "CN": "'${certDomain}'", "names": [{"O": "'${certOrganization}'", "OU": "NLX"}]}' > "${csrFilename}"

## Generate and sign cert using remote CA (cfssl server)
cfssl gencert -ca ca.pem -ca-key ca-key.pem "${csrFilename}" | cfssljson -bare "${certDomain}"
```

The root cert of the ca is `~/some-folder/pki-nlx/preprod/ca.pem`.

To add a keypair as secret in kubernetes, run for alle components that need a cert:

```bash
env=preprod
component=directory-monitor
namespace=nlx-${env}-directory
certDomain=${component}.${env}.nlx.io

kubectl -n ${namespace} create secret generic certs-${component} \
	--from-file=./${certDomain}.pem \
	--from-file=./${certDomain}-key.pem \
	--from-file=./ca.pem
```

To create a key and certificate for an external party, run:

```bash
certDomain=< the domain of the inway/outway that needs a cert >
certOrganization=< name of the organization that needs the cert >

csrFilename="${certDomain}-csr.json"
echo '{"hosts": ["'${certDomain}'"], "key": {"algo": "rsa", "size": 4096}, "CN": "'${certDomain}'", "names": [{"O": "'${certOrganization}'", "OU": "NLX"}]}' > "${csrFilename}"

## Generate and sign cert using remote CA (cfssl server)
cfssl gencert -ca ca.pem -ca-key ca-key.pem "${csrFilename}" | cfssljson -bare "${certDomain}"
```

Preferably, we don't generate a key for the external party, and just verify/sign the CSR:

```bash
cfssl gencert -ca ca.pem -ca-key ca-key.pem the-csr-from-external-party.csr | cfssljson -bare output
```

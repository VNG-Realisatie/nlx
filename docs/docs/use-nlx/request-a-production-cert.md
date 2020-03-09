---
id: request-a-production-cert
title: Request a production certificate
---

In production environment we trust PKIoverheid certificates issued under the [Staat der Nederlanden Private Root CA - G1](https://www.pkioverheid.nl/). A certificate can be requested at one of the Trusted Service Providers (TSP's).

> Double check the type of certificate you are requesting. Only certificates issued under the **G1 private root** will work on NLX. Certificates issued under the EV, G2 or G3 root will not work.

## Connect to production

The process of connecting to the NLX production environment consists of the following phases:

- An organization generates a private key and certificate signing request (CSR).
- A certificate is requested at one of the TSPs using the CSR.
- The organization installs the certificate on their NLX inways / outways and connects to the NLX pre-production environment.
- *(optional)* The organization can request a voluntary audit of their infrastructure from the NLX support team by sending an email to [support@nlx.io](mailto:support@nlx.io).
- After verifying the setup, the organization connects to the production environment using the same certificate.

## Generate a private key and CSR

A private key and certificate signing request for the (pre-)production environment can be generated with the following command:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

Answer the questions accordingly:

- **Country Name**, enter `Netherlands`
- **State**, enter the name of the province, e.g. `Utrecht`
- **Locality Name**, enter the name of the city, e.g. `Utrecht`
- **Organization Name**, enter your organization name
- **Organization Unit Name**, enter the name of your organization unit name (optional)
- **Common name**, for an inway this should correspond to the FQDN of your inway, e.g. `inway.my-organization.nl` or `nlx.my-organization.nl`. For an outway this FQDN does not have to be resolvable. It is possible to use the same certificate for an outway and inway.
- **Email Address**, enter your e-mail address (optional)
- **A challenge password**, omit this value

The command outputs two files: *org.key*, the private key and *org.csr*, the certificate signing request.

> **Warning**
>
> Keep your private key securely stored and do not transfer the private key `org.key` unprotectedly.

## Request a certificate

To request a *Staat der Nederlanden Private Root CA - G1* certificate, contact one of the Trusted Service Providers (TSPs):

- [QuoVadis](https://www.quovadisglobal.nl/DigitaleCertificaten/PKIOverheidCertificaten.aspx)
- [Digidentity](https://sslstore.digidentity.eu/)
- [KPN](https://certificaat.kpn.com/pkioverheidcertificaten/servercertificaten/private-servercertificaat/)

Follow the approval process and submit the certificate signing request *org.csr*. When you retrieved the certificate, name the files accordingly:

- *root.crt*, the certificate of the [Staat der Nederlanden Private Root CA - G1](/certs/PKIoverheid-PrivateRootCA-G1.crt)
- *org.crt*, the certificate you retrieved from the TSP, including the intermediate CA certificates
- *org.key*, the private key you generated with the OpenSSL command

## Environments

Inways and outways use the `DIRECTORY_REGISTRATION_ADDRESS` and `DIRECTORY_INSPECTION_ADDRESS` settings to connect to the directory. Use the following values to connect to the (pre-)production environments:

**Pre-production**
- directory-inspection-api.preprod.nlx.io
- directory-registration-api.preprod.nlx.io

**Production**
- directory-inspection-api.nlx.io
- directory-registration-api.nlx.io

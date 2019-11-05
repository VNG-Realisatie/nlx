---
id: request-a-production-cert
title: Request a production certificate
---

Certificates for the NLX production environment are handed out by the NLX Certificate Authority. Before organizations are able to connect to the production environment, we require you to first connect to the pre-production environment. After an audit of the NLX support team, a production certificate is handed out. This certificate can then be used to setup connections through the NLX production environment.

## Admission process

The admission process of the NLX production environment consists of the following phases:

- An organization generates a private key and certificate signing request (CSR). The CSR is sent and via email to [support@nlx.io](mailto:support@nlx.io).
- The support team will verify the identity of the requester and  the fingerprint of the CSR by video call. If everything is validated we will hand out a pre-production certificate.
- The organization installs the pre-production certificate on their NLX inways / outways.
- The organization requests an audit from the support team.
- The support team verifies the NLX setup at the organization and approves the setup.
- The organization generates a new private key and CSR for the production certificate and sends this to [support@nlx.io](mailto:support@nlx.io).
- The support team verifies the fingerprint of the CSR by video call and hands out a production certificate.
- The organization installs the production certificate on their NLX inways / outways.

## Generate a private key and CSR

A private key and certificate signing request for the (pre)-production environment can be generated with the following command:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

> **Warning**
>
> Keep your private key securely stored and do not transfer the private key `org.key` unprotectedly.

Answer the questions accordingly:

- **Country Name**, enter `Netherlands`
- **State**, enter the name of the province, e.g. `Utrecht`
- **Locality Name**, enter the name of the city, e.g. `Utrecht`
- **Organization Name**, enter your complete organization name lowercase and URL-friendly, e.g. `my-organization`
- **Organization Unit Name**, omit this value
- **Common name**, for an inway this should correspond to the FQDN of your inway, e.g. `inway.my-organization.nl`. For an outway you can use e.g. `outway.my-organization.nl`. For an outway this FQDN does not have to be resolvable.
- **Email Address**, omit this value
- **A challenge password**, omit this value

## Environments
Inways and outways use the `DIRECTORY_REGISTRATION_ADDRESS` and `DIRECTORY_INSPECTION_ADDRESS` settings to connect to the directory. Use the following values to connect to the (pre-)production environments:

**Pre-production**
- directory-inspection-api.preprod.nlx.io
- directory-registration-api.preprod.nlx.io

**Production**
- directory-inspection-api.prod.nlx.io
- directory-registration-api.prod.nlx.io

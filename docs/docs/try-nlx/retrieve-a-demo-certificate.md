---
id: retrieve-a-demo-certificate
title: Retrieve a demo certificate
---

## Introduction

To be able to send traffic through the NLX network, you'll need a certificate and private key.
The certificate and key are used to encrypt traffic between you and other nodes.

In this part we will generate & install all required certificates.

## Demo CA root certificate

The very first certificate we need is the CA root certificate.
This one is used to validate certificates of other organizations.

Download [the root certificate](https://certportal.demo.nlx.io/root.crt) file and save it as `root.crt` in `~/nlx-setup/`.
That is the working directory as described in [part 1](./setup-your-environment.md).

>  If you are using Windows make sure that you select `All files` as file-type when saving the root certificate. If you do not do this Windows will add the extension `txt` to the file.

### Verification

Let's check if our certificate is alright.

```bash
openssl x509 -in root.crt -text | grep Subject:
```

Example of the output: `Subject: C=NL, ST=Noord-Holland, L=Amsterdam, O=Common Ground, OU=NLX`

## The certificate for our service

Now we have the root certificate installed. The next certificate we need is our own generated certificate.
It should include information about the API we will provide or consume.

In order to request or own certificate, we need to generate a key and Certificate Signing Request (CSR).
We can create these using [openssl](https://www.openssl.org/).

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

Answer the questions accordingly:

- **Country Name**, enter any value
- **State**, enter any value
- **Locality Name**, enter any value
- **Organization Name**, please enter a URL-friendly value. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness yet.<br>A good value could be: `my-organization`.
- **Organization Unit Name**, enter any value
- **Common name**, when you would like to offer your API's to the NLX network make sure this corresponds to your external hostname. For this guide we will use `my-organization.nl`.
- **Email Address**, enter any value
- **A challenge password**, leave empty

Now openssl wil generate two files:

- A private key `org.key`.  **Keep this file safe**, limit access to it and do not transfer it in an insecure way.
- A certificate request `org.csr`. Use this file to request a certificate.

### Verification

List the files in your working directory.

```bash
ls -la
```

If all of the above went well, you should see at least two files listed:

* org.csr
* org.key

## Retrieve the certificate

We will use the NLX certportal to retrieve an NLX developer certificate.

First let's copy the contents of `org.csr`. We will use this to generate the demo certificate.
Make sure to copy the complete output, including *-----BEGIN CERTIFICATE REQUEST-----* and *-----END CERTIFICATE REQUEST-----*.

```bash
cat org.csr
```

Open [certportal](https://certportal.demo.nlx.io) and paste the content in the `CSR` field.

Scroll to the bottom of the page and click on **Request certificate**.

The system will instantly sign your csr and return your certificate.
You can either copy paste your certificate and store it in a file or you can click **Download certificate** to download the certificate.

Rename the file from `certificate.crt` to `org.crt` and store the file next to your private key.

>  If you are using Windows make sure that you select `All files` as file type when saving your certificate. If you do not do this Windows will add the extenstion `txt` to the file.

### Verification

Let's check if our certificate is alright.

```bash
openssl x509 -in org.crt -text | grep Subject:
```

The output should contain the answers you've provided when you created the certificate.

Example of the output: `Subject: C=nl, ST=noord-holland, L=haarlem, O=my-organization, OU=my-organization-unit, CN=an-awesome-organization.nl`


## In sum

All required certificates are available now. So far, we have:

- downloaded the CA root certificate.
- generated our own certificate, so we are allowed to communicate with the API's on the NLX network.

Now let's see if we can consume an API from the NLX network.

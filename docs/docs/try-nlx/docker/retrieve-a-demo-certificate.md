---
id: retrieve-a-demo-certificate
title: Retrieve a demo certificate
---

To be able to send traffic through the NLX network, you'll need a certificate and private key.
The certificate and key are used to encrypt traffic between you and other nodes.

In this part we will generate & install all required certificates.
Before you continue, make sure you have [set up your environment](./setup-your-environment).

## Demo CA root certificate

The very first certificate we need is the CA root certificate.
This one is used to validate certificates of other organizations.

Download [the root certificate](https://certportal.demo.nlx.io/root.crt) file and save it as `root.crt` in `~/nlx-setup/`.
That is the working directory as described in [Setup your environment](./setup-your-environment).

>  If you are using Windows make sure that you select `All files` as file-type when saving the root certificate. If you do not do this Windows will add the extension `txt` to the file.

### Verification

Let's check if our certificate is alright.

```bash
openssl x509 -in root.crt -text | grep Subject:
```

Example of the output: `Subject: C=NL, ST=Noord-Holland, L=Amsterdam, O=Common Ground, OU=NLX`

## The certificate for our service

We now have the root certificate installed. The next certificate we need is our own generated certificate.
It should include information about the API we will provide or consume.

In order to request or own certificate, we need to generate a key and Certificate Signing Request (CSR).
We can create these using [OpenSSL](https://www.openssl.org/).

**Note**

> NLX uses the Subject Serial Number field of a certificate to identify an organization.
> If you want to use the Subject Serial Number that you provice when generating the CSR, you need to add the following lines to your OpenSSL config:
>
> ```toml
> [ req_distinguished_name ]
> serialNumber = <your-subject-serial-number>
> ```
>
> For more information about this, see the [organization identification](/reference-information/organization-identification) section.

Now let's generate the CSR:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

Answer the questions accordingly:

- **Country Name**, enter any value
- **State**, enter any value
- **Locality Name**, enter any value
- **Organization Name**, please enter a URL-friendly value with a maximum length of 100 characters.
  A good value could be: `my-organization`.
- **Organization Unit Name**, enter any value
- **Common name**, this should correspond to the Fully Qualified Domain Name (FQDN) of your Inway,
  we will use `my-organization.nl` for this guide. For an Outway this FQDN does not have to be resolvable. It is possible to use the Inway certificate for the Outway and NLX Management.
- **Email Address**, enter any value
- **Serial Number** (optional), enter a serial number with a maximum length of 20 characters. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness.
- **A challenge password**, leave empty

Now OpenSSL will generate two files:

- A private key `org.key`.  **Keep this file safe**, limit access to it and do not transfer it in an insecure way.
- A certificate request `org.csr`. Use this file to request a certificate.

### Verification

List the files in your working directory.

```bash
ls -la
```

If all the steps above went well, you should see at least two files listed:

* org.csr
* org.key

## Retrieve the certificate

We will use the NLX certportal to retrieve an NLX developer certificate.

First let's copy the contents of `org.csr`. We will use this to generate the demo certificate.
Make sure to copy the complete output, including `-----BEGIN CERTIFICATE REQUEST-----` and `-----END CERTIFICATE REQUEST-----`.

```bash
cat org.csr
```

Open the [NLX Certportal](https://certportal.demo.nlx.io) and paste the content in the `CSR` field.

Scroll to the bottom of the page and click on **Request certificate**.

The system will instantly sign your CSR and return your certificate.
You can either copy and paste your certificate and store it in a file or you can click **Download certificate** to download the certificate.

Rename the file from `certificate.crt` to `org.crt` and store the file next to your private key.

>  If you are using Windows, make sure that you select `All files` as file type when saving your certificate. If you do not do this, Windows will add the extension `txt` to the file.

### Verification

Let's check if our certificate is alright.

```bash
openssl x509 -in org.crt -text | grep Subject:
```

The output should contain the answers you've provided when you created the certificate.

Example of the output: `Subject: C=nl, ST=zuid-holland, L=gemeente-stijns, O=my-organization, OU=my-organization-unit, CN=an-awesome-organization.nl/serialNumber=01234567890123456789`.

The value after `serialNumber=` in the Subject's CN field is the primary way to identify your organization on NLX.
For details about this, see the [organization identification](/reference-information/organization-identification) page.

## In sum

All required certificates are available now. So far, we have:

- Downloaded the CA root certificate and stored it in the `~/nlx-setup` folder as `root.crt`.
- Generated our own certificate and private key and stored them in the `~/nlx-setup` folder as `org.crt`(certificate) and `org.key`(private key), so we are allowed to communicate with the API's on the NLX network.

Now let's [get up and running](./getting-up-and-running) to make sure you have all software installed to get started.

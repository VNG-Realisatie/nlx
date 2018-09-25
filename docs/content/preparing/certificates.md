---
title: "Certificates"
description: ""
weight: 100
menu:
  docs:
    parent: "preparing"
---

To connect with other nodes on the NLX network, you'll need a certificate and private key. The certificate and key are used to encrypt traffic between you and other nodes.

## Generate a demo key and csr

First generate a key and csr using [openssl](https://www.openssl.org/). This can be done using the following command:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

Answer the questions accordingly:

- **Country Name**, enter any value
- **State**, enter any value
- **Locality Name**, enter any value
- **Organization Name**, please enter a URL-friendly value. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness yet.<br>A good value could be: `an-awesome-organization`.
- **Organization Unit Name**, enter any value
- **Common name**, when you would like to offer services to the NLX network make sure this corresponds to your external hostname.
- **Email Address**, enter any value
- **A challenge password**, leave empty

Now openssl wil generate two files:

- A private key `org.key`.  **Keep this file safe**, limit access to it and do not transfer it unprotectedly.
- A certificate request `org.csr`. Use this file to request a certificate.

We will use the NLX certportal to retrieve an NLX developer certificate.

## Request a demo certificate

Now point your browser to [certportal.demo.nlx.io](https://certportal.demo.nlx.io) to request a certificate. Enter the content of `org.csr` and click on **Request certificate**. The system will instantly sign your csr and returns the certificate. Download the certificate and store it next to your private key as `org.crt`.

## Downloading the demo CA root certificate

To validate certificates of other organizations, you will need our demo CA's root certificate. It's available for download at https://certportal.demo.nlx.io/root.crt

Now you are ready to develop on NLX.

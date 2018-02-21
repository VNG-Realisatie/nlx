---
title: "Getting started"
description: ""
weight: 100
menu:
  docs:
    parent: "developing-on-nlx"
---

Before you can start developing on NLX you need to request a development organisation certificate. With this certificate you can build applications on top of NLX and also provide services to the NLX network.

## Generate a certificate
First generate a certificate request using [openssl](https://www.openssl.org/). This can be done using the following command:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:2048 -keyout {yourhostname}.key -out {yourhostname}.csr
```

Answer the questions. Now openssl wil generate two files:

- A private key ```{yourhostname}.key```.  **Keep this file safe**, limit access to it and do not transfer it unprotectedly.
- A certificate request ```{yourhostname}.csr```. Use this file to request a certificate.

We will use the NLX certportal to retieve an NLX developer certificate.

## Request a certificate
Now point your browser to [certportal.nlx.io](https://certportal.nlx.io) to request a certificate. Enter the content of ```{yourhostname}.csr``` and click on **Request certificate**. The system will instantly generate a valid developer certificate. Download the certificate and store it next to your private key.

Now you are ready to develop on NLX.

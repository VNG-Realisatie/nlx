---
title: "Getting started"
description: ""
weight: 100
menu:
  docs:
    parent: "developing-on-nlx"
---

Before you can start developing on NLX you need to request a development organisation certificate. This certificate allows you to build applications on top of NLX services and also offer services to the network yourselves.

## Generate a certificate
First generate a certificate request with [openssl](https://www.openssl.org/). This can be done using the following command:

```bash
    openssl req -utf8 -nodes -sha256 -newkey rsa:2048 -keyout {yourhostname}.key -out {yourhostname}.csr
```

Follow the steps and answer the questions.

## Request a certificate
Now point your browser to [certportal.nlx.io](https://certportal.nlx.io) to request a certificate. Enter the content of ```{yourhostname}.csr``` and click on **Request certificate**. The system will instantly generate a valid developer certificate. Download the certificate. With your certificate you are ready to develop on NLX.

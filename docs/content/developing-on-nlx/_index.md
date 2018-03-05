---
title: "Developing on NLX"
description: ""
weight: 10
---

## Getting started
Before you can start developing on NLX you need to request a development organisation certificate. With this certificate you can build applications on top of NLX and also provide services to the NLX network.

## Generate a certificate
First generate a certificate request using [openssl](https://www.openssl.org/). This can be done using the following command:

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:2048 -keyout {yourhostname}.key -out {yourhostname}.csr
```

Answer the questions accordingly:

- **Country Name**, enter any value
- **State**, enter any value
- **Locality Name**, enter any value
- **Organization Name**, please enter a URL-friendly value. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness yet.<br>A good value could be: `an-awesome-organization`.
- **Organization Unit Name**, enter any value
- **Common name**, when you would like to offer services to the NLX network make sure this corresponds to your external hostname or IP address.
- **Email Address**, enter any value
- **A challenge password**, leave empty

Now openssl wil generate two files:

- A private key ```{yourhostname}.key```.  **Keep this file safe**, limit access to it and do not transfer it unprotectedly.
- A certificate request ```{yourhostname}.csr```. Use this file to request a certificate.

We will use the NLX certportal to retrieve an NLX developer certificate.

## Request a certificate
Now point your browser to [certportal.nlx.io](https://certportal.nlx.io) to request a certificate. Enter the content of ```{yourhostname}.csr``` and click on **Request certificate**. The system will instantly generate a valid developer certificate. Download the certificate and store it next to your private key.

## Downloading the CA root certificate
To validate certificates of other organizations, you will need our demo CA's root certificate. It's available for download at https://certportal.nlx.io/root.crt

Now you are ready to develop on NLX.

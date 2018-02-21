---
title: "Creating a service"
description: ""
weight: 120
menu:
  docs:
    parent: "developing-on-nlx"
---

## Deploying an inway
To provide a service to the NLX network you first need to deploy an NLX inway. First make sure you installed the latest version of [Docker](https://www.docker.com) on your machine. Also make sure you generated a private key and aquired a certificate. The steps to aquire a certificate are described in [getting started](../getting-started/).

Now download the certificate of the NLX development CA:

```bash
wget https://dev.nlx.io/root.crt
```

And store it next to your private key and certificate. Now start a new inway:


```bash
docker run -d \
-v ./root.crt:/certs/root.crt \
-v ./{yourhostname}.crt:/certs/inway.crt \
-v ./{yourhostname}.key:/certs/inway.key \
-e TLS_NLX_ROOT_CERT=/certs/root.crt \
-e TLS_ORG_CERT=/certs/inway.crt \
-e TLS_ORG_KEY=/certs/inway.key \
-p 2018:2018 \
nlxio/inway:stable
```

The inway now connects itself to the NLX network. Please **make sure** connections on port ```2018``` at your public IP address are routed to the machine running the NLX inway otherwise connections to your inway and services will fail.

## Configuring a service
This is still work in progress. When the work is finished you are able to configure services on the inway. Then all organisations on NLX can query your service on their outway as follows:

```bash
curl http://{OutwayIP}:12018/{OrganisationName}/{ServiceName}/{endpoint}
```

## Authorisation
A service is able to implement its own autorisation logic by using a NLX-specific header. The inway will automatically append the following header with every request:

    X-NLX-Organisation-CN: {OrganisationName}

where ```{OrganisationName}``` is set to the organisation performing the request.

---
title: "Building an application"
description: ""
weight: 110
menu:
  docs:
    parent: "developing-on-nlx"
---

## Deploying an outway
To build an application that uses services of the NLX network you first need to deploy an NLX outway. First make sure you installed the latest version of [Docker](https://docker.com) on your machine. Also make sure you generated a private key and aquired a certificate. The steps to aquire a certificate are described in [getting started](../getting-started/).

Now download the certificate of the NLX development CA:

```bash
wget https://certportal.nlx.io/root.crt
```

And store it next to your private key and certificate. Now start a new outway:

```bash
docker run -d \
-v ./root.crt:/certs/root.crt \
-v ./{yourhostname}.crt:/certs/outway.crt \
-v ./{yourhostname}.key:/certs/outway.key \
-e TLS_NLX_ROOT_CERT=/certs/root.crt \
-e TLS_ORG_CERT=/certs/outway.crt \
-e TLS_ORG_KEY=/certs/outway.key \
-p 12018:12018 \
nlxio/outway:latest
```

The outway is now running on http://localhost:12018.

## Querying services
To query services on the NLX network, use the following structure:

```bash
curl http://localhost:12018/{OrganisationName}/{ServerName}/{endpoint}
```

For example to query the central NLX echo service, use:

```bash
curl http://localhost:12018/DemoProviderOrganization/PostmanEcho/get?foo=1
```

Congratulations, you now made your first query on the NLX network!
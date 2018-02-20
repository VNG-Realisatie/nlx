---
title: "Building an application"
description: ""
weight: 110
menu:
  docs:
    parent: "developing-on-nlx"
---

## Deploying an outway
To build an application on top of services that are offered on the NLX network you first need to deploy an NLX outway. First make sure you installed the latest version of [Docker](https://docker.com) on your machine. 

First make sure to have a copy of the NLX development network root certificate:

```bash
wget https://dev.nlx.io/root.crt
```

Now start a new outway:

```bash
  docker run -d \
  -v ./root.crt:/certs/root.crt \
  -v ./{yourhostname}.crt:/certs/outway.crt \
  -v ./{yourhostname}.key:/certs/outway.key \
  -e TLS_NLX_ROOT_CERT=/certs/root.crt \
  -e TLS_ORG_CERT=/certs/outway.crt \
  -e TLS_ORG_KEY=/certs/outway.key \
  -p 12018:12018 \
  nlxio/outway
```

The outway is now running on http://127.0.0.1:12018

## Querying services
To query services on the NLX network, use the following structure:

    curl http://127.0.0.1:12018/{OrganisationName}/{ServerName}/{endpoint}


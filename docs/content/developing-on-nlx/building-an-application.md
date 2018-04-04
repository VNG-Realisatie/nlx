---
title: "Building an application"
description: ""
weight: 110
menu:
  docs:
    parent: "developing-on-nlx"
---

## Deploying an outway
To build an application that uses services of the NLX network you first need to deploy an NLX outway. First make sure you installed the latest version of [Docker](https://docker.com) on your machine. Also make sure you generated a private key and aquired a certificate. The steps to aquire a certificate are described in [getting started](../).

Now download the certificate of the NLX development CA:

```bash
wget https://certportal.demo.nlx.io/root.crt
```

And store it next to your private key and certificate. Now start a new outway:

```bash
docker pull nlxio/outway:latest
docker run -d \
-v {/absolute/path/to/root.crt}:/certs/root.crt \
-v {/absolute/path/to/yourhostname.crt}:/certs/outway.crt \
-v {/absolute/path/to/yourhostname.key}:/certs/outway.key \
-e DIRECTORY_ADDRESS=directory.demo.nlx.io:1984 \
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
curl http://localhost:12018/{organization-name}/{service-name}/{endpoint}
```

For example to query the NLX demo application, use:

```bash
curl http://localhost:12018/demo-organization/demo-api/
```

Congratulations, you now made your first query on the NLX network!

## Overview of services
The directory will provide an overview of all services that are currently registered on NLX. Browse to [directory.demo.nlx.io](https://directory.demo.nlx.io/) to see an actual overview of services.
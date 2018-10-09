---
title: "Start an outway"
description: ""
weight: 110
menu:
  docs:
    parent: "using-services-on-nlx"
---

To use a service that is provided via NLX, you need to route traffic through an outway onto the network.

## Start an outway using Docker

First make sure you have installed a recent version of [Docker](https://www.docker.com) on your machine. Also make sure you generated a private key and aquired a signed certificate by following [these steps](../../preparing/certificates).

To start the outway, run the following docker commands:

```bash
docker pull nlxio/outway:latest

docker run --detach \
  --name my-nlx-outway \
  --volume {/absolute/path/to/root.crt}:/certs/root.crt \
  --volume {/absolute/path/to/yourhostname.crt}:/certs/org.crt \
  --volume {/absolute/path/to/yourhostname.key}:/certs/org.key \
  --env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443 \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --publish 4080:80 \
  nlxio/outway:latest
```

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`. The outway is now running and listening on `http://localhost:4080`.

To stop the outway run `docker stop my-nlx-outway && docker rm my-nlx-outway`.

## Querying services

To query services on the NLX network, use the following structure:

```bash
curl http://localhost:4080/{organization-name}/{service-name}/{api-specific-path}
```

For example, to query the NLX demo application use:

```bash
curl http://localhost:4080/vng-realisatie/demo-api/
```

Congratulations, you now made your first query on the NLX network!

## Overview of services

The directory will provide an overview of all services that are currently registered on NLX. Browse to [directory.demo.nlx.io](https://directory.demo.nlx.io/) to see an actual overview of services.

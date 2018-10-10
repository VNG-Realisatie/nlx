---
title: "Start an outway"
description: ""
weight: 110
menu:
  docs:
    parent: "using-services-on-nlx"
---

To use a service that is provided via NLX, you need to route traffic through an outway onto the network. In this guide we'll show you how to run the outway using docker. This guide is written for Mac/Linux users and assumes you have some experience with the terminal/shell. Windows 10 users may be able to follow this tutorial using [ubuntu for windows](https://tutorials.ubuntu.com/tutorial/tutorial-ubuntu-on-windows). However, the windows platform is not officially supported, we advise to use a VM with Ubuntu installed.


## Start an outway using Docker

First make sure you have installed a recent version of [Docker](https://www.docker.com) on your machine. To run the outway, we need a certificate for encrypted/TLS communications with other NLX services on the network. So we need make sure you generated a private key and aquired a signed certificate by following [these steps](../../preparing/certificates).

To start the outway, run the following docker commands:

```bash
docker pull nlxio/outway:latest

docker run --detach \
  --name my-nlx-outway \
  --volume {/absolute/path/to/root.crt}:/certs/root.crt:ro \
  --volume {/absolute/path/to/yourhostname.crt}:/certs/org.crt:ro \
  --volume {/absolute/path/to/yourhostname.key}:/certs/org.key:ro \
  --env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443 \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --publish 4080:80 \
  nlxio/outway:latest
```

We give docker several arguments:

- `--detach` will make the container run in the background and print the container ID
- `--name my-nlx-outway` the name of your docker container 
- `--volume {/absolute/path/to/root.crt}:/certs/root.crt:ro` tells docker to make the `/certs/root.crt` file available inside the container.
- `--volume {/absolute/path/to/org.crt}:/certs/org.crt:ro` tells docker to make the `/certs/org.crt` file available inside the container.
- `--volume {/absolute/path/to/org.key}:/certs/org.key:ro` tells docker to make the `/certs/org.key` file available inside the container.
- `--env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443` sets the environment variable `DIRECTORY_ADDRESS` this address is used by the outway to anounce itself to the directory.
- `--env TLS_NLX_ROOT_CERT=/certs/root.crt`sets the environment variable `TLS_NLX_ROOT_CERT` this is the location of the root certificate.
- `--env TLS_ORG_CERT=/certs/org.crt` sets the environment variable `TLS_ORG_CERT` this is the location of the organisation certificate.
- `--env TLS_ORG_KEY=/certs/org.key` sets the environment variable `TLS_ORG_KEY` this is the location of the organisation private key.
- `--env DISABLE_LOGDB=1` sets the environment variable `DISABLE_LOGDB` the value 1 will disable the transaction logs, the value 0 will enable them.
- `--publish` connects port 4080 on the host machine to port 80 inside the container. This way, we can send requests to the inway.
- ` nlxio/outway:latest` is the name of our docker image (`nlxio/outway`) as stored in the docker registry and the version we want to use (`latest`). The `--` tells docker that all arguments after this one are meant for the outway process, not for docker itself.

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`. The outway is now running and listening on `http://localhost:4080`.

To verify your container is running execute the `docker ps` command, this command will print the active containers. If your container is not running try to start it again without the `--detach` flag, the logs of the container will now be printed in your terminal. 



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

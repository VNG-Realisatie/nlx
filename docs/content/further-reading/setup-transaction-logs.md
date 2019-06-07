---
title: "Setup transaction logs"
description: ""
weight: 10
menu:
  docs:
    parent: "further-reading"
---


## Before you start

Make sure you read our [Get started](../../get-started/setup-your-environment) guide.

Make sure you have installed a recent version of [Docker-Compose](https://docs.docker.com/compose/install/).

Make sure you have installed a recent version of [Git](https://git-scm.com/downloads).


## Introduction

The inway and the outway are able to log meta-data of the requests they process. In this guide you will learn how to setup the transaction-log service and connect it to an outway.

## Docker compose

To make the setup of the transaction-log as easy as possible you can clone our nlx-compose repository. This repository contains a docker-compose file for setting up the transaction-log. Docker-compose is a tool to setup multiple docker containers at once.

Start by cloning our nlx-compose repository using git.

execute

```bash
git clone https://gitlab.com/commonground/nlx-compose
```

This repository contains `transaction-log.yml` this script will
* Setup a [postgres](https://www.postgresql.org/) database.
* Create the transaction-log database structure.
* Setup the transaction-log-api. The transaction-log-api is an API we can use to retrieve log records from the database.

After pulling the repository lets navigate to the correct directory.

execute

```bash
cd nlx-compose
```

## Start the transaction-log

We can now use `docker-compose` to setup our transaction-log

execute

```bash
docker-compose -f transaction-log.yml up -d
```


## Verify

To verify that `docker-compose` runned successfuly we use the transaction-log-api. Log records of outways connected to you transaction-log can be retrieved by making a http call to `http://localhost:12019/out` (log records inways are available on `http://localhost:12019/in`)

Lets make a request to retrieve the log records of an outway

execute

``` bash
curl http://localhost:12019/out
```

There won't be any records in the transaction-log so the response should be

```json
   {"records":[]}
```

Congratulations, the transaction-log is up and running!


## Connect an outway to the transaction log

The next step is to connect an inway or an outway to the transaction-log database. If you've followed the [Get started](../../get-started/setup-your-environment) guide you will either have an inway or an outway running. In this example we will connect the transaction log to the outway setup in the [Get started](../../get-started/consume-an-api) guide.

Stop the outway you creating in the [Get started](../../get-started/consume-an-api) guide, if it is still running.

execute

```bash
docker rm -f my-nlx-outway
```

To connect the transaction-log to our outway we need to start the outway with a different docker command

execute

```bash
docker run --detach \
  --name my-nlx-outway \
  --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
  --volume ~/nlx-setup/yourhostname.crt:/certs/org.crt:ro \
  --volume ~/nlx-setup/yourhostname.key:/certs/org.key:ro \
  --env DIRECTORY_INSPECTION_ADDRESS=directory-api.demo.nlx.io:443 \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env POSTGRES_DSN='postgres://postgres:postgres@postgres/txlog-db?sslmode=disable&connect_timeout=10' \
  --network nlx-network \
  --publish 80:8080 \
  nlxio/outway:latest
```

There are some key differences in the docker command compared to the one used to start the outway in [Get started](../../get-started/consume-an-api).

We will go over them briefly

* We removed the flag `--env DISABLE_LOGDB=1`, this flag disables the transaction-log on the outway and defeats the purpose of this guide.
* We added the flag `--env POSTGRES_DSN='postgres://postgres:postgres@postgres/txlog-db?sslmode=disable&connect_timeout=10'`, this flags stores the postgres connection string in the environment variable `POSTGRES_DSN` on the docker container. The outway uses this connection string to setup a connection to the transaction-log database.
* We added the flag `--network nlx-network` this connects the docker container to the same docker network as the transaction-log (this network is created by the docker-compose-file we executed earlier). Without this parameter the outway wouldn't be able to connect to the transaction-log.


## Verify

Let's make a request through our outway and check if the log records are being writen.

execute

```bash
curl http://localhost/haarlem/demo-api/get?foo1=bar1
```

If successful the response should be

```json
{
	"args": {
	  "foo1": "bar1"
	},
	"headers": {
	  "x-forwarded-proto": "https",
	  "host": "postman-echo.com",
	  "accept": "*/*",
	  "accept-encoding": "gzip",
	  "user-agent": "curl/7.54.0",
	  "x-nlx-logrecord-id": "<arbitrary-logrecord-id>",
	  "x-nlx-request-organization": "my-organization",
	  "x-forwarded-port": "443"
	},
	"url": "https://postman-echo.com/get?foo1=bar1"
}
```

Now the outway should have written records in the transaction-log. Lets query the transaction-log API to verify the log records were written.

execute

``` bash
curl http://localhost:12019/out
```

The response should be

```json
{
  "records": [
    {
      "source_organization": "my-organization",
      "destination_organization": "haarlem",
      "service_name": "demo-api",
      "logrecord-id": "70824crkkpvpu",
      "data": {
        "request-path": "/get"
      },
      "DataSubjects": null,
      "created": "2019-02-06T10:56:18.090384Z"
    }
  ],
  "page": 0,
  "rowsPerPage": 0,
  "rowCount": 1
}

```

Congratulations, you've connected your outway to the transaction-log!

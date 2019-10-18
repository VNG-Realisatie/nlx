---
title: "Setup transaction logs"
description: ""
weight: 10
menu:
  docs:
    parent: "further-reading"
---


## Before you start

- Make sure you read our [Get started](../../get-started/setup-your-environment) guide.
- Make sure you have installed a recent version of [Docker-Compose](https://docs.docker.com/compose/install/).
- Make sure you have installed a recent version of [Git](https://git-scm.com/downloads).

## Introduction

The inway and the outway are able to log metadata of the requests they process. In this guide you will learn how to write transaction logs from the outway service.

## Docker compose

To make the setup of the transaction logs as easy as possible, we have provided a `docker-compose.yml` file in the [nlx-compose repository](https://gitlab.com/commonground/nlx-compose). It contains configuration for a PostgreSQL container and executes the migrations necessary for the txlog-db.

## Installation

Start by cloning the [nlx-compose repository](https://gitlab.com/commonground/nlx-compose).

```bash
git clone https://gitlab.com/commonground/nlx-compose
```

After pulling the repository lets navigate to the correct directory

```bash
cd nlx-compose
```

## Start the required containers

We can now use the Docker compose configuration to start the services we need for the transaction logs.

```bash
docker-compose up -d
```

## Viewing the logs

With a SQL command we are able to view the logs in the database:

```bash
docker-compose exec postgres psql -d txlog-db -c "SELECT * FROM transactionlog.records ORDER BY id DESC;"
 id | direction | created | src_organization | dest_organization | service_name | logrecord_id | data
----+-----------+---------+------------------+-------------------+--------------+--------------+------
(0 rows)
```

As you can see there are nog logs yet.

## Connect an outway to the transaction log

The next step is to connect an inway or an outway to the database. If you've followed the [Get started](../../get-started/setup-your-environment) guide you will either have an inway or an outway running. In this example we will connect the transaction log to the outway setup in the [Get started](../../get-started/consume-an-api) guide.

If your outway from the [Get started](../../get-started/consume-an-api) guide is still running, you will have to stop it first.

execute

```bash
docker rm -f my-nlx-outway
```

To connect the transaction log database to our outway we need to start the outway with a different docker command:

```bash
docker run --rm \
  --name my-nlx-outway \
  --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
  --volume ~/nlx-setup/yourhostname.crt:/certs/org.crt:ro \
  --volume ~/nlx-setup/yourhostname.key:/certs/org.key:ro \
  --env DIRECTORY_INSPECTION_ADDRESS=directory-inspection-api.demo.nlx.io:443 \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env POSTGRES_DSN='postgres://postgres:postgres@postgres/txlog-db?sslmode=disable&connect_timeout=10' \
  --network nlx \
  --publish 80:8080 \
  nlxio/outway:latest
```

There are some key differences in the Docker command compared to the one used to start the outway in [Get started](../../get-started/consume-an-api).

We will go over them briefly

* We removed the flag `--env DISABLE_LOGDB=1`, this flag disables the transaction-log on the outway and defeats the purpose of this guide.
* We added the flag `--env POSTGRES_DSN='postgres://postgres:postgres@postgres/txlog-db?sslmode=disable&connect_timeout=10'`, this flags stores the postgres connection string in the environment variable `POSTGRES_DSN` on the docker container. The outway uses this connection string to setup a connection to the transaction-log database.
* We added the flag `--network nlx`, this connects the docker container to the same docker network as the transaction log database (this network is created by the docker-compose-file we executed earlier). Without this parameter the outway wouldn't be able to connect to the transaction log.


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

Now the outway should have written records to the transaction log database. Let's query the transaction log API, to verify the log records were written.

```bash
docker-compose exec postgres psql -d txlog-db -c "SELECT * FROM transactionlog.records ORDER BY id DESC;"

id | direction |            created            | src_organization | dest_organization | service_name | logrecord_id  |          data
---+-----------+-------------------------------+------------------+-------------------+--------------+---------------+-------------------------
 3 | out       | 2019-06-28 10:58:50.63158+00  | barttest         | haarlem           | demo-api     | dmv593btpr9rh | {"request-path": "get"}
(1 row)
```

Congratulations, you've connected your outway to the transaction log!

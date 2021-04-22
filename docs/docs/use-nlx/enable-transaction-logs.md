---
id: enable-transaction-logs
title: Enable transaction logs
---

## Before you start

- Make sure you read our [try NLX](../try-nlx/setup-your-environment.md) guide.
- Make sure you have installed a recent version of [Docker Compose](https://docs.docker.com/compose/install/).
- Make sure you have installed a recent version of [Git](https://git-scm.com/downloads).

## Introduction

The inway and the outway are able to log metadata of the requests they process. In this guide you will learn how to write transaction logs from the outway service.

## Docker compose

To make the setup of the transaction logs as easy as possible, we have provided a `docker-compose.yml` file in the [nlx-compose repository](https://gitlab.com/commonground/nlx/nlx-compose). It contains configuration for a PostgreSQL container and executes the migrations necessary for the txlog-db.

## Installation

Start by cloning the [nlx-compose repository](https://gitlab.com/commonground/nlx/nlx-compose).

```bash
git clone https://gitlab.com/commonground/nlx/nlx-compose.git
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
docker-compose exec postgres psql -d txlog-db -x -c "SELECT * FROM transactionlog.records ORDER BY id DESC;"

(0 rows)
```

As you can see there are nog logs yet.

## Connect an outway to the transaction log

The next step is to connect an inway or an outway to the database. 
If you've followed the [Try NLX](../try-nlx/setup-your-environment.md) guide you will either have an inway or an outway running. 
In this example we will connect the transaction log to the outway setup in the [Try NLX](../try-nlx/introduction.md) guide.

If your outway from the [Try NLX](../try-nlx/introduction.md) guide is still running, you will have to stop it first.

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

There are some key differences in the Docker command compared to the one used to start the outway in [Try NLX](../try-nlx/introduction.md).

We will go over them briefly

* We removed the flag `--env DISABLE_LOGDB=1`, this flag disables the transaction-log on the outway and defeats the purpose of this guide.
* We added the flag `--env POSTGRES_DSN='postgres://postgres:postgres@postgres/txlog-db?sslmode=disable&connect_timeout=10'`, this flags stores the postgres connection string in the environment variable `POSTGRES_DSN` on the docker container. The outway uses this connection string to setup a connection to the transaction-log database.
* We added the flag `--network nlx`, this connects the docker container to the same docker network as the transaction log database (this network is created by the docker-compose-file we executed earlier). Without this parameter the outway wouldn't be able to connect to the transaction log.


## Verify

Let's make a request through our outway and check if the log records are being writen.

execute

```bash
curl http://localhost/RvRD/basisregistratie/natuurlijke_personen/da02a3ac-4412-11e9-b210-d663bd873d93
```

If successful the response should be

```json
{
  "aanduiding_naamsgebruik": "V",
  "aanschrijving": {
    "adelijke_titel": "",
    "geslachtsnaam": "Vlasman",
    "voorletters": "S.",
    "voornamen": "Sanne",
    "voorvoegsel_geslachtsnaam": ""
  },
  "burgerservicenummer": "663678651",
  "emailadres": "SanneVlasman@gmail.com",
  "geboorte": {
    "datum": "05/07/1970",
    "land": "Nederland",
    "stad": "Utrecht"
  },
  "identificatie": "da02a3ac-4412-11e9-b210-d663bd873d93",
  "kinderen": [
    {
      "url": "/natuurlijke_personen/da02f050-4412-11e9-b210-d663bd873d93"
    },
    {
      "url": "/natuurlijke_personen/da02f1ae-4412-11e9-b210-d663bd873d93"
    }
  ],
  "naam": {
    "adelijke_titel": "",
    "geslachtsnaam": "Vlasman",
    "voorletters": "S.",
    "voornamen": "Sanne",
    "voorvoegsel_geslachtsnaam": ""
  },
  "ouders": [],
  "overlijden": {
    "datum": "",
    "land": "",
    "stad": ""
  },
  "partner": {},
  "postadres": {
    "huisnummer": 21,
    "postcode": "3512JC",
    "straatnaam": "Domplein",
    "woonplaats": "Utrecht"
  },
  "telefoonnummer": "(06)594-38-045",
  "url": "/natuurlijke_personen/da02a3ac-4412-11e9-b210-d663bd873d93",
  "verblijfadres": {
    "huisnummer": 21,
    "postcode": "3512JC",
    "straatnaam": "Domplein",
    "woonplaats": "Utrecht"
  }
}
```

Now the outway should have written records to the transaction log database. Let's query the transaction log API, to verify the log records were written.

```bash
docker-compose exec postgres psql -d txlog-db -x -a -c "SELECT * FROM transactionlog.records ORDER BY id DESC;"

-[ RECORD 1 ]-----+-------------------------------------------------------------------------------
id                | 6
direction         | out
created           | 2020-10-05 14:00:17.415985+00
src_organization  | nlx-test
dest_organization | RvRD
service_name      | basisregistratie
logrecord_id      | 3929112adcf36aa8eed3049adea3ba79
data              | {"request-path": "natuurlijke_personen/da02a3ac-4412-11e9-b210-d663bd873d93"}
```

Congratulations, you've connected your outway to the transaction log!

---
title: "Start an inway"
description: ""
weight: 120
menu:
  docs:
    parent: "providing-services-on-nlx"
---

To provide a service to the NLX network you first need to deploy an NLX inway. There are multiple ways to run an inway. In this guide we'll show you how to run the inway using docker.

## Start an inway using Docker

First make sure you have installed a recent version of [Docker](https://www.docker.com) on your machine. Also make sure you generated a private key and aquired a signed certificate by following [these steps](../../preparing/certificates).

The `service-config.toml` file configures which services are available through the inway. Example:

```toml
[services]

## This block defines a services exposed by this inway.
## A single inway can expose multiple services, therefore this block can be added multiple times.
## The name of the service (in this example PostmanEcho) must be unique for each block.
	[services.MyPublicService]

	## `address` should be set to the address at which the service/API is available.
	## In this example we simply expose the postman-echo.com website.
	address = "<< the address of your local service, e.g.: localhost:8080 >>"

	## `documentation-url` points to the documentation for provided sevice
	documentation-url = "<< URL to online documentation for your service >>"

	## `authorization-model` can or whitelist
	authorization-model = "none"


	## This second service is just an example, and therefore commented out.
	## In this example we add a whitelist authorization model.
	#[services.MyPrivateService]
	#address = "https://postman-echo.com/"
	#documentation-url = "https://docs.postman-echo.com/"

	## We will enable whitelist authorization for this service
	#authorization-model = "whitelist"

	## `authorization-whitelist` is required when `authorization-model` is set to "whitelist".
	## This is a list of organization names (as specified in the peers organization cert) which is allowed access.

	## WARNING: The currently deployed online NLX network is for demo purposes and not ready for connected resources containing sensitive data.
	## When using real personal data, use your own NLX network in an environment you control.

	#authorization-whitelist = ["DemoRequesterOrganization"]
```

where **MyPublicService** is the name of the service. Please note when using default Docker networking settings `localhost` points to the inway Docker container itself. When you run a service on the Docker host, please use the special Docker DNS name: `host.docker.internal`. NLX currently supports HTTP/1.x services like REST/JSON and SOAP/XML. HTTP/2 (gRPC) is work in progress.

## Starting the inway

When you configured the services, start the inway:

```bash
docker pull nlxio/inway:latest

docker run --detach \
  --name my-nlx-inway \
  --volume {/absolute/path/to/root.crt}:/certs/root.crt \
  --volume {/absolute/path/to/org.crt}:/certs/org.crt \
  --volume {/absolute/path/to/org.key}:/certs/org.key \
  --volume {/absolute/path/to/service-config.toml}:/service-config.toml \
  --env DIRECTORY_ADDRESS=directory.demo.nlx.io:1984 \
  --env SELF_ADDRESS={external-inway-hostname-or-ip-address}:4081 \
  --env SERVICE_CONFIG=/service-config.toml \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --publish 4081:80 \
  nlxio/inway:latest
```

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`.

The inway now connects itself to the NLX network and registers its services on the NLX networks. Please **make sure** external connection is possible to the specified port on the specified hostname or IP adress and port  public IP address are routed to the machine running the NLX inway otherwise connections to your inway and services will fail.

To stop the inway run `docker stop my-nlx-inway && docker rm my-nlx-inway`.

## Querying the service

All organizations on the NLX network can query your service on their outway as follows:

```bash
curl http://{some-outway-address}/{organization-name}/{service-name}/{endpoint}
```

## Authorisation

A service is able to implement its own autorization logic by using an NLX-specific header. The inway will automatically append the following header with every request:

```http
X-NLX-Requester-Organization: {organization-name}
```

where `{organization-name}` is set to the organisation performing the request.

NOTE: The organization name is currently free-form and manually entered. We will probably switch to something like OID's in the future.

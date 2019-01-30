---
title: "Start an inway"
description: ""
weight: 120
menu:
  docs:
    parent: "providing-services-on-nlx"
---

To provide a service to the NLX network you first need to deploy an NLX inway. There are multiple ways to run an inway. In this guide we'll show you how to run the inway using docker. This guide is written for Mac/Linux users and assumes you have some experience with the terminal/shell. Windows 10 users may be able to follow this tutorial using [ubuntu for windows](https://tutorials.ubuntu.com/tutorial/tutorial-ubuntu-on-windows). However, the windows platform is not officially supported, we advise to use a VM with Ubuntu installed.

## Start an inway using Docker

First make sure you have installed a recent version of [Docker](https://www.docker.com) on your machine. To run the inway, we need a certificate for encrypted/TLS communications with other NLX services on the network. So we need make sure you generated a private key and aquired a signed certificate by following [these steps](../../preparing/certificates).

The `service-config.toml` file configures which services are available through the inway. Example:

```toml
[services]

# This block defines a services exposed by this inway.
# A single inway can expose multiple services, therefore this block can be added multiple times.
# The name of the service (in this example PostmanEcho) must be unique for each block.
	[services.MyPublicService]

	# `endpoint-url` should be set to the address at which the service/API is available.
	# In this example we simply expose the postman-echo.com website.
	endpoint-url = "<< the address of your local service, e.g.: localhost:8080 >>"

	# `documentation-url` points to the documentation for provided sevice
	documentation-url = "<< URL to online documentation for your service >>"

	# `authorization-model` can or whitelist
	authorization-model = "none"

	# OpenAPI2/3 specification can be provided to the directory. This will allow the directory to render the documentation in the webinterface.
	# This configuration value is optional.
	api-specification-document-url = "https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v2.0/json/petstore.json"


	## This second service is just an example, and therefore commented out.
	## In this example we add a whitelist authorization model.
	#[services.MyPrivateService]
	#endpoint-url = "https://postman-echo.com/"
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
docker pull nlxio/inway:v0.0.29

docker run --detach \
  --name my-nlx-inway \
  --volume {/absolute/path/to/root.crt}:/certs/root.crt:ro \
  --volume {/absolute/path/to/org.crt}:/certs/org.crt:ro \
  --volume {/absolute/path/to/org.key}:/certs/org.key:ro \
  --volume {/absolute/path/to/service-config.toml}:/service-config.toml:ro \
  --env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443 \
  --env SELF_ADDRESS={external-inway-hostname-or-ip-address}:4443 \
  --env SERVICE_CONFIG=/service-config.toml \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --publish 4443:443 \
  nlxio/inway:v0.0.29
```

We give docker several arguments:

- `--detach` will make the container run in the background and print the container ID
- `--name my-nlx-inway` the name of your docker container 
- `--volume {/absolute/path/to/root.crt}:/certs/root.crt:ro` tells docker to make the `/certs/root.crt` file available inside the container.
- `--volume {/absolute/path/to/org.crt}:/certs/org.crt:ro` tells docker to make the `/certs/org.crt` file available inside the container.
- `--volume {/absolute/path/to/org.key}:/certs/org.key:ro` tells docker to make the `/certs/org.key` file available inside the container.
- `--env DDIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443` sets the environment variable `DDIRECTORY_REGISTRATION_ADDRESS` this address is used by the inway to anounce itself to the directory.
- `--env SELF_ADDRESS={external-inway-hostname-or-ip-address}:4443` sets the environment variable `SELF_ADDRESS` to the address of the inway so it can be reached by the NLX network.
- `-env SERVICE_CONFIG=/service-config.toml` sets the environment variable `SERVICE_CONFIG` this is the location of the service-config.toml file which specifies the services conntected to the inway.
- `--env TLS_NLX_ROOT_CERT=/certs/root.crt`sets the environment variable `TLS_NLX_ROOT_CERT` this is the location of the root certificate.
- `--env TLS_ORG_CERT=/certs/org.crt` sets the environment variable `TLS_ORG_CERT` this is the location of the organisation certificate.
- `--env TLS_ORG_KEY=/certs/org.key` sets the environment variable `TLS_ORG_KEY` this is the location of the organisation private key.
- `--env DISABLE_LOGDB=1` sets the environment variable `DISABLE_LOGDB` the value 1 will disable the transaction logs, the value 0 will enable them.
- `--publish 4443:443` connects port 4443 on the host machine to port 443 inside the container. This way, we can send requests to the inway.
- ` nlxio/inway:v0.0.29` is the name of our docker image (`nlxio/inway`) as stored in the docker registry and the version we want to use (`v0.0.29`).

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`. 

The inway is now running and listening on `https://localhost:4443`.

To verify your container is running execute the `docker ps` command, this will print the containers that are currently running. If your container is not running try to start it again without the `--detach` flag, the logs of the container will now be printed in your terminal. 

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

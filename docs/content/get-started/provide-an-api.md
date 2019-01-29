---
title: "Part 4: Provide an API"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Introduction

To provide an API to the NLX network, you need to route traffic through an **inway** service.
We will use the certificate which we've setup in [part 2]({{< ref "/create-certificates.md" >}}), to make sure traffic is encrypted between your and other nodes.

## The API

In this example we will use [postman-echo.com](https://postman-echo.com) as an example API.

We have to define our API in a TOML-file. You can save the contents below as `service-config.toml`.

```toml
[services]
    
# This block defines an API exposed by this inway.
# A single inway can expose multiple API's, therefore this block can be added multiple times.
# The name of the API (in this example PostmanEcho) must be unique for each block.
    [services.MyPublicAPI]

    # `endpoint-url` should be set to the address at which the API is available.
    # In this example we simply expose the postman-echo.com website.
    endpoint-url = "postman-echo.com"

    # `documentation-url` points to the documentation for provided API
    documentation-url = "postman-echo.com"

    # `authorization-model` can or whitelist
    authorization-model = "none"
```

> If you are providing your own API using Docker, make sure to specify the IP-address of the host machine as endpoint-url. 
Localhost won't work, because it is not available to the outside world.

## Setting up the inway

Let's setup the inway service. First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).
    
```bash
docker pull nlxio/inway:v0.0.29
```

The following command will run the inway using the Docker image we just fetched.

```bash
docker run --detach \
              --name my-nlx-inway \
              --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
              --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
              --volume ~/nlx-setup/org.key:/certs/org.key:ro \
              --volume ~/nlx-setup/service-config.toml:/service-config.toml:ro \
              --env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443 \
              --env SELF_ADDRESS=an-awesome-organization.nl:4443 \
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
- `--env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443` sets the environment variable `DIRECTORY_ADDRESS` this address is used by the inway to anounce itself to the directory.
- `--env SELF_ADDRESS={external-inway-hostname-or-ip-address}:4443` sets the environment variable `SELF_ADDRESS` to the address of the inway so it can be reached by the NLX network.
- `-env SERVICE_CONFIG=/service-config.toml` sets the environment variable `SERVICE_CONFIG` this is the location of the service-config.toml file which specifies the services conntected to the inway.
- `--env TLS_NLX_ROOT_CERT=/certs/root.crt`sets the environment variable `TLS_NLX_ROOT_CERT` this is the location of the root certificate.
- `--env TLS_ORG_CERT=/certs/org.crt` sets the environment variable `TLS_ORG_CERT` this is the location of the organisation certificate.
- `--env TLS_ORG_KEY=/certs/org.key` sets the environment variable `TLS_ORG_KEY` this is the location of the organisation private key.
- `--env DISABLE_LOGDB=1` sets the environment variable `DISABLE_LOGDB` the value 1 will disable the transaction logs, the value 0 will enable them.
- `--publish 4443:443` connects port 4443 on the host machine to port 443 inside the container. This way, we can send requests to the inway.
- ` nlxio/inway:v0.0.29` is the name of our docker image (`nlxio/inway`) as stored in the docker registry and the version we want to use (`v0.0.29`).

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`. 

You will get back the container id of the container you created from this image.
By running this command, we've launched our very own NLX inway. It is running and listening on `http://localhost:4443`.

The inway now connects itself to the NLX network and registers its API's on the NLX networks.

Please **make sure** external connection is possible to the specified port on the specified hostname or IP adress and port.
Public IP address are routed to the machine running the NLX inway otherwise connections to your inway and API's will fail.

### Verification

To verify the container is running, execute:

```bash
docker ps
```

You should see your container id in the list of containers. The image name  of the container should be `nlxio/inway:v0.0.29`.

Take a look at the [directory](https://directory.nlx.io) to see if your API is present. It's status should show a green icon.

To verify the inway you just created you will need to use an outway because an inway only accepts requests from an outway. 
Now let's verify our inway is working as expected using the outway we have setup in [part 3]({{< ref "/consume-an-api.md" >}}).

```bash
curl http://localhost:4080/an-awesome-organization/DocsTestMyPublicAPI/
```

geeft : nlx outway: unknown service?

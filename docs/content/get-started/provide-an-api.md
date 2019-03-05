---
title: "Part 4: Provide an API"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Introduction

To provide an API to the NLX network, you need to route traffic through an **inway** service.
We will use the certificate which we've setup in [part 2](../create-certificates), to make sure traffic is encrypted between your and other nodes.

Please note that:

* **you need a domain name to provide an inway** (an IP address won't work)
* **the domain should be the same as the domain you used to generate the certificates** (that was in [part 2](../create-certificates)).

It is not recommended to follow this guide when you're working from your home network. 
Preferably you are able to start the inway service on a machine which is publicly accessible and it's port 4443 is open to the public.

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
    endpoint-url = "https://postman-echo.com/"

    # `documentation-url` points to the documentation for provided API
    documentation-url = "https://docs.postman-echo.com/"

    # `authorization-model` can or whitelist
    authorization-model = "none"

    # `public-support-contact` contains an email address which NLX users can contact if they need support using your API.
    # This email address is published in the directory.
    # this field is optional
    # public-support-contact = "support@my-organization.nl"

    # `tech-support-contact` contains an email address which we (the NLX organization) can contact if they have any questions about your API
    # This email address will NOT be published in the directory
    # this field is optional
    # tech-support-contact   = "tech@my-organization.nl"
```

> If you're specifying you own API, please note that `localhost` won't work. If your API is running on the same machine as 
your inway, the endpoint URL should be your network IP and not `localhost`.   

## Setting up the inway

Let's setup the inway service. First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).
    
```bash
docker pull nlxio/inway:latest
```

The following command will run the inway using the Docker image we just fetched.

```bash
docker run --detach \
              --name my-nlx-inway \
              --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
              --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
              --volume ~/nlx-setup/org.key:/certs/org.key:ro \
              --volume ~/nlx-setup/service-config.toml:/service-config.toml:ro \
              --env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443 \
              --env SELF_ADDRESS=my-organization.nl:4443 \
              --env SERVICE_CONFIG=/service-config.toml \
              --env TLS_NLX_ROOT_CERT=/certs/root.crt \
              --env TLS_ORG_CERT=/certs/org.crt \
              --env TLS_ORG_KEY=/certs/org.key \
              --env DISABLE_LOGDB=1 \
              --publish 4443:443 \
              nlxio/inway:latest
```

We give docker several arguments:

- `--detach` will make the container run in the background and print the container ID
- `--name my-nlx-inway` the name of your docker container 
- `--volume {/absolute/path/to/root.crt}:/certs/root.crt:ro` tells docker to make the `/certs/root.crt` file available inside the container.
- `--volume {/absolute/path/to/org.crt}:/certs/org.crt:ro` tells docker to make the `/certs/org.crt` file available inside the container.
- `--volume {/absolute/path/to/org.key}:/certs/org.key:ro` tells docker to make the `/certs/org.key` file available inside the container.
- `--env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443` sets the environment variable `DIRECTORY_REGISTRATION_ADDRESS` this address is used by the inway to anounce itself to the directory.
- `--env SELF_ADDRESS={external-inway-hostname-or-ip-address}:4443` sets the environment variable `SELF_ADDRESS` to the address of the inway so it can be reached by the NLX network.
- `--env SERVICE_CONFIG=/service-config.toml` sets the environment variable `SERVICE_CONFIG` this is the location of the service-config.toml file which specifies the services connected to the inway.
- `--env TLS_NLX_ROOT_CERT=/certs/root.crt`sets the environment variable `TLS_NLX_ROOT_CERT` this is the location of the root certificate.
- `--env TLS_ORG_CERT=/certs/org.crt` sets the environment variable `TLS_ORG_CERT` this is the location of the organisation certificate.
- `--env TLS_ORG_KEY=/certs/org.key` sets the environment variable `TLS_ORG_KEY` this is the location of the organisation private key.
- `--env DISABLE_LOGDB=1` sets the environment variable `DISABLE_LOGDB` the value 1 will disable the transaction logs, the value 0 will enable them.
- `--publish 4443:443` connects port 4443 on the host machine to port 443 inside the container. This way, we can send requests to the inway.
- ` nlxio/inway:latest` is the name of our docker image (`nlxio/inway`) as stored in the docker registry and the version we want to use (`latest`).

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`. 

You will get back the container id of the container you created from this image.
By running this command, we've launched our very own NLX inway. It is running and listening on `http://localhost:4443`.

The inway now connects itself to the NLX [directory](https://directory.nlx.io).

### Verification

To verify the container is running, execute:

```bash
docker ps
```

You should see your container id in the list of containers. The image name  of the container should be `nlxio/inway:latest`.

Take a look at the [directory](https://directory.nlx.io) to see if your API is present. It's status should show a green icon.

## Querying your own API's

Now let's try to fetch some data from our inway using our outway.
To do so, we have to use the following structure:

```bash
curl http://localhost:4080/my-organization/DocsTestMyPublicAPI/get?foo1=bar1
```

### Verification

The response of the `curl` command should look similar to the following output.

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

Congratulations, you can now consider yourself a member of the NLX club!

## In sum

You have provided your API to the NLX network. In this part, you have:

- setup & configured the NLX inway.
- used your outway to consume your own inway API.

That's all folks! There's some more advanced concepts which you can explore from these docs.

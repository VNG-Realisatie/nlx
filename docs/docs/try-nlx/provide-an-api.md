---
id: provide-an-api
title: Provide an API
---

## Introduction

To provide an API to the NLX network, you need to route traffic through an **inway** service.
To make sure traffic is encrypted between your and other nodes, we will use the certificate which we've setup in [part 2](./retrieve-a-demo-certificate.md).

Please note that:

* You need a domain name to provide an inway (an IP address will not work).
* The domain should be the same as the domain you used to generate the certificates (that was in [part 2](./retrieve-a-demo-certificate.md)).

It is not recommended to follow this guide when you're working from your home network.
Preferably, you are able to start the inway service on a machine which is publicly available. Make sure and the port of the inway (we recommend using port 443) is open to the public.

## The API

In the following example we will use [Swagger Petstore](https://petstore.swagger.io) as an example API.

We have to define our API in a TOML-file. You can save the contents below as `service-config.toml`. Please read our documentation about the [service configuration](../reference-information/service-configuration.md) to learn more about the configuration options.

```toml
[services]

# This block defines an API exposed by this inway.
# A single inway can expose multiple API's, therefore this block can be added multiple times.
# The name of the API (in this example SwaggerPetstore) must be unique for each block.
    [services.SwaggerPetstore]

    # `endpoint-url` should be set to the address at which the API is available.
    # In this example we expose the petstore.swagger.io website.
    endpoint-url = "https://petstore.swagger.io"

    # `documentation-url` should point to the documentation for the API
    documentation-url = "https://petstore.swagger.io"

    # `authorization-model` can be whitelist or none (allow all)
    authorization-model = "none"

    # `api-specification-document-url` defines the specification document for the API.
    # When using the `ca-cert-path` option, the server behind this URL should provide a certificate signed by that root certificate.
    api-specification-document-url = "https://petstore.swagger.io/swagger.json"

    # `ca-cert-path` can be used when the server behind the endpoint url is providing a TLS certificate signed by a custom root certificate.
    # ca-cert-path = "/path/to/custom-root-ca.crt"

    # `public-support-contact` contains an email address which NLX users can contact if they need support using your API.
    # This email address is published in the directory.
    # this field is optional
    # public-support-contact = "support@my-organization.nl"

    # `tech-support-contact` contains an email address which we (the NLX organization) can contact if they have any questions about your API
    # This email address will NOT be published in the directory
    # this field is optional
    # tech-support-contact = "tech@my-organization.nl"
```

> If you're specifying you own API, please note that `localhost` won't work. If your API is running on the same machine as
your inway, the endpoint URL should be your network IP and not `localhost`.

## Setting up the inway

Let's setup the inway service. First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).

```bash
docker pull nlxio/inway:latest
```

The following command will run the inway using the Docker image we just fetched.
<!--DOCUSAURUS_CODE_TABS-->
<!--Linux & macOS-->
```bash
docker run --rm \
  --name my-nlx-inway \
  --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
  --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
  --volume ~/nlx-setup/org.key:/certs/org.key:ro \
  --volume ~/nlx-setup/service-config.toml:/service-config.toml:ro \
  --env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443 \
  --env SELF_ADDRESS=my-organization.nl:443 \
  --env SERVICE_CONFIG=/service-config.toml \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --publish 443:8443 \
  nlxio/inway:latest
```
<!--Windows-->
```powershell
docker run --rm `
--name my-nlx-inway `
--volume ~/nlx-setup/root.crt:/certs/root.crt:ro `
--volume ~/nlx-setup/org.crt:/certs/org.crt:ro `
--volume ~/nlx-setup/org.key:/certs/org.key:ro `
--volume ~/nlx-setup/service-config.toml:/service-config.toml:ro `
--env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443 `
--env SELF_ADDRESS=my-organization.nl:443 `
--env SERVICE_CONFIG=/service-config.toml `
--env TLS_NLX_ROOT_CERT=/certs/root.crt `
--env TLS_ORG_CERT=/certs/org.crt `
--env TLS_ORG_KEY=/certs/org.key `
--env DISABLE_LOGDB=1 `
--publish 443:8443 `
nlxio/inway:latest
```
<!--END_DOCUSAURUS_CODE_TABS-->

We give Docker several arguments:

- `--detach` will make the container run in the background and print the container ID
- `--name my-nlx-inway` the name of your docker container
- `--volume {/absolute/path/to/root.crt}:/certs/root.crt:ro` tells docker to make the `/certs/root.crt` file available inside the container.
- `--volume {/absolute/path/to/org.crt}:/certs/org.crt:ro` tells docker to make the `/certs/org.crt` file available inside the container.
- `--volume {/absolute/path/to/org.key}:/certs/org.key:ro` tells docker to make the `/certs/org.key` file available inside the container.
- `--env DIRECTORY_REGISTRATION_ADDRESS=directory-registration-api.demo.nlx.io:443` sets the environment variable `DIRECTORY_REGISTRATION_ADDRESS` this address is used by the inway to anounce itself to the directory.
- `--env SELF_ADDRESS={external-inway-hostname-or-ip-address}:443` sets the environment variable `SELF_ADDRESS` to the address of the inway so it can be reached by the NLX network.
- `--env SERVICE_CONFIG=/service-config.toml` sets the environment variable `SERVICE_CONFIG` this is the location of the service-config.toml file which specifies the services connected to the inway.
- `--env TLS_NLX_ROOT_CERT=/certs/root.crt`sets the environment variable `TLS_NLX_ROOT_CERT` this is the location of the root certificate.
- `--env TLS_ORG_CERT=/certs/org.crt` sets the environment variable `TLS_ORG_CERT` this is the location of the organization certificate.
- `--env TLS_ORG_KEY=/certs/org.key` sets the environment variable `TLS_ORG_KEY` this is the location of the organization private key.
- `--env DISABLE_LOGDB=1` sets the environment variable `DISABLE_LOGDB` the value 1 will disable the transaction logs, the value 0 will enable them.
- `--publish 443:8443` connects port 443 on the host machine to port 8443 inside the container. This way, we can send requests to the inway.
- ` nlxio/inway:latest` is the name of our docker image (`nlxio/inway`) as stored in the docker registry and the version we want to use (`latest`).

To get started quickly, we will disable transaction logs for now by setting the environment variable `DISABLE_LOGDB=1`.

You will get back the container id of the container you created from this image.
By running this command, we've launched our very own NLX inway. It is running and listening on `https://localhost:443`.

The inway now connects itself to the NLX [directory](https://directory.nlx.io).

### Verification

To verify the container is running, execute:

```bash
docker ps
```

You should see your container id in the list of containers. The image name  of the container should be `nlxio/inway:latest`.

Take a look at the [directory](https://directory.nlx.io) to see if your API is present. Its status should show a green icon.

> The directory will remove stale services automatically. A service will be marked stale when it is not announced for more than 24 hours.

## Querying your own API's

Now let's try to fetch some data from our inway using our outway using `curl`:

```bash
curl http://localhost/my-organization/SwaggerPetstore/v2/pet/20002085
```

### Verification

The response of the `curl` command should look similar to the following output.

```json
{
  "args": {},
  "headers": {
    "x-forwarded-proto": "https",
    "host": "petstore.swagger.io",
    "accept": "*/*",
    "accept-encoding": "gzip",
    "user-agent": "curl/7.54.0",
    "x-nlx-logrecord-id": "<arbitrary-logrecord-id>",
    "x-nlx-request-organization": "my-organization",
    "x-forwarded-port": "443"
  },
  "url": "https://petstore.swagger.io/v2/pet/20002085"
}
```

Congratulations, you can now consider yourself a member of the NLX club!

## In sum

You have provided your API to the NLX network. In this part, you have:

- setup & configured the NLX inway.
- used your outway to consume your own inway API.

That's all folks! There's some more advanced concepts which you can explore from these docs.

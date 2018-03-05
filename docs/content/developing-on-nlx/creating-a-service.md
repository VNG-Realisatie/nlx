---
title: "Creating a service"
description: ""
weight: 120
menu:
  docs:
    parent: "developing-on-nlx"
---

## Deploying an inway
To provide a service to the NLX network you first need to deploy an NLX inway. First make sure you installed the latest version of [Docker](https://www.docker.com) on your machine. Also make sure you generated a private key and aquired a certificate. The steps to aquire a certificate are described in [getting started](../).

Now download the certificate of the NLX development CA:

```bash
wget https://certportal.demo.nlx.io/root.crt
```

And store it next to your private key and certificate. Now create a `service-config.toml` and adjust it to the services you would like to offer to the network. Use the following as an example:

**service-config.toml**

    [services]

        [services.my-inway]
        address = "http://{ip-address-to-the-local-service}:{port}/"

Now start a new inway:


```bash
docker run -d \
-v /absolute/path/to/root.crt:/certs/root.crt \
-v /absolute/path/to/{yourhostname}.crt:/certs/inway.crt \
-v /absolute/path/to/{yourhostname}.key:/certs/inway.key \
-v /absolute/path/to/service-config.toml:/service-config.toml \
-e DIRECTORY_ADDRESS=directory.demo.nlx.io:1984 \
-e SELF_ADDRESS={external-inway-hostname-or-ip-address}:2018 \
-e SERVICE_CONFIG=/service-config.toml \
-e TLS_NLX_ROOT_CERT=/certs/root.crt \
-e TLS_ORG_CERT=/certs/inway.crt \
-e TLS_ORG_KEY=/certs/inway.key \
-p 2018:2018 \
nlxio/inway:latest
```

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
	#authorization-whitelist = ["DemoRequesterOrganization"]

```

The inway now connects itself to the NLX network and registers its services on the NLX networks. Please **make sure** external connection is possible to the specified port on the specified hostname or IP adress and port  public IP address are routed to the machine running the NLX inway otherwise connections to your inway and services will fail.

## Querying the service
Now all organizations on the NLX network can query your service on their outway as follows:

```bash
curl http://{outway-ip}:12018/{organization-name}/{service-name}/{endpoint}
```

## Authorisation
A service is able to implement its own autorisation logic by using a NLX-specific header. The inway will automatically append the following header with every request:

    X-NLX-Requester-Organization: {organization-name}

where ```{organization-name}``` is set to the organisation performing the request.

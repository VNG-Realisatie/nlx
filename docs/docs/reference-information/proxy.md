---
id: outway-as-proxy
title: Outway as a proxy
---

## Introduction

**Please note this is an alpha feature**

The outway can run as a proxy. This allows applications to request services on NLX in the following format: `http://service-name.organization-name.nlx.local`. This can be useful when predictability of URLs is important, such as accessing linked data.

## Configuration

To use the outway as a proxy you need to:

* Start the outway as a proxy
* Configure your http client to use the outway as a proxy

### Start the outway as a proxy

You can start the outway as a proxy by setting the environment variable `USE_AS_HTTP_PROXY` to `1` when starting the outway docker image.

In this example the command to start the outway used in the [try nlx section](../try-nlx/introduction.md) has been modified to start the outway as a proxy:

```bash
docker run --rm \
  --name my-nlx-outway \
  --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
  --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
  --volume ~/nlx-setup/org.key:/certs/org.key:ro \
  --env DIRECTORY_INSPECTION_ADDRESS=directory-inspection-api.demo.nlx.io:443 \
  --env TLS_NLX_ROOT_CERT=/certs/root.crt \
  --env TLS_ORG_CERT=/certs/org.crt \
  --env TLS_ORG_KEY=/certs/org.key \
  --env DISABLE_LOGDB=1 \
  --env USE_AS_HTTP_PROXY=1 \
  --publish 80:8080 \
  nlxio/outway:latest
```

### Configure your http client

The HTTP client you will use to address the outway needs to be configured to use the outway as a proxy. This can be done by setting the environment variable `http_proxy` to the address of your outway. If the http client you use does not respect the `http_proxy` variable, please review the documentation of the http client on how to configure it to use a proxy.

Set the environment variable `http_proxy`:

```bash
export http_proxy=http://{outway-ip}
```

## Using the outway as a proxy.

To address a service on the NLX network the URL should have to following structure: `http://service.organization.services.nlx.local`

For example, to query the BRP demo API use:

```bash
curl http://basisregistratie.brp.service.nlx.local/natuurlijke_personen/da02ca58-4412-11e9-b210-d663bd873d93
```

If the URL does not contain `service.nlx.local`, the request will not be routed through the NLX network but through the public internet.
Calls routed through the public internet **will not** be logged in the transaction log of the outway.

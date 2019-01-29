---
title: "Part 3: Consume an API"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Introduction

To use an API that is provided via NLX, you need to route traffic through an **outway** onto the network.
We will use the certificate which we've setup in [part 2]({{< ref "/create-certificates.md" >}}), to make sure traffic is encrypted between your and other nodes.

## Setting up the outway

Let's setup the outway service. First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).
    
```bash
docker pull nlxio/outway:v0.0.20
```

The following command will run the outway using the Docker image we just fetched.

```bash
docker run --detach \
             --name my-nlx-outway \
             --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
             --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
             --volume ~/nlx-setup/org.key:/certs/org.key:ro \
             --env DIRECTORY_ADDRESS=directory-api.demo.nlx.io:443 \
             --env TLS_NLX_ROOT_CERT=/certs/root.crt \
             --env TLS_ORG_CERT=/certs/org.crt \
             --env TLS_ORG_KEY=/certs/org.key \
             --env DISABLE_LOGDB=1 \
             --publish 4080:80 \
             nlxio/outway:v0.0.20
```

You will get back the container id of the container you created from this image.
By running this command, we've launched our very own NLX outway. It is running on `http://localhost:4080`.

### Verification

To verify the container is running, execute:

```bash
docker ps
```

You should see your container id in the list of containers. The image name  of the container should be `nlxio/outway:v0.0.20`.

If the service is not present, it might have stopped for unknown reasons. You can see all your containers including stopped ones using

```bash
docker ps -a
```

To inspect the logs of a (stopped) container, use the following command

```bash
docker logs -f <container-id>
```

## Querying API's

Now let's try to fetch some data from an API in the NLX network!
To do so, we have to use the following structure:

```bash
curl http://localhost:4080/{organization-name}/{service-name}/{api-specific-path}
```

For example, to query the NLX demo application use:

```bash
curl http://localhost:4080/vng-realisatie/demo-api/
```

Congratulations, you now made your first query on the NLX network!

Take a look at the [directory](https://directory.nlx.io) to see which other API's are available to fetch data from.

### Verification

The response of the `curl` command should look similar to the following output.

```json
{
  "message": "Hi there, greetings from the nlx-demo API!", 
  "local_time": "2019-01-22T15:19:13.618494", 
  "nlx_request_organization": "<your-organization-name>", 
  "nlx_request_logrecord_id": null
}
```

## In sum
    
In this part, we have:

- setup a local NLX outway service, which we can use to get data from the network.
- made a real request to the VNG Realisatie Demo API API.

Now let's see if we can provide our own API's to the network in [part 4]({{< ref "/provide-an-api.md" >}}). 

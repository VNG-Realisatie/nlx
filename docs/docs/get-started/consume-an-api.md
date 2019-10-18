---
id: consume-an-api
title: Part 3: Consume an API
---

## Introduction

To use an API that is provided via NLX, you need to route traffic through an **outway** onto the network.
We will use the certificate which we've setup in [part 2](./create-certificates), to make sure traffic is encrypted between your and other nodes.

## Setting up the outway

Let's setup the outway service. First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).

```bash
docker pull nlxio/outway:latest
```

The following command will run the outway using the Docker image we just fetched.
You might need to give `org.key` group read permisson so docker can access the key.

```bash
chmod g+r org.key
```

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
  --publish 80:8080 \
  nlxio/outway:latest
```

You will get back the container id of the container you created from this image.
By running this command, we've launched our very own NLX outway. It is running on [`http://localhost`](http://localhost).

### Verification

To verify the container is running, execute:

```bash
docker ps
```

You should see your container id in the list of containers. The image name  of the container should be `nlxio/outway:latest`.

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
curl http://localhost/{organization-name}/{service-name}/{api-specific-path}
```

For example, to query the Haarlem demo API application use:

```bash
curl http://localhost/haarlem/demo-api/get
```

### Verification

The response of the `curl` command should look similar to the following output.

```json
{
  "args": {},
  "headers": {
    "x-forwarded-proto":"https",
    "host":"postman-echo.com",
    "accept":"*/*",
    "accept-encoding":"gzip",
    "user-agent":"curl/7.61.0",
    "x-nlx-logrecord-id":"<log-record-id>",
    "x-nlx-request-organization":"<your-organization-name>",
    "x-forwarded-port":"443"
  },
  "url":"https://postman-echo.com/get"
}
```

Congratulations, you now made your first query on the NLX network!

API's provided on the NLX network are published in the NLX directory. You can use this directory to see which other API's are available.
Take a look at the [directory](https://directory.nlx.io).

## In sum

In this part, we have:

- setup a local NLX outway, which we can use to get data from the network.
- made a real request to the VNG Realisatie Demo API.

Now let's see if we can provide our own API's to the network.

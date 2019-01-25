---
title: "Part 2: Consuming an API"
description: ""
menu:
  docs:
    parent: "get-started"
---

## Intro - Setting up an outway

To use an API that is provided via NLX, you need to route traffic through an outway onto the network.
To be able to connect with the other nodes on the NLX network, you'll need a certificate and private key. The certificate and key are used to encrypt traffic between you and other nodes.

## Before we can request a certificate

In order to request a certificate, we need to generate a key and Certificate Signing Request (CSR). 
We can create these using [openssl](https://www.openssl.org/).

```bash
openssl req -utf8 -nodes -sha256 -newkey rsa:4096 -keyout org.key -out org.csr
```

Answer the questions accordingly:

- **Country Name**, enter any value
- **State**, enter any value
- **Locality Name**, enter any value
- **Organization Name**, please enter a URL-friendly value. Also make sure this value is unique for the network in the [directory overview](https://directory.demo.nlx.io) as we do not check for uniqueness yet.<br>A good value could be: `an-awesome-organization`.
- **Organization Unit Name**, enter any value
- **Common name**, when you would like to offer your API's to the NLX network make sure this corresponds to your external hostname. For this guide we will use `an-awesome-organization.nl`.
- **Email Address**, enter any value
- **A challenge password**, leave empty

Now openssl wil generate two files:

- A private key `org.key`.  **Keep this file safe**, limit access to it and do not transfer it in an insecure way.
- A certificate request `org.csr`. Use this file to request a certificate.

### Verification

List the files in your working directory.

```bash
ls -la
```

If all of the above went well, you should see at least two files listed:

* org.csr
* org.key

## Request the demo certificate

We will use the NLX certportal to retrieve an NLX developer certificate.

First let's copy the contents of `org.csr`. We will use this to generate the demo certificate.
Make sure to copy the complete output, including *-----BEGIN CERTIFICATE REQUEST-----* and *-----END CERTIFICATE REQUEST-----*.

```bash
cat org.csr
```
 
Open [certportal](https://certportal.demo.nlx.io) and paste the content in the `CSR` field.
 
Scroll to the bottom of the page and click on **Request certificate**. 

The system will instantly sign your csr and return your certificate. 
You can either copy paste your certificate and store it in a file or you can click **Download certificate** to download the certificate. 

Rename the file from `certificate.crt` to `org.crt` and store the file next to your private key.

### Verification

Let's check if our certificate is alright.

```bash
openssl x509 -in org.crt -text | grep Subject:
```

The output should contain the answers you've provided when you created the certificate.

Example of the output: `Subject: C=nl, ST=noord-holland, L=haarlem, O=an-awesome-organization, OU=an-awesome-organization-unit, CN=an-awesome-organization.nl`


## Using the certificate with an outway

All prerequisites are available now. Let's setup the outway service.

First, fetch the Docker image from the [Docker Hub](https://hub.docker.com/u/nlxio).
    
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
    
So far, we have:

- generated our own certificate, so we are allowed to communicate with the API's on the NLX network.
- setup a local NLX outway service, which we can use to get data from the network.
- made a real request to the VNG Realisatie Demo API API.

Now let's see if we can provide our own API's to the network in [part 3]({{< ref "/providing-an-api.md" >}}). 

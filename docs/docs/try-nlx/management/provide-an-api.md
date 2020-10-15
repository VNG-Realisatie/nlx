---
id: provide-an-api
title: Provide an API
---

## Introduction

To provide an API to the NLX network, you need to route traffic through an **inway** service.
To make sure traffic is encrypted between your and other nodes, we will use the certificate which we've setup in [Retrieve a demo certificate ](../retrieve-a-demo-certificate.md).

Please note that:

* You need a domain name to provide an inway (an IP address will not work).
* The domain should be the same as the domain you used to generate the certificates (that was in [Retrieve a demo certificate](../retrieve-a-demo-certificate.md)).

It is not recommended to follow this guide when you're working from your home network.

Preferably, you are able to start the inway service on a machine which is publicly available. 
Make sure the port of the inway (we recommend using port `443`) is open to the public.


## Verification

Assuming you followed [Getting up and running](./getting-up-and-running.md) the inway and Management API should be already running.

You can confirm that by inspecting the Management API logs:

```
docker-compose -f docker-compose.management.yml logs -f api
```

Next, let's take a look at the logs of our inway:

```
docker-compose -f docker-compose.management.yml logs -f inway
```


## Create a service

In the following example we will use [Swagger Petstore](https://petstore.swagger.io) as an example API.

To provide our API in the NLX network we have to create a service in the Management UI.
You can do that by going to the services page where you click on the 'Add service' button.

Note that for demo purposes you can omit most fields only fill in the 'Servicename' and 'API Endpoint URL' field.
For the 'ServiceName' use `SwaggerPetStore` and for the 'API Endpoint URL' use `https://petstore.swagger.io/v2`

Next to that you also have to select the inway `Demo Inway` to be used by this service.

![Add service screen](/img/nlx-management-add-service-screen.png "Add service screen")

Whenever you're ready, click 'Add service' to save the details and register the service in the demo directory.

> If you're specifying you own API, please note that `localhost` will not work. If your API is running on the same machine as
your inway, the endpoint URL should be your network IP and not `localhost`.

After adding the service, you should see the service in the services page and in the directory.

> The directory will remove stale services automatically. A service will be marked stale when it is not announced for more than 24 hours.


## Querying your own API's

Now let's try to fetch some data from our inway using our outway using `curl`:

```bash
curl http://localhost/my-organization/SwaggerPetstore/v2/pet/20002085
```

The response of the `curl` command should look similar to the following output (where `ORGANIZATION_NAME`/`PUBLIC_KEY_FINGERPRINT` are derived from the certificate generated in [step 3](../retrieve-a-demo-certificate.md)).

```
nlx-inway: permission denied, organization "ORGANIZATION_NAME" or public key "PUBLIC_KEY_FINGERPRINT" is not allowed access.
```

We are denied access because we first need to request access. This is one of the key features of NLX Management!

In order to request access, follow these steps:

1. Navigate to the 'Directory' in NLX Management. 
1. Select the service `SwaggerPetstore` from the list and click on 'Toegang aanvragen'.
1. Now navigate to the 'Services' page and again select the service `SwaggerPetstore`.
1. You should see one access request under the section 'Toegansverzoeken'. 
1. Expand the section and click on 'Accepteren' to accept the access request. 
1. You now have an access grant for the service.

Let's try to fetch the data again.

```bash
curl http://localhost/my-organization/SwaggerPetstore/v2/pet/20002085
```

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


Congratulations, you can now consider yourself a member of the NLX club! 🎉

## In sum

You have provided your API to the NLX network. In this part, you have:

- Created a service for the demo API
- Used your outway to consume your own inway API.

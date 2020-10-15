---
id: consume-an-api
title: Consume an API
---

## Introduction

To use an API that is provided via NLX, you need to route traffic through an **outway** onto the network.
We will use the certificate which we've setup in [Retrieve a demo certificate ](../retrieve-a-demo-certificate.md), to make sure traffic is encrypted between your and other nodes.


### Verification

Assuming you followed [Getting up and running](./getting-up-and-running.md) the outway should already be running.
You can confirm that by checking the outway logs:

```
docker-compose -f docker-compose.management.yml logs -f outway
```


## Querying APIs

Now let's try to fetch some data from an API in the NLX network!
To do so, we have to use the following structure:

```bash
curl http://localhost/{organization-name}/{service-name}/{api-specific-path}
```

For example, to query the BRP demo API use:

```bash
curl http://localhost/BRP/basisregistratie/natuurlijke_personen/da02ca58-4412-11e9-b210-d663bd873d93
```

You can also run the outway as a HTTP proxy. This allows applications to call services on NLX by using `http://service-name.organization-name.services.nlx.local`.
For more information read [the reference information](../../reference-information/proxy.md)

The response of the `curl` command should look similar to the following output.

```json
{
  "aanduiding_naamsgebruik": "V",
  "aanschrijving": {
    "adelijke_titel": "",
    "geslachtsnaam": "Poll",
    "voorletters": "L.",
    "voornamen": "Linneke",
    "voorvoegsel_geslachtsnaam": ""
  },
  "burgerservicenummer": "58249163",
  "emailadres": "LinnekePoll@gmail.com",
  "geboorte": {
    "datum": "25/11/1984",
    "land": "Nederland",
    "stad": "Utrecht"
  },
  "identificatie": "da02ca58-4412-11e9-b210-d663bd873d93",
  "kinderen": [],
  "naam": {
    "adelijke_titel": "",
    "geslachtsnaam": "Poll",
    "voorletters": "L.",
    "voornamen": "Linneke",
    "voorvoegsel_geslachtsnaam": ""
  },
  "ouders": [],
  "overlijden": {
    "datum": "",
    "land": "",
    "stad": ""
  },
  "partner": {},
  "postadres": {
    "huisnummer": 11,
    "postcode": "3521AZ",
    "straatnaam": "Stadsplateau",
    "woonplaats": "Utrecht"
  },
  "telefoonnummer": "(06)432-51-968",
  "url": "/natuurlijke_personen/da02ca58-4412-11e9-b210-d663bd873d93",
  "verblijfadres": {
    "huisnummer": 11,
    "postcode": "3521AZ",
    "straatnaam": "Stadsplateau",
    "woonplaats": "Utrecht"
  }
}
```

Congratulations!, You have made your first query on the NLX network!

APIs provided on the NLX network are published in the NLX directory.
Take a look at the [directory](https://directory.nlx.io) to see which APIs are available.

## In sum

In this part, we have:

- Setup a local NLX outway, which we can use to get data from the network.
- Made a real request to the VNG Realisatie Demo API.

Now let's see if we can [provide our own APIs](./provide-an-api.md) to the network.

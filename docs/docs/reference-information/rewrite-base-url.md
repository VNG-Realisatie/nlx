---
id: rewrite-base-url
title: Rewrite Base URLs
---

## Problem

Some API responses contain hardcoded URIs pointing back to the API or another resource.
This could pose a problem for the client when following actions are required based on
these URIs when the new request again needs to go through NLX.

### Example

The [Daara API](https://demo.ldproxy.net/daraa) includes the base URL of the API in its response data.
For example, the [GET /collections](https://demo.ldproxy.net/daraa/collections/?f=json) endpoint.

```json
{
  "title": "Daraa",
  "description": "This is a test dataset for the Open Portrayal Framework thread in the OGC Testbed-15 as well as for the OGC Vector Tiles Pilot Phase 2. The data is OpenStreetMap data from the region of Daraa, Syria, converted to the Topographic Data Store schema of NGA.",
  "links": [
    {
      "rel": "self",
      "type": "application/json",
      "title": "This document",
      "href": "https://demo.ldproxy.net/daraa/collections?f=json"
    },
    {
      "rel": "alternate",
      "type": "text/html",
      "title": "This document as HTML",
      "href": "https://demo.ldproxy.net/daraa/collections?f=html"
    }
  ]
}
```

## Possible solution

We have identified this behavior and are investigating the permanent solution. Until then, we suggest to use a proxy
which will replace the base URL of the original API with the base URL of your proxy address.
This proxy will pass the requests to your Outway.

## Reference code

We have created an example setup with Nginx for this proxy. The files are listed on the
[proxy-rewrite-base-url at GitLab](https://gitlab.com/commonground/nlx/proxy-rewrite-base-url).
Instructions on how to build & run the container can be found in the README.

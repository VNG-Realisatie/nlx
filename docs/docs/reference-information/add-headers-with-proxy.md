---
id: add-headers-with-proxy
title: Add HTTP headers with a proxy
---

## Problem

When you make requests on behalf of another organization, you need to add the following HTTP headers:
- `X-NLX-Request-Order-Reference`
- `X-NLX-Request-Delegator`

to the request you make from your client to the Outway.

However, when you're not in control of the client, you might not be able to add these headers. E.g. when the client is made by a third party.

## Solution

We have identified this problem and are investigating a permanent solution. Until then, we suggest to introduce a proxy service.
This service will add the required HTTP headers to the request before passing it to the Outway.

## Reference code

We have created an example setup with Nginx for this proxy. The files are listed on the
[proxy-add-http-headers repository at GitLab](https://gitlab.com/commonground/nlx/proxy-add-http-headers).
Instructions on how to build & run the container can be found in the [README](https://gitlab.com/commonground/nlx/proxy-add-http-headers/-/blob/main/README.md).

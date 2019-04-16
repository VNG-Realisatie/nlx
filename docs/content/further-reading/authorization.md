---
title: "Authorization"
description: ""
weight: 20
menu:
  docs:
    parent: "further-reading"
---


## Introduction

To enable authorization, organizations can plug-in their own authorization service to an outway. Once configured, all requests will be authorized by this service before routing the request to the targeted API in the NLX network.

## How it works

Once an authorization service is configured the outway will, after receiving a request from an application, extract all the HTTP-headers from the request. Only the HTTP headers along with the destination organization and service will be send to the authorization service. The authorization service can use this information to determine if authorization should be granted. The authorization service will send the result back to the the outway. If the authorization is granted the outway will strip the HTTP headers starting with `X-NLX` from the request and continue sending the request (body + HTTP headers) to the destination inway.   
 
## The authorization interface

In order to keep the NLX components as flexible as possible, each organization will have to implement the "authorization interface" on their own authorization service. Implementing this interface will enable communication between the outway and the authorization service.
We have described this interface using Open API Specification (OAS). [This specification can be found here](https://gitlab.com/commonground/nlx/tree/master/outway/authorization-interface.yaml).
A reference implementation has also been made available in the [NLX repository](https://gitlab.com/commonground/nlx/blob/master/auth-service/). 

## Configuring the outway

After you have implemented the `authorization interface` on your authorization service, you will have to configure the outway to use it. This can be done by setting the environment variable `AUTHORIZATION_SERVICE_ADDRESS` in the docker image of the outway. This variable should contain the URL of your authorization service.
For example, the URL of our authorization service is `https://auth.nlx.io`, we can start the outway by running the following docker command:

```bash
docker run --detach \
             --name my-nlx-outway \
             --volume ~/nlx-setup/root.crt:/certs/root.crt:ro \
             --volume ~/nlx-setup/org.crt:/certs/org.crt:ro \
             --volume ~/nlx-setup/org.key:/certs/org.key:ro \
             --env DIRECTORY_INSPECTION_ADDRESS=directory-inspection-api.demo.nlx.io:443 \
             --env TLS_NLX_ROOT_CERT=/certs/root.crt \
             --env TLS_ORG_CERT=/certs/org.crt \
             --env TLS_ORG_KEY=/certs/org.key \
             --env AUTHORIZATION_SERVICE_ADDRESS=https://auth.nlx.io \
             --env AUTHORIZATION_ROOT_CA=~/nlx-setup/root.crt:/certs/root.crt:ro \
             --env DISABLE_LOGDB=1 \
             --publish 4080:80 \
             nlxio/outway:latest
```

Please note we are also setting the environment variable `AUTHORIZATION_ROOT_CA`. This variable contains the path to a root Certificate Authority(CA). To keep everything as secure as possible, your authorization service **must** only accept connections with TLS enabled. The configured root CA will be used by the outway to verify the certificates of your authorization service.

## Headers

HTTP-headers are used to pass around authorization information. NLX provides a set of HTTP-headers which will be logged in the transaction-log and stripped from any request as soon as this request leaves the organization.

- `X-NLX-Requester-User`  should contain information about the user, for example a userID
- `X-NLX-Requester-Claims` should contain claims granted to the user
- `Proxy-Authorization` should contain the credentials to authenticate a user or application with a server. For example, a JWT token. This header will **always** be stripped, even if the request does not leave the organisation. Take a look at the [Mozilla Proxy-Authorization header documentation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authorization) for more information about this header. 





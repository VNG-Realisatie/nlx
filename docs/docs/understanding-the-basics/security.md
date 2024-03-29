---
id: security
title: Security
---

The main security mechanism is based on [TLS client authentication](https://blog.cloudflare.com/introducing-tls-client-auth/). The nodes in the NLX network trust a [Certificate authority](https://en.wikipedia.org/wiki/Certificate_authority) to sign certificates of organizations. Outway and Inway nodes identify themselves by using their signed certificate. All traffic between the Outway and Inway is therefore also encrypted using TLS.

NLX serves as a data-exchange layer between applications and services. It therefore only provides an authentication mechanism on *organizational level*. This means that authentication and authorization on *user level* is the responsibility of the application. A service is able to implement authorization on *organizational level* when required.

## Applications
The proof of concept does not provide a security mechanism for applications yet. Therefore the HTTP endpoint of the Inway is unprotected. In future versions of NLX the Inway endpoint will be protected by TLS client authentication as well. An organization will be able to configure an organization certificate autority. The CA signs certificates per application to control the access. NLX will provide additional configuration options to restrict the access to services per application.

## Services
Services receive traffic from the NLX network through the Inway. The Inway performs authentication on organizational level and attaches a header with the organization name of the requester:

    X-NLX-Requester-Organization: {OrganizationSerialNumber}

With this header a service is able to implement its own authorization logic.

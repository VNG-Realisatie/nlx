# Outway proxy protocol

## Introduction

The NLX outway component proxies traffic from an application, through an inway, to the service (API) endpoint for an internal or external organization. It performs tasks such as logging, applying security (mutual TLS between outway and inway), and user authentication.

This document describes challenges the current design of the NLX outway is facing, and proposes a solution to overcome them.

## History and concerns

The initial design of the Outway, conceived during the PoC-phase of NLX, makes it basically a direct reverse HTTP proxy. Requests are sent from the application, directly to the outway. The application must set the HTTP `Host` header to the outway's address and the `path` value must be prefixed with the organization and service name to which the request must be forwarded.

This worked well, and provided a simple way to get started with building the PoC. However, it also presented a number of questions and concerns.

- How to work with Linked Data?
- What about applications that are not NLX-aware?
- Can we use the outway for API's that are not (yet) provided by an NLX inway?

These are all valid questions, that we want and need to address. But the current outway proxy design doesn't seem to allow for a clean and good solution. Lets dive into each use-case to identify their need and the chalenges they pose.

### Linked Data

Wikipedia summarizes Linked Data as follows.

> In computing, linked data (often capitalized as Linked Data) is a method of publishing structured data so that it can be interlinked and become more useful through semantic queries. It builds upon standard Web technologies such as HTTP, RDF and URIs, but rather than using them to serve web pages only for human readers, it extends them to share information in a way that can be read automatically by computers. Part of the vision of linked data is for the internet to become a global database.

Linked data is often used in API's within the dutch governments. REST API endpoints may refer to objects available from a different service in a different organization. This new service may be inside or outside the initially enquired organization.

### Applications that are not NLX-aweare

Some organizations may use applications that are not actively maintained anymore. These organizations may still want to use NLX. With the current design, there is no proper way to direct the outgoing requests from an application to the NLX outway without modifying the application software. This means these unmaintained application cannot use NLX.

### Using regular API's (no NLX inway) through the outway

Given that orgainzation A uses NLX and organization B does not use NLX. Organization A might still want to log their side of the conversation, even though organization B does not (or uses another means of logging). With the current design, NLX must be used by both organizations.

This use-case is also interesting in conjunction with the Linked Data use-case. A service on NLX may return a reference to an object in an organization that has no NLX inways installed yet.

## Proposed solution

This document describes on a high-level the approach that we may take to make NLX more transparent and compatible with existing systems.

Instead of modifying the URI for an HTTP request to point to the outway, we let applications use the outway as a generic proxy. Most HTTP libraries and cli's support this by setting the HTTP_PROXY and HTTPS_PROXY environment variables, or by passing the proxy address during the initialization of an HTTP client.

The outway accepts any HTTP traffic, also for API's that do not have an inway available. The application make requests without modifying the requested path or Host header.

Instead of (currently): `https://outway.local:8080/voorbeeld-organization/some-api/v1/getContacts`

The request becomes: `https://some-api.voorbeeld-organization.nl/v1/getContacts`

This is semantically clean, and works great with linked-data.

### Discovering NLX-enabled services

The outway will need a way to discover whether an API is NLX-enabled or not. Two solutions come to mind. Either one or both of those solutions can be implemented.

- DNS records
- Directory

#### DNS records

A DNS record indicates that an API is NLX-enabled. For example: `_nlx.some-api.voorbeeld-organization.nl` indicates that some-api.voorbeeld.organization.nl is NLX-enabled.

The record will need to have a key=value structure so that we may configure.

`v=nlx1; org=voorbeeld-haarlem; path:/v1="svc=some-api-v1"; path:/v2="svc=some-api-v2";`

Syntax:

- Semicolons separate the key=value pairs.
- keys must match the regexp: `[a-zA-Z0-9\-\_\:\/]+`.
- values must match the regexp: `[a-zA-Z0-9]`.
- values may be quoted.

Defined keys:

- `v=nlx1` indicates that the record is in fact an NLX record with syntax version 1.
- `org` sets the organization for this domain.
- `path:{path}` 
- - `{path}` must match the regexp `[a-zA-Z\/\-_]+`.
- - The value for a path key contains a wrapped key=value structure, and must be quoted.
- - `svc` sets the service name for the path.
- - Optionally: `in` points to a DNS record of inways. This would remove dependency on directory.
- - The record of inways must be an A record. It may have multiple A values, in which case traffic is evenly distributed accross inways.

**Considerations:**

- What if a domain claims to be organization X service Y? It's important that organization X service Y still matches the Host and HTTP path corectly.

**Pro's:**

- We may be able to do most management via DNS.

**Con's:**

- How does directory discover which services are available?

#### Directory

The directory keeps a list of addresses related to a service.

**Pro's:**

- Easy for outways to match which service to use.

**Con's:**

- Validation required in directory.

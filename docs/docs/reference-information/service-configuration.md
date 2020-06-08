---
id: service-configuration
title: Service configuration
---

A single inway can expose multiple services. You can tell an inway which services to expose by providing the inway a [toml](https://github.com/toml-lang/toml) file which contains the service configuration.

Below is an example configuration named `service-config.toml`.
```toml
version = 2
[services]

# This block defines an service exposed by this inway.
# A single inway can expose multiple services, therefore this block can be added multiple times.
# The name of the API (in this example SwaggerPetstore) must be unique for each block.
    [services.SwaggerPetstore]
    # In this example we expose the petstore.swagger.io website.
    endpoint-url = "https://petstore.swagger.io"
    documentation-url = "https://petstore.swagger.io"
    api-specification-document-url = "https://petstore.swagger.io/swagger.json"
    authorization-model = "whitelist"
    public-support-contact = "support@my-organization.nl"
    tech-support-contact = "tech@my-organization.nl"
    ca-cert-path = "/path/to/custom-root-ca.crt"
    [[services.SwaggerPetstore.authorization-whitelist]]
    organization-name = "DemoRequesterOrganization1"

    [[services.SwaggerPetstore.authorization-whitelist]]
    organization-name = "DemoRequesterOrganization2"
    public-key-hash = "tGbzEuAy88OB0zZWm+dolZoakhIKScV7zTK3wA15Ci8="

    [[services.SwaggerPetstore.authorization-whitelist]]
    public-key-hash = "yX0i/6NJZxaZWw7+LoCoq/vlA+06qb/5j/cg4n/zT/A="
```
# Top level fields

## version
***Required***
Should be set to the version of the config that is used, a deprecation warning is logged when the value is empty or less than 2.

***Example***
```toml
version = 2
```
# Service configuration fields

## endpoint-url
***Required***
Should be set to the address at which the API is available. Please make sure the inway can reach the API on this address!

***Example***
```toml
endpoint-url = "https://petstore.swagger.io"
```

## documentation-url
Should be set to the url at which the documentation for this API is available.

***Example***
```toml
documentation-url = "https://petstore.swagger.io"
```

## api-specification-document-url
If there is an [OpenAPI Specification](https://swagger.io/specification/) (OAS) available for the exposed API you can supply an URL to the OAS in this field. The OAS will be published to the [directory](https://directory.nlx.io).
When using the [ca-cert-path](#field-ca-cert-path) option, the server behind this URL should provide a certificate signed by that root certifictate.
The following OAS versions are supported: 2.0, 3.0.0, 3.0.1, 3.0.2

***Example***
```toml
api-specification-document-url = "https://petstore.swagger.io/swagger.json"
```

<a name="field-authorization-model"></a>

## authorization-model
***Required***
The authorization model tells the inway how to authorise outways who are trying to consume this service.
Currently there are two options available:

1. `none` All outways with a valid NLX certificate can consume this service from the inway. No authorization check will be performed.
1. `whitelist` An outway has to have a valid NLX certificate and the organization name and/or public key fingerprint of this certificate should be present in the [authorization-whitelist](#field-authorization-whitelist) of the inway. If not, the inway will not accept requests from this outway.

***Example***
```toml
authorization-model = "whitelist"
```

<a name="field-authorization-whitelist"></a>

## authorization-whitelist
A whitelist of organizations who are authorized to consume the service. When using the `authorization-whitelist` field the [authorization-model](#field-authorization-model) of the service should be set to `whitelist`.
Each entry in the whitelist consists of an `organization-name` and/or `public-key`:
* `public-key` is the preferred method of authorization as this:
  * This restricts the certificates from a particular organization that can be used to setup NLX connections.
  * This allows organizations to compartiment their security by having different security zones.
  * This protects the NLX system of compromised CA's by pinning to specific public keys of certificates.
* `organization-name` offers backward compatibility with the previous version of the whitelist.

***Example***
```toml
[[services.SwaggerPetstore.authorization-whitelist]]
organization-name = "DemoRequesterOrganization1"

[[services.SwaggerPetstore.authorization-whitelist]]
organization-name = "DemoRequesterOrganization2"
public-key-hash = "tGbzEuAy88OB0zZWm+dolZoakhIKScV7zTK3wA15Ci8="

[[services.SwaggerPetstore.authorization-whitelist]]
public-key-hash = "yX0i/6NJZxaZWw7+LoCoq/vlA+06qb/5j/cg4n/zT/A="
```

***Example v1***

_This syntax is deprecated and will cause a warning to be logged_
```toml
authorization-whitelist = ["DemoRequesterOrganization1", "DemoRequesterOrganization2"]
```

<a name="field-ca-cert-path"></a>

## ca-cert-path
Can be used if the API you are trying to expose is providing a TLS certificate signed by a custom root certificate. The root certificate has to be available on the machine running the inway and the absolute path to the root certificate should be the value of this field.

***Example***
```toml
ca-cert-path = "/path/to/custom-root-ca.crt"`
```

## public-support-contact
Contains an email address which NLX users can contact if they need your support when using this service. This email address is published in the [directory](https://directory.nlx.io).

***Example***
```toml
public-support-contact = "support@my-organization.nl"
```

## tech-support-contact
Contains an email address which we (the NLX organization) can contact if we have any questions about your API.
This email address will NOT be published in the directory.

***Example***
```toml
tech-support-contact = "tech@my-organization.nl"
```

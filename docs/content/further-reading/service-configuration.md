---
title: "Service configuration"
description: ""
weight: 20
menu:
  docs:
    parent: "further-reading"
---


A single inway can expose multiple services to the NLX network, you can tell an inway which services to expose by providing the inway a [toml](https://github.com/toml-lang/toml) file which contains the service configuration.

Below is an example configuration named `service-config.toml`.
```toml
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
    authorization-whitelist = ["DemoRequesterOrganization"]
    public-support-contact = "support@my-organization.nl"
    tech-support-contact = "tech@my-organization.nl"
    # ca-cert-path = "/path/to/custom-root-ca.crt"
```
# Service configuration fields
## endpoint-url 
***required***   
Should be set to the address at which the API is available. Please make sure the inway can reach the API on this address!

`example: endpoint-url = "https://petstore.swagger.io"`
## documentation-url
Should be set to the url at which the documentation for this API is available.

`example: documentation-url = "https://petstore.swagger.io"`
## api-specification-document-url
If there is an [OpenAPI Specification](https://swagger.io/specification/)(OAS) available for the exposed API you can supply an URL to the OAS in this field. The OAS will be published to the [directory](https://directory.nlx.io).
When using the [ca-cert-path](#field-ca-cert-path) option, the server behind this URL should provide a certificate signed by that root certifictate. 
The following OAS versions are supported: 2.0, 3.0.0, 3.0.1, 3.0.2

`example: api-specification-document-url = "https://petstore.swagger.io/swagger.json"`
<a name="field-authorization-model"></a>
## authorization-model
***required***  
The authorization model tells the inway how to authorise outways who are trying to consume this service.
Currently there are two options available:

1. `none` All outways with a valid NLX certificate can consume this service from the inway. No authorization check will be performed.
1. `whitelist` An outway has to have a valid NLX certificate and the organization name in this certificate should be present in the [authorization-whitelist](#field-authorization-whitelist) of the inway. If not, the inway will not accept requests from this outway.

`example: authorization-model = "whitelist"`
<a name="field-authorization-whitelist"></a>
## authorization-whitelist
A whitelist of organizations who are authorized to consume the service. When using the `authorization-whitelist` field the [authorization-model](#field-authorization-model) of the service should be set to `whitelist`.

`example: authorization-whitelist = ["DemoRequesterOrganization1", "DemoRequesterOrganization2"] `
<a name="field-ca-cert-path"></a>
## ca-cert-path 
Can be used if the API you are trying to expose is providing a TLS certificate signed by a custom root certificate. The root certificate has to be available on the machine running the inway and the absolute path to the root certificate should be the value of this field.  

`example: "/path/to/custom-root-ca.crt"`
## public-support-contact
Contains an email address which NLX users can contact if they need your support when using this service. This email address is published in the [directory](https://directory.nlx.io).

`example: public-support-contact = "support@my-organization.nl"`
## tech-support-contact  
Contains an email address which we (the NLX organization) can contact if they have any questions about your API.
This email address will NOT be published in the directory.

`example: tech-support-contact = "tech@my-organization.nl"`

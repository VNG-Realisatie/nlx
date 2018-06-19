# outway
The NLX outway proxies outgoing requests to an inway. The outway process is operated by each organization that is making requests to other organizations.

## URI path scheme

The path scheme is as follows:

`/{OrganizationName}/{ServiceName}/{OriginalPath}`

* `OrganizationName` is the name of the organization that provides a service.
* `ServiceName` is the name of the service to which a request is sent.
* `OriginalPath` is the path inside the service, this is specific for each service and is documented in the services' API documentation.

# inway

The NLX inway proxies incomming requests to the service endpoints of an organization. The inway process is operated by each organization that provides services.

## URI path scheme

The path scheme is as follows:

`/{ServiceName}/{OriginalPath}`

* `ServiceName` is the name of the service which is requested.
* `OriginalPath` is sent as path to the service, this is specific for each service and must be specified by organizations in the service API documentation.

### Special endpoints

Special endpoints provide the centralized NLX service to fetch information from the _inway_.

#### `/-/health`

The health endpoint tells the directory whether the _inway_ is available and stable.

# Organization management: configuration

**DRAFT**

An organization starts by spawning the management components. These components provides a UI and API endpoints to inspect, modify and automate the configuration of all other NLX components within an organization.
Inways, outways and other components fetch their configuration when they are started, and automatically keep track of updates to the configuration.

## Services

A Service is an NLX abstraction of an API provided and used by organizations.

When creating a Service, several parameters must be configured.

- Name, must be unique within the organization
- The type of the service, e.g. REST, HTTP/JSON, gRPC, SOAP.
- Link to human readable documentation
- Link to system readable documentation ([OpenAPI](https://github.com/OAI/OpenAPI-Specification), [gRPC definitions](https://grpc.io/))
- Security model (public, private)

### Service security

A private Service needs requesting organizations to be whitelisted. The configuration management dashboard provides an interface to add and remove whitelisted organizations.

## Inways and Outways

Inways and outways self-register with the configuration management system, but must be acknowledged by an admin user. The admin user needs to enter the public key of the approved system. This is also the case for the insight-api.

### Inway / outway deployment stratiegies

The organization config management system must work well with auto-scaling and re-creation of inways/outways. We do this by handling separate instances of an xway as a single "inway" or "outway" in the config API. They are identified by the config api as a "inway" or "outway" when they identify themselves with the same self-init credentials.

For example:

An organization provides a stateless API. For redundancy and performance reasons this API is deployed to a scaling number of machines/pods. Each machine or pod has an inway which needs to be configured to provide the API as a NLX Service. By letting the inways share a drive that contains the keys, admins don't have to perform manual jobs every time a new inway spawns or is destroyed.

## Organization status

The centralized configuration management provides a dashboard where all connected gateways (and other components) are listed. For each gateway several values are visible:

- Overall status
- Number of instances
- A collapsable (or click for next page) list of instances. For each instance:
  - Status (alive, out-of-sync, shutdown) (State becomes out-of-sync after 10+ seconds of not watching(?). State becomes "shutdown" when the gateway properly said goodbye after stopping the listener.)
  - Version numbers of the instances
  - Version status (indicating importance of updating the gateway)

Network status should focus on whether the inway is watching and version checks.

Monitoring availability and performance is not a task of NLX. Existing open-source off-the-shelf solutions, such as Prometheus, perform these tasks well. This also goes for metrics, we may decide to add metrics endpoints for popular metric services (such as prometheus). We could also decide to add existing services to the common NLX organization helm chart / manifests.

## Security

### Authentication

Basic authentication and authorization is provided by the management system. A simple user and role system allows organizations to manage their employees' access ([RBAC](https://en.wikipedia.org/wiki/Role-based_access_control)) to the management system.

Open for discussion/brainstorm: authn adapters which run on-premise at the organization and connect to AD or other SSO systems to provide access tokens to users.

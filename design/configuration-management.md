
# Centralized configuration management

A centralized configuration management is provided by NLX. This centralized component provides UI's and API endpoints to modify the configuration of NLX components in an organization.
NLX components fetch their configuration when they are started, and automatically keep track of updates to the configuration.

## Services

A Service is an NLX abstraction of an API provided and used by organizations.

When creating a Service, several parameters can be configured.

 - Name, must be unique within the organization
 - The type of the service, e.g. REST, HTTP/JSON, gRPC, SOAP.
 - Link to human readable documentation
 - Link to system readable documentation ([OpenAPI](https://github.com/OAI/OpenAPI-Specification), [gRPC definitions](https://grpc.io/))
 - Security model (public, private)

### Service security

A private Service needs requesting organizations to whitelisted. The central configuration management provides an interface to add and remove whitelisted organizations.

## Configuration groups

To ease the management of multiple gateways (inways, outways) performing the same task InwaysGroups and OutwayGroups can be created. These groups 

An inway or outway can only be in one group. Need a machine to be in multiple groups? Run multiple Inways/Outways.

Example:

An organization provides a stateless API. For redundancy and performance reasons this API is deployed to 3 seperate machines and listens on `localhost:8080`. Each machine has an inway which needs to be configured to provide the API as a NLX Service. The organization creates an InwayGroup which has the service configured with an endpoint on `localhost:8080`. An inway is started on each machine and added to the InwayGroup. Each Inway then automatically registers to the directory and starts providing the Service backed by the `localhost:8080` endpoint to handle the request.

## Gateways

When a Gateway (Inway or Outway) is started, it generates a unique identifier for itself and registers to the directory. An employee can assign the Inway/Outway to an InwayGroup or OutwayGroup which triggers automatic configuration of the Inway/Outway.

To allow automatic deployment and autoscaling of Inways and Outways, they must be able to self-register to a group. The configuration flags/env-vars of `nlx-inway` and `nlx-outway` provide options to automatically self-register. e.g.: `nlx-inway --auto-register-group=my-first-group --auto-register-external-address=inway-001.environment.nlx.organization.com:443`

Note on Kubernetes and loadbalanced inways: when multiple `nlx-inway` processes are proxied by a single loadbalancer, which is the case in common kubernetes setups, they should register and behave as a single inway. This means the inway unique identifier must be manually forced to be equal for all `nlx-inway`'s behind the same loadbalancer. The inways must have the same configurations.

An Inway must have a publicly resolveable domain name pointing to it. The CN of the TLS certificate used by the Inway must match the public domain name.

## Network status

The centralized configuration management provides a dashboard where all connected gateways are listed. For each gateway several values are visible:

 - Healthy {yes, no, unreachable} (health-check reachable by healthchecker, health-check outcome can still be unhealthy)
 - Version number
 - Version status (indicating importance of updating)
 - Configuration in sync (when out of sync for too long, gateway should be taken down and inways become unhealthy)

## Security

### Authentication

Basic authentication and authorization is provided by the centralized management system. A simple user and role system allows organizations to manage their employees' access ([RBAC](https://en.wikipedia.org/wiki/Role-based_access_control)) to the management system.

Open for discussion/brainstorm: authn adapters which run on-premise at the organization and connect to AD or other SSO systems to provide access tokens to users.

### Encryption of secret configuration values

Most information stored in the directory and configuration manager is not secret, but some values must not be publicly available. This design does not contain a concrete proposal for encryption of secret values. Instead, we should first identify the needs organizations, and the kind of configuration values that should be encrypted. After that we should identify possible existing solutions (vault?) or design a NLX-specific system.

Investigate true end-to-end encryption/signing of values using [OpenPGP.js](https://github.com/openpgpjs/openpgpjs)

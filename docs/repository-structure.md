
# Repository structure

The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows cross-component changes to be combined in a single feature branch and pull-request.

## Components

- `inway`: a gateway that provides service endpoints to the NLX network
- `outway`: a gateway that forwards requests to a service endpoint on the NLX network
- `directory`: the central service where inways can register services and outways can request the list of available services.

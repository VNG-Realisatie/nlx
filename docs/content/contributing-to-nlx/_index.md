---
title: "Contributing to NLX"
description: ""
---

## Getting started
The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows changes that affect multiple components to be combined in a single feature branch and pull-request.

The root of the repository contains a `docker-compose.yml` which is configured to start a complete development environment.

Simply run the following command in the root of the repository to build and start all components:

```bash
docker-compose up
```

Where applicable, [`modd`](https://github.com/cortesi/modd) is used to rebuild and restart a component when changes in its source files are detected.

## Components
The repository currently holds the following components:

- Docs: The documentation you are currently reading.
- Outway: A gateway that forwards requests to a service endpoint on the NLX network.
- Inway: A gateway that provides service endpoints to the NLX network.
- Directory: the central service where inways can register services and outways can request the list of available services.

Each component generally contains a number of files:

- `Dockerfile`: The Dockerfile is used to build and release the component.
- `Makefile`: Common build instructions are kept in a makefile. The Makefile is usually ran from the build stage in the Dockerfile, but should be usable directly from the developer's machine as well.
- `modd.conf`: The modd.conf file contains instructions for [`modd`](https://github.com/cortesi/modd), which watches for changes in the source code and builds/runs the components.



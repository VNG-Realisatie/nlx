---
title: "Components"
description: ""
weight: 20
---

Each component generally contains a number of files:

- `Dockerfile`: The Dockerfile is used to build and release the component.
- `Makefile`: Common build instructions are kept in a makefile. The Makefile is usually ran from the build stage in the Dockerfile, but should be usable directly from the developer's machine as well.
- `modd.conf`: The modd.conf file contains instructions for [`modd`](https://github.com/cortesi/modd), which watches for changes in the source code and builds/runs the components.

The repository currently holds the following Go components:

- [`docs`](./docs): The documentation you're currently reading.
- [`inway`](./inway): A gateway that provides service endpoints to the NLX network
- [`outway`](./outway): A gateway that forwards requests to a service endpoint on the NLX network
<!-- - [`directory`](./directory): the central service where inways can register services and outways can request the list of available services. -->


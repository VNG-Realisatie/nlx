---
title: "Components"
description: ""
weight: 20
menu:
  docs:
    parent: "working-on-nlx-core"
---

Each component generally contains a number of files:

- `Dockerfile`: The Dockerfile is used to build and release the component.
- `Makefile`: Common build instructions are kept in a makefile. The Makefile is usually ran from the build stage in the Dockerfile, but should be usable directly from the developer's machine as well.
- `modd.conf`: The modd.conf file contains instructions for [`modd`](https://github.com/cortesi/modd), which watches for changes in the source code and builds/runs the components.

The repository currently holds the following Go components:

- Docs: The documentation you're currently reading.
- Inway: A gateway that provides service endpoints to the NLX network.
- Outway: A gateway that forwards requests to a service endpoint on the NLX network.
- Directoyr: the central service where inways can register services and outways can request the list of available services.

## Docs
These documentation pages are built using [hugo](https://gohugo.io).
To run the docs locally, cd into the `/docs` directory and run `hugo server`.

[More information about hugo.](https://gohugo.io/documentation/)

## Inway
Placeholder for `inway` documentation.

## Outway
Placeholder for `outway` documentation.
Test the outway by sending an HTTP GET request to: `http://localhost:12018/DemoProviderOrganization/PostmanEcho/get?foo=1`

# Directory
Placeholder for `directory` documentation.
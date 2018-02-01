---
title: "Introduction"
description: ""
weight: 10
---

The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows changes that affect multiple components to be combined in a single feature branch and pull-request.

The root of the repository contains a `docker-compose.yml` which is configured to start a complete development environment. Simply run `docker-compose up` in the root of the repository to build and start all components. Where applicable, [`modd`](https://github.com/cortesi/modd) is used to rebuild and restart a component when changes in its source files are detected.

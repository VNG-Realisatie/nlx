---
title: "Jenkins"
description: ""
weight: 120
menu:
  docs:
    parent: "contributing-to-nlx"
---

Jenkins performs some automated tasks.

- Every push to the nlx repository on github triggers a build on the latest commit of that push. Jenkins builds and verifies the commit and sets the commit status to github. This status is also shown when creating a Pull Request.
- Every push to the master branch is build and released using the short git commit hash as docker tag. When the release was succesfull, it is deployed to the test environment.
- When a tag is pushed, it is built and deployed to the test and acc environments.

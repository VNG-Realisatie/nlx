---
id: introduction
title: Introduction
---

## Deploying to a production like environment

This guide is made to help you take the next step from trying NLX to deploying NLX in a production like environment.

We do not (yet) favour a certain deployment strategy or technology, because there are virtually no limitations to 
the amount of variations of setups at the (potential) users of the NLX network.

We do not know how many requests you will need to manage, how much data will be transferred, how many organisations will act on behalf of you. 
Therefore, there is not a one size fits all deployment strategy for NLX.

Below are the most common strategies to deploy NLX to a production like environment:

### Deployment strategies

1. **Helm charts**  
   We use these ourselves to deploy the demo, pre-prod and prod environment and are confident that it will cover most demands of most users.  
   The [charts are located in the NLX repository](https://gitlab.com/commonground/nlx/nlx/-/tree/master/helm).
1. **Docker containers**  
   When you deploy every component in a separate container, we feel you we be able to size these according to your needs.
   The [Docker images are available from Docker Hub](https://hub.docker.com/u/nlxio).
1. **Native packages**
   We have created [native packages for Debian and Red Hat distro's](https://gitlab.com/commonground/nlx/packaging).

Obviously, we are willing to help you in every way we can when you want to start using NLX in production. 
See the [contact page](../support/contact.md) to find out how to reach us.

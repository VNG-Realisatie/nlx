---
title: "Implementation of NLX in your own infrastructure"
description: ""
weight: 20
menu:
  docs:
    parent: "further-reading"
---

## Implementation of NLX in your own infrastructure

Implementing NLX is a crucial part of building your Common Ground infrastructure. The current NLX infrastructure is based upon [Openstack](https://www.openstack.org/),  [Kubernetes (K8s)](https://kubernetes.io/) and [Docker](https://www.docker.com/). This current infrastructure choice of the development team however isn't mandatory and because the NLX software is developed open source it's also possible to deploy it on your own infrastructure of choice.  
The most common implementations we expect to see regarding NLX are the following five options:

1. **Dedicated machine:** Install the codebase as executable on a dedicated (virtual) machine.  
2. **Hosted containers:** Install the predefined container in your own onsite or offsite hosting environments.
3. **SaaS Cloud:** Use NLX as part of a SaaS proposition in the cloud environment of a solution provider.
4. **K8s Cloud:** Install in a kubernetes native cloud environment (just like the [nlx.io demo](https://demo.nlx.io/)).
*An alternative approach might be to:*  
5. **Embedded CI/CD:** Embed the NLX codebase as part of your own application codebase and include NLX in the CI/CD of the project you are developing.  

Bear in mind that you only need to run an outway to support your clients or an inway to expose your api's. Running NLX as stand alone system won't be a common situation and thus for most cases the infrastructure on which the client or api is running will be the default option to also run the NLX inway or outway.

### Considerations and remarks for the different types

**1. Dedicated machine:**

- [x] this type will be most demanding, besides setting up your own infrastructure you will need to set up the operating system, required libraries and the NLX software itself.
- [x] deploying NLX on a custom prepared OS currently isn't part of the devops activities of the NLX team. This type of installation will be possible but has yet to be done.
- [x] before choosing this option consult with the NLX devops team via support@nlx.io to evaluate the work ahead and possible pitfalls.
- [x] you can choose your own OS distribution.
- [x] you must apply your own OS hardening configuration.

**2. Hosted containers:**

- [x] this type of implementation doesn't require the use of kubernetes but will require knowledge and experience of docker.
- [x] setting up the docker containers for production will require you to create your own deployment scripts to setup each environment and manage the application life cycle.
- [x] using the available containers on [dockerhub](https://hub.docker.com/search?q=nlx&type=image) might speed up implementation.
- [x] use your own Host OS and docker installation.
- [x] Deploy onsite or offsite.

**3. SaaS Cloud:**

- [x] this type will outsource the entire technical NLX setup to the contracted SaaS provider and use NLX as add-on service.
- [x] the SaaS provider will need to invest time to adopt NLX and build a multi-tennant setup for multiple clients.
- [x] this will take time and effort of the SaaS provider and currently isn't a commodity. Your implementation will be the launching customer.

**4. K8s Cloud:**

- [x] this is the preferred option of the Common Ground and the way we have set up the [NLX.io environments]({{< ref "/environments.md" >}}) including our demo and insight applications.
- [x] his will provide a modern, hosted, scalable and resilient environment with the option to outsource the main infrastructure work to the specialist cloud provider.
- [x] Relatively complex to set up and maintain at first and it's advised to include developers in the team and use a devops approach to manage the deployments.

**5. Embedded CI/CD:**

- [x] because the source of NLX is available online it's possible to include the source of our components in your own project.
- [x] this will allow you to have full control of the CI/CD pipeline actions.
- [x] it's possible to write and execute additional unit and integration tests.
- [x] NLX can be deployed as part of your own application stack.  

### Pro's and Con's of the different implementation types

Different implementation types have different advantages and disadvantages. The most important ones are listed in the following table:

|Implementation type | Pro's | Con's  |
|:-------------------|:------|:-------|
|1. Dedicated machine| 1. choose and harden your own OS. | 1. no advantages of pre configured containers or infrastructure. |
|2. Hosted containers| 1. reuse the preconfigured containers on your own host OS and docker installation. | 1. the need to trust the preconfigured container setup or generate your own containers. |
|3. SaaS Cloud| 1. outsource the entire technical NLX setup to the SaaS provider; 2. use NLX as add-on. | 1. currently isn't a commodity. |
|4. K8s Cloud| 1. hosted, scalable and resilient environment;  2. outsource the main infrastructure work to the specialist K8s cloud provider | 1. relatively complex to set up and maintain at first. |
|5. Embedded CI/CD| 1. full control of the CI/CD; 2. ability to incorporate NLX in your own application. | 1. no advantages of pre configured containers or infrastructure. |

### Advice

We advise organizations to consider reusing the available containers and reuse the infrastructure setup of NLX based on [Openstack](https://www.openstack.org/) and [Kubernetes](https://kubernetes.io/) and [Docker](https://www.docker.com/). Especially when developing new clients and/or api's.

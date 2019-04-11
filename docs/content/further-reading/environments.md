---
title: "Environments"
description: ""
weight: 20
menu:
  docs:
    parent: "further-reading"
---


## Environments

NLX is build on a modern infrastructure based on OpenStack and Kubernetes, with a development street that automates most of the build, test and deployment tasks. The project and code repository are available on [Gitlab](https://gitlab.com/commonground/nlx). NLX developers have their own CI/CD flow with several private and public environments. The following Environments are publicly available for different audiences:

* `Acceptance (Acc)`; for testing new features and formal acceptance - *available for team members and stakeholders only.*
* `Demo`; for testing the general setup and functionality - available for all interested parties.
* `Preproduction(Preprod)`; for testing the interaction of NLX with developers own software - available for all developers after acquiring a certificate.
* `Production`; for production functionality only - available to all users and developers after testing the interaction on preproduction.

To make use of the NLX preproduction or production environment you will need **verified and signed certificates**. Read the instructions on the page ["further-reading/production"](../production) for obtaining a certificate.

Besides the public environments there are private development and test environments for the NLX team to develop new features:
* `Development`; private environments for development - available for team members only.
* `Test`; environment for testing new features - available for team members and stakeholders.

All new features are developed in separate [Gitlab branches](https://gitlab.com/commonground/nlx/branches), after the initial tests the branches are merged into the master branch and deployed to the test and acceptance environment using the [CI/CD pipelines of Gitlab](https://gitlab.com/commonground/nlx/pipelines). The pipeline for NLX requires manual code review and acceptance followed bij automated build, unit test , integration tests, packaging and deployment.

After testing by the stakeholders (in the Acc environment) the version of the master is tagged and this identical version is deployed to preproduction, production and demo.
All components like the directory, certportal, docs, demo and insight are available on all environments. The components are accessible using the following url construction:
* *< componentname.environment.product.topleveldomain >.*

Leaving the environment empty will redirect to production. For example:

|URL   | Environment | Component  |
|:-----|:----------:|:-------------|
| [Insight.demo.nlx.io](https://insight.demo.nlx.io/) |`Demo` |Insight application for viewing log entries|
|[Docs.preprod.nlx.io](https://docs.preprod.nlx.io)|`Preproduction`| Technical documentation of NLX|
|[Certportal.acc.nlx.io](https://certportal.acc.nlx.io/) |`Acceptance`| NLX certificate portal for this environment|
|[Directory.acc.nlx.io](https://directory.acc.nlx.io/) |`Acceptance`|Directory with all available api’s for this environment  |
|[Demo.nlx.io](https://demo.nlx.io/) |`Production`| Demo application (applying for a parking permit in Haarlem)|
|[NLX.io](https://nlx.io/) |`Production`| NLX home page |

*Note that: All non production CA's (Certificate Authority's) are un secured and for development and demonstration purpose only. Preproduction & Production require a restricted CA configuration and related certificates.*

An overview of the environments and the deployment flow is illustrated in the following archimate model:

![NLX Environments overview:](environments-ebb71.png "NLX Environments overview")

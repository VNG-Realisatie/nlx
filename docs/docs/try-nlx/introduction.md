---
id: introduction
title: Introduction
---

> Note that this tutorial is not suited for production environments. 
> Its only purpose is to enable you to setup a local test environment.

This guide is here to help you get started with NLX using NLX Management.
NLX Management provides you with a web interface for working with NLX. 
You will learn how to set up a test environment, provide and consume APIs on the NLX network.

The target audience is **system operators**.

![Screenshot of the NLX Management web interface](/img/nlx-management-web-interface-screenshot.png)

The following features are available:

* Manage your inways
* Provide services on the NLX network
* Manage access to your services
* Request access to services available on the NLX network


## Components

The following components are part of NLX Management:

* Management API
* Management UI
* PostgreSQL database
* An [OpenID Connect](https://openid.net/connect/) (OIDC) provider
* An Inway

![NLX Management Overview](/img/docs-nlx-management-overview.svg)

### Management API

The Management API is at the heart of NLX Management. It is used by the Management UI to manage 
your NLX setup and your Inways use the Management API to retrieve their configuration.


### Management UI

The Management UI is a web interface you use to manage your NLX setup. It is designed to be easy 
to understand and operate.


### PostgreSQL Database

The configuration of your Inways, the services you provide and access requests to your services 
are stored in an PostgreSQL database. PostgreSQL is used as a database because it is safe and reliable.


### OIDC provider

Users need to login before they can use NLX management. NLX management does **not** come with its 
own identity and access management tool but supports OIDC. If you are using OIDC in your organization 
you can easily configure NLX Management to use your OIDC provider.

If your organization does not support OIDC you can use an identity service 
like [Dex](https://github.com/dexidp/dex) to set it up.


## In sum

You've learned about all the components used by NLX. 
Next up, let's [setup our local environment](./setup-your-environment).

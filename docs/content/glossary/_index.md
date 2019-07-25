---
title: "Glossary"
description: ""
weight: 200
---

# A

## API

API stands for application programming interface, which is a software intermediary that allows two pieces of software to talk to each other.

[wiki](https://en.wikipedia.org/wiki/Application_programming_interface)

## Application

_Application_ is one of the main concepts of the NLX network, An application is only available within a organization and usually provide users an interface to query and mutate data. Applications call API's of [services](#service) in the NLX network.

# C

## CA

In cryptography, a certificate authority or certification authority (CA) is an entity that issues digital certificates. A digital certificate certifies the ownership of a public key by the named subject of the certificate. This allows others (relying parties) to rely upon signatures or on assertions made about the private key that corresponds to the certified public key. A CA acts as a trusted third party—trusted both by the subject (owner) of the certificate and by the party relying upon the certificate. The format of these certificates is specified by the X.509 standard.

[wiki](https://en.wikipedia.org/wiki/Certificate_authority)

# D

## Directory

The direcory lists all the services available on the NLX network.

## Docker

Docker is a tool designed to make it easier to create, deploy, and run applications by using containers. Containers allow a developer to package up an application with all of the parts it needs, such as libraries and other dependencies, and ship it all out as one package. By doing so, thanks to the container, the developer can rest assured that the application will run on any other Linux machine regardless of any customized settings that machine might have that could differ from the machine used for writing and testing the code.

[Official website](https://www.docker.com/)

[Video  : Learn docker in 12 minutes](https://www.youtube.com/watch?v=YFl2mCHdv24)

# G

## Golang

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. Go is the main programming language of NLX

[Official website](https://golang.org/)


# H

## Helm

Helm helps you manage Kubernetes applications — Helm Charts helps you define, install, and upgrade even the most complex Kubernetes application.
Charts are easy to create, version, share, and publish — so start using Helm and stop the copy-and-paste madness.

[Official website](https://helm.sh/)

## HTTP

The Hypertext Transfer Protocol (HTTP) is an application protocol for distributed, collaborative, hypermedia information systems. HTTP is the foundation of data communication for the World Wide Web, where hypertext documents include hyperlinks to other resources that the user can easily access, for example by a mouse click or by tapping the screen. HTTP was developed to facilitate hypertext and the World Wide Web.

[wiki](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)

## HTTPS

Hypertext Transfer Protocol Secure (HTTPS) is an extension of the Hypertext Transfer Protocol (HTTP) for secure communication over a computer network, and is widely used on the Internet. In HTTPS, the communication protocol is encrypted using Transport Layer Security (TLS), or formerly, its predecessor, Secure Sockets Layer (SSL). The protocol is therefore also often referred to as HTTP over TLS

[wiki](https://en.wikipedia.org/wiki/HTTPS)


# K

## Kubernetes

Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications.

[Official website](https://kubernetes.io/)

# M

## Mutual TLS

Mutual TLS is a an optional feature of TLS and it enables the server to authenicate the identity of the client.

# I

## Inway

The inway is used by an organization to provide [services](#service) to the NLX network.

[How to start an inway](../providing-services-on-nlx/start-an-inway/)

# O

## OpenAPI

The OpenAPI Initiative (OAI) was created by a consortium of forward-looking industry experts who recognize the immense value of standardizing on how REST APIs are described. As an open governance structure under the Linux Foundation, the OAI is focused on creating, evolving and promoting a vendor neutral description format.

[Offical website](https://www.openapis.org/)
[OpenAPI repository](https://github.com/OAI/OpenAPI-Specification)

## Outway

The outway is used by an organization to query [services](#service) available on the NLX network.

[How to start an outway](../using-services-on-nlx/start-an-outway/)

# R

## REST

Representational State Transfer (REST) is an architectural style that defines a set of constraints to be used for creating web services.

[Wiki](https://en.wikipedia.org/wiki/Representational_state_transfer)

# S

## Service

A service is one of the main concepts within the NLX network. Services provide an API to [applications](#application). All requests made to the service will go through the [inway](#inway) of the service and are logged in the [transaction log](#transaction-log-txlog).


# T

## TLS

TLS stand for transport layer security and is a crypthographic protocol designed to provide secure communication over a computer network.

[Wiki](https://en.wikipedia.org/wiki/Transport_Layer_Security)

## Transaction log (txlog)

Requests made by an [application](#application) to a [service](#service) are logged in the transaction log. The logging is done by the [outway](#outway) of the application and the [inway](#inway) of the [service](#service).

[Transaction logs](../transactionlogs/)

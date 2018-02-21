---
title: "Introduction"
description: ""
weight: 10
---

**NLX** is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape with many organisations.

The need for the creation of NLX arises from the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground), which aims to convert or replace the current monolithic information systems of Dutch municipalities (and wider government) with a state of the art, API-first software landscape, fulfilling societal demand for automation, transparency and privacy.

In this vision, everyone should be able to use API's in other organisations as easy as their own, while core data objects should only be manipulated by the one administratively responsible and used by all others. An additional advantage is that public data can easily be made available to everyone. To make this technically feasible in a controllable manner, NLX comes in play.

NLX provides a developer friendly way to use standardised resources between organisations. It provides an API gateway for querying services on the central network as well as an API gateway to offer services to the central network.

## High level overview
Within the NLX network there are two main concepts: **applications** and **services**. Applications are only available within organisations and usually provide users an interface to query and mutate data. Services provide an API to applications and are accessable within and across organisations.

NLX provides two different types of API gateways: the **inway** and **outway**. These gateways are usually deployed centrally within the organization although it is possible for one organisation to deploy multiple instances of inway and outway.

<div class="mermaid">
graph LR;
    App[Application] --> Out
    Out[Outway] --> Net{NLX network}
    Net --> In[Inway]
    In --> Serv[Service]
    Serv --> In
    In --> Net
    Net --> Out
    Out --> App
</div>

Here you can see a full request-response flow on NLX. An application performs a request on the outway within the same organisation. The outway routes the request to the inway of the organisation providing the service. The inway routes the request to the service. The service responds to the request and this response is routed through the NLX network, to the outway and arrives at the application.

## Security
The security mechanism is based on TLS Client Authentication (mutual TLS). The nodes in the NLX network trust a [Certificate Autority](https://en.wikipedia.org/wiki/Certificate_authority) to sign certificates of organisations. Outway and inway nodes identify themselves by showing the signed certificate. All traffic between the outway and inway is therefore encrypted using TLS.

The security of communication between the application and the outway and the service and the inway is work in progress.

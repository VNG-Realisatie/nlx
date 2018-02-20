---
title: "Introduction"
description: ""
weight: 10
---

**NLX** is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

The need for the creation of NLX arises from the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground), which aims to convert or replace the current monolithic information systems of Dutch municipalities (and wider government) with a state of the art, API-first software landscape, fulfilling societal demand for automation, transparency and privacy.

In this vision, everyone should be able to use API's in other organisations as easy as their own, while core data objects should only be manipulated by the one administratively responsible and used by all others. An additional advantage is that public data can easily be made available to everyone. To make this technically feasible in a controllable manner, NLX comes in play.

In essence, NLX provides a developer friendly way to use standardised resources on the scale of a whole country.

Next to a technological challenge the implementation of NLX will require changes in governance fitting on agile development and federated structures.

## Components
The NLX network exists of the following components:

- Outway: through this component applications within organisation can query services across the NLX network.
- Inway: enables an organisation to provide services on the NLX network.
- Directory: a central component within the NLX network that keeps track of all organisations and services.
- Certportal: a central component within the NLX network that distributes developer certificates.

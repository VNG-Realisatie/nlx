---
title: "NLX Product Vision"
description: ""
weight: 130
menu:
  docs:
    parent: "introduction"
---

**NLX** is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

## Index

* [Introduction](#introduction)
* [Core Requirements](#core-requirements)
* [NLX functionality in more detail](#nlx-functionality-in-more-detail)
* [Context](#context)


## Introduction

The need for the creation of NLX arises from the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground), which aims to convert or replace the current monolithic information systems of Dutch municipalities (and wider government) with a state of the art, API-first software landscape, fulfilling societal demand for automation, transparency and privacy.

In this vision, everyone should be able to use API's in other organisations as easy as their own, while core data objects should only be manipulated by the one administratively responsible and used by all others. An additional advantage is that public data can easily be made available to everyone. To make this technically feasible in a controllable manner, NLX comes in play.

In essence, NLX provides a developer friendly way to use standardised resources on the scale of a whole country.

Next to a technological challenge the implementation of NLX will require changes in governance fitting on agile development and federated structures.

## Core requirements

Functional:

- [ ] Facilitate federated authentication and authorisation
- [ ] Automate secure data connection setup
- [ ] Protocol API requests, for:
    - [ ] GDPR Purpose limitation principle
    - [ ] Publication of data use to data subject
    - [ ] Automated inter-organisational billing
    - [ ] Quality improvement
    - [ ] Auditing
    - [ ] Monitoring

Non-functional:

- [ ] Blazingly fast
- [ ] Developer-friendly
- [ ] Secure
- [ ] Open
- [ ] Federated
- [ ] Scalable
- [ ] Reliable
- [ ] Resilient


## NLX functionality in more detail

Although the NLX system will be fairly complex and will require extensively detailed requirements, it is possible to grasp the core of NLX by describing the three core functional requirements. This is what NLX is all about.

*For more details, read the [NLX Functional Design](https://gitlab.com/commonground/nlx/blob/master/design/functional-design.md).*


### Facilitate federated authentication and authorisation
Organisations offering a service should be able to authorise other organisations to use the service. How that other organisation deals internally with identification, authentication and authorisation should be irrelevant. NLX should provide in a way to identify and authenticate other organisations, and perform al necessary actions to convert internal identity into an external one when a request leaves the own organisation. As a consequence, a resource owner will only know which organisation uses the resource, not which person or which process inside that organisation is using it. By enhancing the ways we audit, resource owners can still have trust in the way their resources are used.

### Automate secure data connection setup
When using a service from an external organisation, NLX should automatically set up a secure connection to that other organisation. This is meant to be the API equivalent of what connections that meet the [Digikoppeling](https://www.logius.nl/diensten/digikoppeling/) standard do for the current Dutch eBMS- and SOAP-based connections. Instead of system administrators in every organisation being responsible for building connections to every other organisation (like with Digikoppeling), NLX will create certificate based secure peer-to-peer connections on demand.

### Protocol API requests

NLX should protocol (formally log) all requests that flow through it. This means logging at both client and server side. It also means logging is done close to the resource. The logs serve several purposes:

* *GDPR Purpose limitation principle*

  By adding a mandatory *purpose limitation claim* to every API request that involves personal data, NLX logs data usage compliant with GDPR. These logs can be used for publication and auditing (next two bullets).

* *Publication of meta data to data subject*

  Meta data from the request logs, be it about usage or manipulation of data, can be made available for data subjects. Note that this is no current requirement, but it is expected that GDPR will develop in this direction and impacts design choices. The decision to publish this meta data is to be made, case by case, by the provider of the service. User interface might be a central portal providing access to logs detailing every time a person's personal data is used or modified by an organisation, including the reason (based on the purpose limitation claim).

* *Auditing*

  The NLX system is based on both trust and control. Trust, even when receiving the first API request from an organisation. The fact that an organisation has access to the NLX network and has an identity that can be authenticated is enough to authorise it - even though that authorisation might result in no access. This way, administration effort goes down drastically, most interactions are fully automated. Every organisation is assumed to only request services that are necessary, including lawful purpose limitation and privacy by design. Control, by auditing everything. The framework of audits allows for trust in organisations without an established relation. It also means more effort will go into auditing, which should be done in a standardised way.

* *Automated inter-organisational billing*

  Although not a popular concept, the Common Ground vision will require a new distribution of costs. Main reason is the fact that those organisations responsible for very popular data will have to maintain services and infrastructure for a much larger audience than in the current situation, which comes with higher costs. If implemented right, cost distribution can be done fairly simple. Every request should be logged with its calculated cost. Eventually a clearing house construction might help to distribute costs with minimal administrative overhead.

* *Quality improvement*

  Analysis of logs allows for quality improvement, spotting weak links that slow down operation, repeated errors, and so on.

* *Monitoring*

  By monitoring the NLX logs all kinds of alerting is possible. Security alerts triggering incident response, quality improvement PDCA cycles, triggering specific compliancy audits and so on. This can be enhanced with machine learning for behavioural analysis (of organisations, not data subjects).


## How NLX should be as a system

Looking at the required functionality alone, it might appear as if enough software exists to provide a solution of the shelf. However, when looking at how NLX should 'be' as a system, it becomes apparent there is nothing like it yet.


### Blazingly fast
When using resources that are spread out over several other organisations instead of SQL Queries to a local database, performance is of utmost importance. NLX will be optimised for speed.

### Developer-friendly
This means: state of the art, fun to use, well-known modern techniques. Short 'Time To First Successful Call' for every service. Excellent documentation. No ambiguity in use cases whatsoever. Reference implementations and examples for many code languages. Organisation should have to spend minimal effort in administration. The tool chain should be flexible enough to accommodate different organisational structures.

### Secure
It's obvious that a system providing federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape has to be very secure. Especially when it serves in governmental environments with sensitive, personal data. It should at least be fully compliant with all mandatory and recommended frames of reference. Security by design should be a priority from the very beginning.

### Open
Software functioning in the core of government should be as transparent as possible. Hence, Open Source. Also, the NLX system should not be restricted to governmental organisations alone. The architecture should be open and technically allow for other organisations to join when demand exists.

### Distributed
Instead of introducing some kind of star topology (e.g. an unwanted bottle neck) NLX should function fully peer-to-peer and decentralised. NLX should function on the edge of every organisation's network. Requests from organisation A to organisation B go through NLX components placed in both organisations, via an on demand secure peer-to-peer connection. Another reason for this is the fact that every organisation will offer API's - if not providing central resources, at least there will be services designed to transfer case ownership and other process related services. Distributed design, where possible following federation concepts, is essential for scalability.

### Scalable
When NLX fulfils it's intended role the system will, distributed, process trillions of requests per year. Extreme scalability is essential.

### Reliable
As NLX provides access to much of the data used in government processes, reliability is obviously a core requirement. As a whole, the services provided by NLX should be available 100%. Think redundancy, automatic rerouting, and so on. Here the solution makes use of proven concepts for distributed availability.

### Resilient
If NLX somehow enters a state of emergency, the system should be resilient. It should easily bounce back from disruptions, preferably prevent them. If NLX is completely destroyed, it should be possible to rebuild the whole NLX system in a very short time.


## Context

* GEMMA
* Current transition from SOAP to RESTful API's

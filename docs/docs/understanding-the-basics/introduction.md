---
id: introduction
title: Introduction
---

**NLX** is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API ecosystem with many organizations.

NLX is built as a core component of the [Common Ground vision](https://github.com/VNG-Realisatie/common-ground), which aims to convert or replace the current monolithic information systems of Dutch municipalities (and wider government) with a state of the art, API-first software landscape, fulfilling societal demand for automation, transparency and privacy.

Important business benefits of NLX are:

* Lower integration costs because of standardised API integration layer
* Better data quality because data is used at the data source
* Better AVG compliance because end users get insight into data usage
* Better logging and auditing of data usage
* More agile software systems because NLX facilitates component based software systems

To reduce integration costs and simplify building component based software systems, everyone should be (technically) able to use API's in other organizations as easy as in their own, while core data objects should only be manipulated by the one administratively responsible and used by all others. An additional advantage is that public data can easily be made available to everyone. NLX makes it possible to easily and securely connect to any API in the ecosystem.

**NOTE**: We are currently running NLX software in production and looking for organizations to participate.
[participating organizations](https://directory.nlx.io/) contact: support@nlx.io

## High level business overview

Starting with NLX implies starting with API’s. More specific it's starting with
generic API’s instead of one-off integration solutions. NLX is useful for both
organizations that want to consume API's and organizations that want to
provide API's to other consuming organizations.

NLX is developed mainly by VNG Realisatie, but anyone can contribute on Gitlab.
To start using NLX, it is not necessary to contribute to NLX development, since
NLX can be installed 'as is'. A good way to start is by reading the
documentation and having your system engineers or suppliers follow the
[Try NLX guide](../try-nlx/setup-your-environment.md) to learn the concepts
of NLX and set up the first test implementation. Some knowledge of IT, the
internet, API's and security is required. The NLX team is available through
Gitlab for support to developers.

Setting up the first Outway for testing is done in a matter of a few hours,
setting up an Inway requires having a basic API and a publicly resolvable URL.
Configuring the Inway is done in a few hours, setting up the API might take
more time depending on your infrastructure. Launching a service on a public
available network (the internet) requires more than just bringing the NLX and
API software to production. A project should address all functional and non
functional requirements and business context before launch.

Consuming API’s is easier than providing an API. A good place to start is to
consume an already available API. Such as for example the [BAG
API](https://zakelijk.kadaster.nl/-/bag-api) of the Dutch Kadaster or the [KVK
API](https://developers.kvk.nl/) of the Dutch Chamber of Commerce.

Providing API’s through NLX is quite straightforward when the API’s are already
available. Suitable API’s to start with as a provider are two examples of API’s
that Amsterdam has made available. Consider providing an API for monuments or
garbage bin’s by reusing the already developed API specifications: [Monumenten
API](https://api.data.amsterdam.nl/monumenten) or [Afval
API](https://api.data.amsterdam.nl/afval/).

## High level technical overview

Within the NLX integration landscape there are two main concepts:
**applications** and **services**. Applications are only available within
organizations and usually provide users a **client** interface to query and
mutate data. Services provide **API's** to applications and are accessable
within and across organizations.

NLX provides a developer friendly way to use standardised resources between
organizations. It provides a gateway for querying services on the ecosystem as
well as a gateway to offer services to the ecosystem. NLX provides support
for HTTP/1.x services like REST/JSON and SOAP/XML. HTTP/2 support (gRPC) has 
been added.

NLX provides two different types of gateways: the **Inway** and **Outway**.
Through an Inway an organization can provide services to the NLX ecosystem and
through an Outway an organization can query services on the NLX ecosystem. The
gateways are usually deployed centrally within the organization although it is
possible for one organization to deploy multiple instances of Inway and Outway
on different locations.

![figure 1](/img/introduction-fig-1.svg)

Here you can see a full request-response flow on NLX. An application performs a
request on the Outway within the same organization. The Outway routes the
request to the Inway of the organization providing the service. The Inway
routes the request to the service. The service responds to the request and this
response is routed through the NLX landscape, to the Outway and arrives at the
application.

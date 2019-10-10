---
title: "IRMA"
description: "I Reveal My Attributes"
menu:
  docs:
    parent: "usecases"
---

## IRMA: I Reveal My Attributes

IRMA is a complementairy technology to NLX and the Common Ground. https://privacybydesign.foundation/irma/

An organization that needs access to personal data attributes could use IRMA.
A person loads their personal data into the IRMA app on their phone using Digid **or any other official goverment
registration.** Once their attributes are loaded on their phone they can share their data securely 
with an organization. 

The person with an prepared IRMA app interacting with the goverment / organization can choose to reveal 
only the information needed and can proof they are the owner of the information requested. 
If a persion agrees in the IRMA app to share asked attributes then IRMA app will proceed
to infrom the asking IRMA organization server with the requested information.

The organization can validate the provided attributes to be correct.

A person can provide and prove:

 - BSN
 - licence plate
 - 18+ years old. (and no more information)
 - 65+ years old.
 - Are a *registered* docter
 - many more.

Organizations participating 

 - BRP
 - DUO
 - AGB
 - iDIN
 - BIG

[check all possible sources of attributes](https://privacybydesign.foundation/uitgifte/ttps://privacybydesign.foundation/irma/)

# DEMO

[NLX and IRMA in action](demo.nlx.io) **local browser and mobile needed to scan the QR code.**


## Summary

An API exposed on NLX (inway) can tell the  consuming service where the IRMA server of the organization
is running.
With IRMA and NLX there is an easy way to authenticate a user and prove that the data belongs 
to the person requesting information.
Securely shareing of personal data opens up a lot of possibilities:

  - [insigth](https://insight.demo.nlx.io/) access logs that belong to *you*


## Requirements

An organization needs 

  - run an IRMA server
  - run published NLX API

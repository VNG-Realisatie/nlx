# Functional design

## Index

* [Introduction](#introduction)
* [Logging in NLX](#logging-in-nlx)

## Introduction

_Functional design will be growing from sprint to sprint._

## Logging in NLX

### Requirements

NLX nodes will be required to keep a log of all the incoming and outgoing service calls. The logs serve a number of purposes as described in the [NLX Product Vision](./docs/content/functional/product-vision.md):
* GDPR Purpose limitation principle
* GDPR Right of access for data subjects
* Assisting non-repudiation controls
* Auditing and accountability
* Quality improvement
* Monitoring
* Inter-organisational billing

The NLX logs must comply with a number of EU rules and regulations concerning data protection:
* GDPR Purpose limitation principle
* GDPR Data minimisation principle
* Information security (ISO 27001:2013)


### General notes
- [ ] Both on client and server side there will be an NLX node. Both sides will log all requests, just the details available differ. Good to keep in mind, there are two sides, two logs.
  - [ ] Even more in use cases that go through multi-layers of NLX nodes.
- [ ] There are multiple reasons for those logs. Each of those reasons should have an independent implementation.
- [ ] Everything is always logged for all purposes - authorisation mechanisms (and encryption) make sure the logs are only visible to those with the right to see them.
  - [ ] This includes logs of requests made for e.g. criminal detection

### GDPR Purpose Limitation principle   

_Note: During Field Lab  (March 5th - 9th) this will be specified and build, if possible._
- [ ] Every request should include metadata detailing the Purpose Limitation
- [ ] All possible (verified, approved) Purpose Limitation claims should be available from a register
- [ ] What exactly is included in a Purpose Limitation Claim has to be specified in detail
  - [ ] Base every claim in law and regulation
  - [ ] Allow for claims that are based on consent
  - [ ] Allow for (certified) validation ?
  - [ ] Codify it all ?
- [ ] Processing of the data stored in the NLX should be limited to the purposes listed in the functional design of the NLX logging requirements.
- [ ] This begs for extensive exploration of use cases.

### GDPR Data minimisation principle   
The principle of data minimisation is essentially the idea that, subject to limited exceptions, an organisation should only process the personal data that it actually needs to process in order to achieve its processing purposes.
- [ ] NLX should only log personal data which is needed for the functional purposes of the NLX logs as described in these logging requirements;

### GDPR Right of access for data subjects
The Right of Access is a data subject right.This gives citizens the right to get access to their personal data and information about how these personal data are being processed.
- [ ] An NLX node has to provide, upon request, an overview of the categories of data that are being processed as well as a copy of the actual log data.

Exercising the right to access involves the data subject logging into a portal, and requesting access to all log records that belong to him or her. Which means:
- [ ] NLX nodes should offer an API to produce log records that belong to a data subject
  - [ ] The client needed to connect to this API is not part of NLX but relevant as demo
- [ ] Somehow, log records should be related to data subjects when relevant
  - [ ] Which means probably that this relation should be known by the API based on the information model
- [ ] Not all data that belongs to a data subject can be shown to the subject. Data concerning active criminal and civil investigations should be withheld during the course of the investigation. This raises the following question:
  - [ ] How does NLX know what data to withhold?
- [ ] Requests for access to the data by the data subject must  be logged as well

### Inter-organisational billing
The NLX logs can be used to implement inter-organisational billing based on the use of services offered by the NLX node. Whether this is used or not is depending on policy, but the design choices should make it easy to implement. This way, the overhead for billing can be kept minimal.
- [ ] Count the number of successful requests per API per organisation
- [ ] Offer an API to exchange this info including a way to mark records as processed
- [ ] Organisational suggestion: Most efficient would be to use a clearing house construction: One organisation that calculates all and deals with invoicing

### Assisting non-repudiation controls
Non-repudiation is a method of assuring that something that’s actually valid cannot be disowned or denied. From the point of NLX, non-repudiation applies to communication between NLX nodes and the transfer of data. Its aim is to ensure that organisations that exchange data are unable to deny the authenticity of their signatures on the these data transfer, or that they were the originator of a particular message or transfer. The NLX logs will have to include artefacts which may be used to dispute the claims of organisations that denies being the originator of an action or communication. These artefacts are:
- [ ] An identity
- [ ] The authentication of the identity
- [ ] Evidence connecting the identified party to a communication or action

### Quality improvement
_To be specified later_
- [ ] Find ways to optimise the whole landscape
- [ ] Find ways to optimise single API’s
- [ ] Find organisations that miss opportunities etc.

### Auditing and accountability
Essential for the NLX concept. The logs should provide all information needed for public accountability and performing audits. Through regular audits organisations will be asked to prove that the use of services for which they have been authorised has been compliant with laws, regulations and bilateral agreements. By performing these audits check organisations can be asked to prove that they have used their autorisations correctly.

### Monitoring
By monitoring we aim to trigger processes with events. Events can be advanced (correlating other events, or found with Machine Learning methods) or simple - but in the end it should trigger an action. For example, an audit, of a specific analysis for Quality Improvement.

### Information security (ISO 27001:2013)
Annex A of ISO 27001:2013 subsection A.12.4 describes a number of requirements concerning logging and monitoring:
- [ ] Event logging: Register information about access and actions of users, errors, events, etc. in NLX nodes;
- [ ] Protection of log information: The NLX logs must be protected to prevent unauthorised changes or removal;
- [ ] Administrator and operator logs: Privileges of administrators and operators of NLX are different from the normal user privileges, which means that they can perform more actions on NLX nodes. NLX should register information about all users, regardless of the privileges that they have on the systems;
- [ ] Clock synchronization: All NLX nodes should be configured with the same time and date; otherwise, if an incident occurs and we want to carry out a traceability test of what has happened in the different NLX nodes involved, it can be difficult if each one has a different configuration. Therefore, the ideal scenario would be that NLX nodes have a synchronized time.

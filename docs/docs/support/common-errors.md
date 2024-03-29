---
id: common-errors
title: Common errors
---

## Overview

![Figure 1](/img/common-errors-figure-1.svg)

## Identify NLX errors

You can identify NLX errors by looking at the HTTP status code.
When the status code is `540`, the error is related to NLX.

If the error has occured in the Inway, the error source will be `inway:`.
When the error occured in the Outway,  the error source will be `outway:`.

This enables you to determine if the error is coming from NLX or from the service that is being provided.

## Technical details

To debug common errors while setting up your NLX service.
Some basic knowlegde about the NLX components is needed.
NLX has three kinds of components.

 - Directory service which the NLX team hosts at https://directory.nlx.io
   keeps track of all the available Inways and does basic health checks.
 - NLX Outways to consume services on the NLX network and
 - NLX Inways provide services on the NLX network to
   be accessed by Outways.

Any participating NLX organization has Outway(s) and or Inway(s).

Inways will register themselves at the directory so
Outways can find them and connect to them.
Outways will ask the directory for available services on startup
and refreshes this list periodically.

The directory monitoring service will check if the registered Inways
are accessible from the public internet and the directory does a
periodic health check.
If the directory monitoring service cannot reach an Inway the service
will appear to be unhealthy.

## Common  errors

 - **C1** Something went wrong between the Client and the Outway
   Could be multiple problems, for example the requested URL is invalid.
   Check the error message for more information
 - **O1** Outway fails to connect to an Inway.
   No direct connection is possible from the Outway to the Inway so there
   are most likely firewall issues. Check **M1** to figure out if the Inway or the
   Outway is having networking problems.
 - **OAS1** Something went wrong between the Outway and the Authorization server
 - **A1** Inway fails to connect to a service. The organization providing
   the service has to bring the service back online or has some internal
   networking problems.
 - **IAS1** Something went wrong between the Inway and the Authorization server
 - **M1** Directory monitoring cannot reach the inway to perform a healthcheck.
   The registered Inway service will appear as unhealthy on [https://directory.nlx.io](https://directory.nlx.io).
   An Outway will still recieve the advertised address of the Inway but
   most likely most Outway's cannot connect to an unhealthy Inway.

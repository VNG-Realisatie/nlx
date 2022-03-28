---
id: transaction-log-headers
title: Transaction log headers
---

NLX logs requests on both the **outway** and **inway**. NLX adds a globally unique id of the request `X-NLX-Logrecord-ID` for every request, which makes a request traceable throughout the network. The log record ID is created from 128 bits cryptographically secure random data and then encoded into a hexadecimal string of 32 characters.

The application that performs the request can append some additional headers that will be saved in the transaction log as well:

* `X-NLX-Request-User-Id`, the id of the user performing the request
* `X-NLX-Requester-User`, data about the user performing the request
* `X-NLX-Requester-Claims`, claims the requester possesses
* `X-NLX-Request-Application-Id`, the id of the application performing the request
* `X-NLX-Request-Subject-Identifier`, an subject identifier for purpose registration (doelbinding)
* `X-NLX-Request-Process-Id`, a process id for purpose registration (doelbinding)
* `X-NLX-Request-Data-Elements`, a list of requested data elements
* `X-NLX-Request-Data-Subject`, a key-value list of data subjects related to this request. e.g. `bsn=12345678,kenteken=ab-12-fg`

The headers set by the application are optional.

All request headers are logged before the request leaves the outway. The fields `X-NLX-Request-User-Id`, `X-NLX-Request-Application-Id`, `X-NLX-Request-Subject-Identifier`, `X-NLX-Requester-Claims` and `X-NLX-Requester-User` are stripped off the request before it is forwarded to the inway.

The value of a `X-NLX-*` header is limited to 1024 characters.

---
title: "Transaction log headers"
description: ""
weight: 10
menu:
  docs:
    parent: "further-reading"
---

NLX logs requests on both the **outway** and **inway**. NLX adds a globally unique id of the request `X-NLX-Logrecord-ID` for every request, which makes a request traceable throughout the network. The application that performs the request can append some additional headers that will be saved in the transaction log as well:

* `X-NLX-Request-User-Id`, the id of the user performing the request
* `X-NLX-Request-Application-Id`, the id of the application performing the request
* `X-NLX-Request-Subject-Identifier`, an subject identifier for purpose registration (doelbinding)
* `X-NLX-Request-Process-Id`, a process id for purpose registration (doelbinding)
* `X-NLX-Request-Data-Elements`, a list of requested data elements
* `X-NLX-Request-Data-Subject`, a key-value list of data subjects related to this request. e.g. `bsn=12345678,kenteken=ab-12-fg`

The headers set by the application are optional.

All request headers are logged before the request leaves the outway. The fields `X-NLX-Requester-User-Id`, `X-NLX-Request-Application-Id`, `X-NLX-Request-Subject-Identifier`, `X-NLX-Requester-Claims` and `X-NLX-Request-User` are stripped off the request before it is forwarded to the inway.
When a request arrives at the inway, the inway logs the request together with the name of the requesting organization. Finally, the inway appends the header named `X-NLX-Request-Organization` before sending the request to the API.

The value of a `X-NLX-*` header is limited to 1024 characters.

Logging is done to stdout in JSON format. You can use for example [fluentd](https://www.fluentd.org/) to collect the logs from the containers and forward them to a store. This is an example configuration for fluentd when using Kubernetes:

```
<match fluent.**>
@type null
</match>

<source>
@type tail
path /var/log/containers/*.log
pos_file /var/log/fluentd-containers.log.pos
time_format %Y-%m-%dT%H:%M:%S.%NZ
tag kubernetes.*
format json
read_from_head true
</source>

<filter kubernetes.**>
@type kubernetes_metadata
</filter>

<match kubernetes.var.log.containers.**fluentd**.log>
@type null
</match>

<match kubernetes.var.log.containers.**kube-system**.log>
@type null
</match>
```

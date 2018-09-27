---
title: "Headers"
description: ""
weight: 10
menu:
  docs:
    parent: "transactionlogs"
---

NLX logs requests on both the **outway** and **inway**. When an application performs a request it can append some additional headers for central logging:

* `X-NLX-Request-User-Id`, the id of the user performing the request
* `X-NLX-Request-Application-Id`, the id of the application performing the request
* `X-NLX-Request-Subject-Identifier`, an subject identifier for purpose registration (doelbinding)
* `X-NLX-Request-Process-Id`, a process id for purpose registration (doelbinding)
* `X-NLX-Request-Data-Elements`, a list of requested data elements
* `X-NLX-Request-Data-Subject`, a key-value list of data subjects related to this request. e.g. `bsn=12345678,kenteken=ab-12-fg`

The outway appends a globally unique `X-NLX-Request-Id` to make a request traceable through the network. All the headers are logged before the request leaves the outway. Then the fields `X-NLX-Request-User-Id`, `X-NLX-Request-Application-Id`, and `X-NLX-Request-Subject-Identifier` are stripped of and the request is forwarded to the inway. When a request arrives at the inway, the inway logs the request and also the requesting organization.

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

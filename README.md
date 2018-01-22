NLX
===
[![Build Status](https://jenkins.nlx.io/job/nlx/badge/icon?style=plastic)](https://jenkins.nlx.io/job/nlx/) ![Repo Status](https://img.shields.io/badge/status-concept-lightgrey.svg?style=plastic)

**NLX** is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

## Goal

The current goal of this project is to deliver a Proof of Concept for the [NLX Product Vision](./product-vision.md).

## Build and run

Make sure you have [installed Go](https://golang.org/doc/install) and [configured a `GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH) with `${GOPATH}/bin` added to your `PATH`.

To build and run all [components](./docs/repository-structure.md), execute the following in a shell.

```bash
go get github.com/cortesi/modd/cmd/modd
cd $GOPATH/src/github.com/VNG-Realisatie
git clone git@github.com:VNG-Realisatie/nlx.git
cd nlx
modd
```

[`modd`](https://github.com/cortesi/modd) will watch for changes in the source code and builds/runs the components. For details see [`modd.conf`](./modd.conf).

Alternatively to `modd`, you could use the [`go` tool](https://golang.org/cmd/go/) to build components and run manually.

## Licence

Copyright © VNG Realisatie 2017  
[Licensed under the EUPL](LICENCE.md)

NLX
===
[![Build Status](https://jenkins.nlx.io/job/nlx-release-master/badge/icon?style=plastic)](https://jenkins.nlx.io/job/nlx/) ![Repo Status](https://img.shields.io/badge/status-concept-lightgrey.svg?style=plastic)

**NLX** is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

## Goal

The current goal of this project is to deliver a Proof of Concept for the [NLX Product Vision](./docs/content/introduction/product-vision.md).

## Questions and contributions

Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Documentation

Please find the latest documentation of NLX on https://docs.nlx.io/

## Build and run

Make sure you have installed the following tools:

- [docker](https://docs.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)
- [go](https://golang.org/doc/install)

Also you will need to have [configured a `GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH) with `${GOPATH}/bin` added to your `PATH`.

Open a new shell so the `GOPATH` you have configured earlier is correctly loaded in the environment variables. Then run the following:

```bash
mkdir -p $GOPATH/src/github.com/VNG-Realisatie
cd $GOPATH/src/github.com/VNG-Realisatie
git clone git@github.com:VNG-Realisatie/nlx.git
cd nlx
docker-compose up
```

## Licence

Copyright © VNG Realisatie 2017  
[Licensed under the EUPL](LICENCE.md)

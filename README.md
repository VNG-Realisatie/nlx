NLX
===
[![Build Status](https://jenkins.nlx.io/job/nlx-release-master/badge/icon?style=plastic)](https://jenkins.nlx.io/) ![Repo Status](https://img.shields.io/badge/status-concept-lightgrey.svg?style=plastic)

NLX is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components to the current **Proof of Concept** of the [NLX Product Vision](https://docs.nlx.io/introduction/product-vision/). Do **not** use this code in production.

## Questions and contributions
Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Documentation
Please find the latest documentation of NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service on top of NLX.

## Build and run NLX locally
If you would like to deploy an NLX network on your local machine or contribute to the NLX code please follow these steps.

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

## License
Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

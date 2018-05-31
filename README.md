<h1 align="center"><img alt="NLX" src="logo.png" width="200"></h1>

[![Build Status](https://jenkins.nlx.io/job/nlx-release-master/badge/icon?style=plastic)](https://jenkins.nlx.io/) ![Repo Status](https://img.shields.io/badge/status-proof%20of%20concept-lightgrey.svg?longCache=true&style=plastic)

NLX is an open source inter-organisational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components to the current **Proof of Concept** of the [NLX Product Vision](https://docs.nlx.io/introduction/product-vision/). Do **not** use this code in production.

## Developing for NLX
Please find the latest documentation for using NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service that uses or provides API access over NLX.

## Questions and contributions
Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Building and running an NLX network locally
If you want to develop locally, or run your own NLX network you need to get all of the components of the NLX network to run.

All of the components that make up the NLX platform are in this repository.
Some components are meant to run as centralized NLX services, while others should run on-premise at organizations that want to connect to the network.

### Requirements
Make sure you have installed the following tools:

- [docker](https://docs.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)
- [go](https://golang.org/doc/install)

Also you will need to have [configured a `GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH) with `${GOPATH}/bin` added to your `PATH`.
After you set the `GOPATH`, be sure to reopen your terminal/shell to be sure the environment variables have been set correctly.

### Cloning
```bash
mkdir -p $GOPATH/src/github.com/VNG-Realisatie
cd $GOPATH/src/github.com/VNG-Realisatie
git clone git@github.com:VNG-Realisatie/nlx.git
cd nlx
```

### Running
You can now start all the components with
```bash
docker-compose up
```

You can now see what services are up and on what port you can reach them using `docker ps`

### Developing

Where applicable, [`modd`](https://github.com/cortesi/modd) is used to rebuild and restart a component when changes in its source files are detected.
There is no need to build individual components.

### Troubleshooting
If you are running into issues after pulling changes you might need to rebuild your containers using `docker-compose build`

If you are running into other issues, please [Post an Issue on GitHub](https://github.com/VNG-Realisatie/nlx/issues/new).

## Deploying and releasing
Our [CI system Jenkins](https://jenkins.nlx.io/) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, Jenkins builds and deploys it to the test and staging environments.

## License
Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

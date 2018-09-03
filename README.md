<h1><img alt="NLX" src="logo.png" width="200"></h1>

[![pipeline status](https://gitlab.com/commonground/nlx/badges/master/pipeline.svg)](https://gitlab.com/commonground/nlx/commits/master)  [![coverage report](https://gitlab.com/commonground/nlx/badges/master/coverage.svg)](https://gitlab.com/commonground/nlx/commits/master)  [![Repo Status](https://img.shields.io/badge/status-proof%20of%20concept-lightgrey.svg?longCache=true)](https://docs.nlx.io/introduction/product-vision/)

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
mkdir -p $GOPATH/src/go.nlx.io
cd $GOPATH/src/go.nlx.io
git clone https://gitlab.com/commonground/nlx
cd nlx
```

If you wish to contribute, fork the project and set the push origin to your fork.

```bash
git remote set-url --push origin git@gitlab.com:<YOUR-GITLAB-USERNAME>/nlx.git
```

### Running

You can now start all the components with

```bash
docker-compose up
```

You can now see what services are up and on what port you can reach them using `docker ps`

#### Ports in docker-compose

The NLX components default to standard ports (tcp/80, tcp/443) for http-based traffic. These ports are mapped to unique ports in docker-compose.yml.

Web frontends (serving HTML):

- ` 8001`: directory-ui HTTP
- ` 8002`: docs HTTP
- ` 8003`: certportal HTTP

Database

- ` 5432`: postgres container for directory and logdb (If you already have a postgresql running on your host, this will create a conflict)

API's:

- `10443`: directory gRPC/HTTPS
- `10080`: directory non-TLS HTTP
- `20080`: outway request proxy
- `30443`: inway requests proxy
- `40080`: logdb-api
- `50080`: unsafe-ca

All these ports are TCP ports.

### Developing

Where applicable, [`modd`](https://github.com/cortesi/modd) is used to rebuild and restart a component when changes in its source files are detected.
There is no need to build individual components.

### Troubleshooting

If you are running into issues after pulling changes you might need to rebuild your containers using `docker-compose build`

If you are running into other issues, please [Post an Issue on GitLab](https://gitlab.com/commonground/nlx/issues).

## Deploying and releasing

The [CI system of GitLab](https://gitlab.com/commonground/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

<h1><img alt="NLX" src="logo.png" width="200"> System</h1>

NLX is an open source inter-organizational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components required to act out the [NLX Product Vision](https://docs.nlx.io/understanding-the-basics/product-vision/).

# WARNING

***This project been deprecated.***

New users should use [NLX based on the FSC Standard](https://gitlab.com/commonground/nlx/fsc-nlx).

## Developing on NLX

Please find the latest documentation for using NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service that uses or provides API access over NLX.

## Questions and contributions

Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Troubleshooting

If you are running into other issues, please [Post an Issue on GitLab](https://gitlab.com/commonground/nlx/nlx/issues).

## Building and running an NLX network locally

The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows changes that affect multiple components to be combined in a single feature branch and merge-request.

If you want to develop locally, or run your own NLX network, you will likely want to start all the components.

### Cloning

Clone NLX in your workspace.

Note for Go developers: We advise to not clone NLX inside the GOPATH. If you must, be sure to set the environment variable `GO111MODULE=on`.
Go doesn't need to be located in the GOPATH since it uses Go module support.

```bash
git clone https://gitlab.com/commonground/nlx/nlx
cd nlx
```

### Development setup

Make sure you have installed the following tools:

- [Docker Desktop / Docker engine](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://golang.org/doc/install)
- [NodeJS LTS](https://nodejs.org/en/download/)
- [Sass](https://sass-lang.com/install)
- [modd](https://github.com/cortesi/modd)

Install the npm dependencies by running:

```bash
(cd management-ui && npm install)
(cd docs/website && npm install)
```

Make sure the TLS key files have the correct permissions to run the NLX components
```bash
./pki/fix-permissions.sh
```

Update the `/etc/hosts` file on your system:

<details>
  <summary>Show hosts</summary>

```
# NLX
127.0.0.1     dex.shared.nlx.local
127.0.0.1     directory-api.shared.nlx.local

127.0.0.1     management-api.organization-a.nlx.local
127.0.0.1     management-proxy.organization-a.nlx.local
127.0.0.1     inway.organization-a.nlx.local
127.0.0.1     outway.organization-a.nlx.local
127.0.0.1     outway-2.organization-a.nlx.local
127.0.0.1     management.organization-a.nlx.local
127.0.0.1     txlog-api.organization-a.nlx.local

127.0.0.1     management-api.organization-b.nlx.local
127.0.0.1     management-proxy.organization-b.nlx.local
127.0.0.1     inway.organization-b.nlx.local
127.0.0.1     management.organization-b.nlx.local
127.0.0.1     auth.organization-b.nlx.local
127.0.0.1     txlog-api.organization-b.nlx.local

127.0.0.1     management-api.organization-c.nlx.local
127.0.0.1     inway.organization-c.nlx.local
127.0.0.1     outway.organization-c.nlx.local
127.0.0.1     management.organization-c.nlx.local
127.0.0.1     auth.organization-c.nlx.local
127.0.0.1     txlog-api.organization-c.nlx.local

::1           dex.shared.nlx.local
::1           directory-api.shared.nlx.local

::1           management-api.organization-a.nlx.local
::1           inway.organization-a.nlx.local
::1           management.organization-a.nlx.local
::1           txlog-api.organization-a.nlx.local

::1           management-api.organization-b.nlx.local
::1           inway.organization-b.nlx.local
::1           management.organization-b.nlx.local
::1           auth.organization-b.nlx.local

::1           management-api.organization-c.nlx.local
::1           inway.organization-c.nlx.local
::1           management.organization-c.nlx.local
::1           auth.organization-c.nlx.local
```
</details>
</br>

To start the services in development daemons with up-to-date databases, run: `./scripts/start-development.sh`. Make sure Docker is running.

During the starting routines of the services, you might see a few services erroring that are dependent on a service that has not yet been started.
This is expected behavior and will resolve itself within 5 seconds.

Services will reload automatically when the code changes and is saved.

Finally, run the Management UI using:
```bash
(cd management-ui && npm start)
(cd management-ui && npm run start:b)
```

This will start the management dashboard locally:

- [management-ui (A)](http://management.organization-a.nlx.local:3011) (HTTP: 3011)
- [management-ui (B)](http://management.organization-b.nlx.local:3021) (HTTP: 3021)

To log in locally, see credentials in `dex.dev.yaml`.

To test if the applications are running correctly, create a service called "Test" for "Organization A" and cURL the outway with `curl localhost:7917/Organization-A/Test` on the command line.

To start the directory UI or documentation website locally:

- `(cd directory-ui && npm start)` to run the [directory-ui](http://localhost:3001) (HTTP: 3001)
- `(cd docs/website && npm start)` to run the [docs](http://localhost:3002) (HTTP: 3002)

## Deploying and releasing

The [CI system of GitLab](https://gitlab.com/commonground/nlx/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

For branches prefixed with `review/` (f.e. `review/feature-name`), a review environment will automatically be created for a shared testing environment.

## Live environments

There are multiple environments for NLX

- `acceptance`: follows the master branch automatically
- `demo`, `pre-production` and `production`: updated after manually triggering a release

Overview of available links per environment:

- [links.demo.nlx.io](https://links.demo.nlx.io)
- [links.acc.nlx.io](https://links.acc.nlx.io)

## Regenerating protobuf and mock files

Te execute the commands in the following sections, you will need to install [Earthly](https://earthly.dev/get-earthly).
If you want to regenerate all files (protobuf and mocks), use the command:

```shell
earthly +all
```

### Compiling gRPC proto files

Protobuf files can be regenerated using:

```shell
earthly +proto
```

You might want to regenerate the mocks after recompiling the protobuf files.

### Mocks

Mocks can be regenerated using:

```shell
earthly +mocks
```

## Further documentation

* [Development setup using Minikube](technical-docs/development-setup-using-minikube.md)
* [Technical notes](technical-docs/notes.md)
* [Official documentation website](https://docs.nlx.io)

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

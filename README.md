<h1><img alt="NLX" src="logo.png" width="200"> System</h1>

NLX is an open source inter-organizational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components required to act out the [NLX Product Vision](https://docs.nlx.io/understanding-the-basics/product-vision/).

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

### Running the complete stack using modd

Make sure you have installed the following tools:

- [Docker Desktop / Docker engine](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [go](https://golang.org/doc/install)
- [modd](https://github.com/cortesi/modd)
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) for PostgreSQL (`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`)

Install the npm dependencies by running:

```bash
(cd directory-ui && npm install)
(cd management-ui && npm install)
(cd docs/website && npm install)
```

Start a PostgreSQL container through Docker Compose with:

```bash
docker-compose -f docker-compose.dev.yml up -d
```

Run the directory migrations with:

```bash
migrate -database "postgres://postgres:postgres@127.0.0.1:5432/nlx?sslmode=disable" -path directory-db/migrations up
```

Run the management API migrations with:

```bash
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
```

Create admin users for the Management API
```
go run ./management-api create-user --email admin@nlx.local --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api create-user --email admin@nlx.local --password development --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
```

Optionally you can setup the databases for the transaction logs:

```bash
migrate -database "postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_a?sslmode=disable" -path txlog-db/migrations up
migrate -database "postgres://postgres:postgres@127.0.0.1:5432/nlx_txlog_b?sslmode=disable" -path txlog-db/migrations up
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
127.0.0.1     directory-inspection-api.shared.nlx.local
127.0.0.1     directory-registration-api.shared.nlx.local

127.0.0.1     management-api.organization-a.nlx.local
127.0.0.1     inway.organization-a.nlx.local
127.0.0.1     management.organization-a.nlx.local

127.0.0.1     management-api.organization-b.nlx.local
127.0.0.1     inway.organization-b.nlx.local
127.0.0.1     management.organization-b.nlx.local

::1           dex.shared.nlx.local
::1           directory-inspection-api.shared.nlx.local
::1           directory-registration-api.shared.nlx.local

::1           management-api.organization-a.nlx.local
::1           inway.organization-a.nlx.local
::1           management.organization-a.nlx.local

::1           management-api.organization-b.nlx.local
::1           inway.organization-b.nlx.local
::1           management.organization-b.nlx.local
```
</details>
</br>

Run the services with:

```bash
modd # To run without transaction logs
TXLOG_A=1 modd # To run transaction logs enabled for organization A
TXLOG_B=1 modd # To run transaction logs enabled for organization B
TXLOG_A=1 TXLOG_B=1 modd # To run translaction logs for both organizations
```

This will start the following services:

- [S] directory-inspection-api (gRPC: 7901, HTTP: 7902)
- [S] directory-registration-api (gRPC: 7903)
- [S] directory-monitor
- [A] management-api (gRPC: 7911, HTTP: 7912)
- [A] inway (gRPC: 7913)
- [A] outway (HTTP: 7915)
- [B] management-api (gRPC: 7921, HTTP: 7922)
- [B] inway (gRPC: 7923)

And the following frontend applications:

- [directory-ui](http://localhost:3001) (HTTP: 3001)
- [docs](http://localhost:3002) (HTTP: 3002)

Services will reload automatically when the code changes and is saved.

Finally, run the Management UI using:
```bash
(cd management-ui && npm start)
(cd management-ui && npm run start:b)
```

This will start the management dashboard locally:

- [management-ui (A)](http://management.organization-a.nlx.local:3011) (HTTP: 3011)
- [management-ui (B)](http://management.organization-b.nlx.local:3021) (HTTP: 3021)

To log in locally, see credentials in `dex.dev.yaml`

## Deploying and releasing

The [CI system of GitLab](https://gitlab.com/commonground/nlx/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

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

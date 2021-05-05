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
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [modd](https://github.com/cortesi/modd)

Install the npm dependencies by running:

```bash
(cd directory-ui && npm install)
(cd management-ui && npm install)
(cd docs/website && npm install)
(cd insight-ui && npm install)
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
go run ./management-api/cmd/nlx-management-api/ migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api/cmd/nlx-management-api/ migrate up --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
```

Create admin users for the Management API
```
go run ./management-api/cmd/nlx-management-api/ create-user --email admin@nlx.local --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_a?sslmode=disable"
go run ./management-api/cmd/nlx-management-api/ create-user --email admin@nlx.local --role admin --postgres-dsn "postgres://postgres:postgres@127.0.0.1:5432/nlx_management_org_b?sslmode=disable"
```

Optionally you can setup the database for the transaction logs:

```bash
docker-compose -f docker-compose.dev.yml exec -u postgres postgres createdb nlx-txlog-a
migrate -database "postgres://postgres:postgres@127.0.0.1:5432/nlx-txlog-a?sslmode=disable" -path txlog-db/migrations up

docker-compose -f docker-compose.dev.yml exec -u postgres postgres createdb nlx-txlog-b
migrate -database "postgres://postgres:postgres@127.0.0.1:5432/nlx-txlog-b?sslmode=disable" -path txlog-db/migrations up
```

Make sure the TLS key files have the correct permissions to run the NLX components
```bash
./pki/fix-permissions.sh
```

Finally run the project with:

```bash
modd

# To run transaction logs enabled for organization A
TXLOG_A=1 modd

# To run transaction logs enabled for organization B
TXLOG_B=1 modd

# Or both
TXLOG_A=1 TXLOG_B=1 modd
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
- [insight-ui](http://docs.shared.nlx.local:3003) (HTTP: 3003)
- [management-ui (A)](http://management.organization-a.nlx.local:3011) (HTTP: 3011)
- [management-ui (B)](http://management.organization-b.nlx.local:3021) (HTTP: 3021)

Services will reload automatically when the code changes.

Update the `/etc/hosts` file:

```
# NLX
127.0.0.1     dex.shared.nlx.local
127.0.0.1     directory-inspection-api.shared.nlx.local
127.0.0.1     directory-registration-api.shared.nlx.local

127.0.0.1     etcd.organization-a.nlx.local
127.0.0.1     management-api.organization-a.nlx.local
127.0.0.1     inway.organization-a.nlx.local
127.0.0.1     management.organization-a.nlx.local

127.0.0.1     etcd.organization-b.nlx.local
127.0.0.1     management-api.organization-b.nlx.local
127.0.0.1     inway.organization-b.nlx.local
127.0.0.1     management.organization-b.nlx.local

::1           dex.shared.nlx.local
::1           directory-inspection-api.shared.nlx.local
::1           directory-registration-api.shared.nlx.local

::1           etcd.organization-a.nlx.local
::1           management-api.organization-a.nlx.local
::1           inway.organization-a.nlx.local
::1           management.organization-a.nlx.local

::1           etcd.organization-b.nlx.local
::1           management-api.organization-b.nlx.local
::1           inway.organization-b.nlx.local
::1           management.organization-b.nlx.local
```

To log in locally, see credentials in `dex.dev.yaml`


## Deploying and releasing

**NOTE:** Automated releases are currently not available.

The [CI system of GitLab](https://gitlab.com/commonground/nlx/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

## Live environments

There are multiple live environments for NLX

- `acceptance`: follows the master branch automatically
- `demo`, `pre-production` and `production`: updated after manually triggering a release

## Compiling gRPC proto files

You will need to install [Earthly](https://earthly.dev/get-earthly) to (re)compile the protobuf files for all services.
After installing Earthly you can run `earthly +proto` to compile everything.

You might want to regenerate the mocks after recompiling the protobuf files.
You can do so as follows for these directories:

```shell
(cd directory-inspection-api && sh regenerate-gomock-files.sh)
```

```shell
(cd directory-registration-api && sh regenerate-gomock-files.sh)
```

```shell
(cd management-api && make -B)
```


## Further documentation

* [Development setup using Minikube](technical-docs/development-setup-using-minikube.md)
* [Technical notes](technical-docs/notes.md)
* [Official documentation website](https://docs.nlx.io)

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

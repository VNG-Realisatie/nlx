<h1><img alt="NLX" src="logo.png" width="200"> System</h1>

NLX is an open source inter-organizational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components required to act out the [NLX Product Vision](https://docs.nlx.io/understanding-the-basics/product-vision/).

## Developing on NLX

Please find the latest documentation for using NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service that uses or provides API access over NLX.

## Questions and contributions

Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

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

### Running the complete stack using Minikube

Make sure you have installed the following tools:

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [helm](https://helm.sh/docs/intro/)

For autocompletion and local development tasks, it's also recommended to install the following:

- [go](https://golang.org/doc/install)

Setup minikube on your local development machine.

Read the [minikube README](https://github.com/kubernetes/minikube) for more information.

Configure the vm driver for minikube:

- for Linux: `minikube config set vm-driver kvm2`
- for Mac: `minikube config set vm-driver hyperkit`

For developers, it's advised to setup minikube with 4 cores, 8GB RAM and at least 100G storage.
e.g.: `minikube start --cpus 4 --memory 8192 --disk-size=100G`


Add the minikube hostnames to your machine's resolver so you can reach the services from your browser.

> see https://github.com/kubernetes/minikube/tree/master/deploy/addons/ingress-dns

```bash
minikube addons enable ingress-dns
```

On MacOS:

```bash
sudo mkdir -p /etc/resolver
sudo tee /etc/resolver/minikube <<EOF
nameserver $(minikube ip)
search_order 1
timeout 5
EOF
```

To let the docker commands make use of Minikube execute the following before proceeding or add it to your shell profile:

```bash
eval $(minikube docker-env)
```

Once minikube is running, install Traefik as ingress controller for web and rest-api requests.

```bash
helm repo add traefik https://containous.github.io/traefik-helm-chart
helm repo update

kubectl create namespace traefik
helm install traefik traefik/traefik --namespace traefik --values helm/traefik-values-minikube.yaml
```

Also install KubeDB, an operator that manages postgres instances. Follow the [kubedb.com instructions for installing using helm](https://kubedb.com/docs/0.12.0/setup/install/#using-helm) and click the 'Helm' tab.

Install cert-manager to issue certificates automatically.

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update

kubectl create namespace cert-manager
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.16.1/cert-manager.crds.yaml

helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v0.16.1
```

> Also see: https://cert-manager.io/docs/installation/kubernetes/#installing-with-helm


When Traefik and KubeDB are running, you can start all the NLX components by executing:

```bash
helm repo add stable https://charts.helm.sh/stable
helm dependency build ./helm/deploy/gemeente-stijns
helm dependency build ./helm/deploy/rvrd

helm upgrade --install shared ./helm/deploy/shared
helm upgrade --install rvrd ./helm/deploy/rvrd
helm upgrade --install gemeente-stijns ./helm/deploy/gemeente-stijns
```

You may now test the following sites:

- http://traefik.minikube:9000/                     Webinterface showing the status of the traefik ingress controller
- http://docs.shared.nlx.minikube/                  Documentation
- http://certportal.shared.nlx.minikube/            Portal to generate TLS certificates
- http://directory.shared.nlx.minikube/             Overview of all services in the network
- http://insight.shared.nlx.minikube/               Insight in logs concerning a specific person
- http://parkeren.gemeente-stijns.nlx.minikube/             Demo application for requesting a parking permit
- http://nlx-management.gemeente-stijns.nlx.minikube/       NLX management UI of example demo organization Gemeente Stijns used to manage NLX
- http://nlx-management.rvrd.nlx.minikube/          NLX management UI of example organization RvRD used to manage NLX

To test a full request through outway>inway, use the RvRD example service through the exampleorg outway: `curl http://outway.nlx-dev-gemeente-stijns.minikube/RvRD/basisregistratie/natuurlijke_personen`

If you want to connect over IP instead of using a hostname, the ingress controller cannot route the request properly. Therefore you must setup a port-forward directly to the application you want to expose. This is useful, for example, when testing IRMA using a phone on the same WiFi network as your host machine.

```bash
kubectl port-forward deployment/rvrd-irma-server 2222:session
socat tcp-listen:3333,fork tcp:127.0.0.1:2222
```

You can now let your phone connect to the IRMA api server of RvRD on `your.host.machine.ip:3333`

## A note on frontend debugging IE11

To locally debug and develop for IE11, please adjust `package.json` to the following:

1. Find `"browserslist":`
1. Copy the part under `production` to `development`
1. `$ rm -rf node_modules/.cache`
1. `$ npm i && npm start`
1. Open your local app in VM or Browserstack

## Troubleshooting

If you are running into other issues, please [Post an Issue on GitLab](https://gitlab.com/commonground/nlx/nlx/issues).

## Deploying and releasing

**NOTE:** Automated releases are currently not available.

The [CI system of GitLab](https://gitlab.com/commonground/nlx/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

## Live environments

There are multiple live environments for NLX

- `acceptance`: follows the master branch automatically
- `demo`, `pre-production` and `production`: updated after manually triggering a release

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

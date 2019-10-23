<h1><img alt="NLX" src="logo.png" width="200"> System</h1>

NLX is an open source inter-organizational system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API landscape.

This repository contains all of the components required to act out the [NLX Product Vision](https://docs.nlx.io/introduction/product-vision/).

## Developing on NLX

Please find the latest documentation for using NLX on [docs.nlx.io](https://docs.nlx.io). This is a good place to start if you would like to develop an application or service that uses or provides API access over NLX.

## Questions and contributions

Read more on how to ask questions, file bugs and contribute code and documentation in [`CONTRIBUTING.md`](CONTRIBUTING.md).

## Building and running an NLX network locally

The NLX project consists of multiple components that together make up the entire NLX platform. Some components run as centralized NLX services, others run on-premise at organizations. All components are maintained in a single repository. This means that a developer has all the tools and code to build and test the complete NLX platform in a single repository. It simplifies version and dependency management and allows changes that affect multiple components to be combined in a single feature branch and merge-request.

If you want to develop locally, or run your own NLX network, you will likely want to start all the components.

### Requirements

Make sure you have installed the following tools:

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [helm](https://docs.helm.sh/using_helm/)
- [skaffold](https://github.com/GoogleContainerTools/skaffold#installation)

For autocompletion and local development tasks, it's also recommended to install the following:

- [go](https://golang.org/doc/install)

This project uses the new go module feature so it is not required to setup a `GOPATH`.

### Cloning

Clone NLX in your workspace.

Note for Go developers: We advise to not clone NLX inside the GOPATH. If you must, be sure to set the environment variable `GO111MODULE=on`.
Go doesn't need to be located in the GOPATH since it uses Go module support.

```bash
git clone https://gitlab.com/commonground/nlx/nlx
cd nlx
```

### Running complete stack in kubernetes/minikube

Setup minikube on your local development machine.

Read the [minikube README](https://github.com/kubernetes/minikube) for more information.

Configure the vm driver for minikube:

- for Linux: `minikube config set vm-driver kvm2`
- for Mac: `minikube config set vm-driver hyperkit`

For developers, it's advised to setup minikube with 4 cores, 8GB RAM and at least 100G storage.
e.g.: `minikube start --cpus 4 --memory 8192 --disk-size=100G`

Once minikube is running, initialize helm by running `helm init` followed by `helm repo update`

Next, install Traefik as ingress controller for web and rest-api requests.

```bash
helm install stable/traefik --name traefik --namespace traefik --values helm/traefik-values-minikube.yaml
```

Also install KubeDB, an operator that manages postgres instances. Follow the [kubedb.com instructions for installing using helm](https://kubedb.com/docs/0.12.0/setup/install/#using-helm).

When Traefik and KubeDB are running, you can start all the NLX components by executing:

```bash
skaffold dev --profile minikube
```

Finally, add the minikube hostnames to your machine's `/etc/hosts` file so you can reach the services from your browser.

```bash
sh initialize-hostnames.sh
```

You may now test the following sites:

- https://traefik.minikube/                         Webinterface showing the status of the traefik ingress controller
- http://docs.nlx-dev-directory.minikube/           Documentation
- http://certportal.nlx-dev-directory.minikube/     Portal to generate TLS certificates
- http://directory.nlx-dev-directory.minikube/      Overview of all services in the network
- http://application.nlx-dev-haarlem.minikube/      Demo application for requesting a parking permit
- http://outway.nlx-dev-haarlem.minikube/           Outway of the Haarlem example organization
- http://txlog.nlx-dev-rdw.minikube/                Transaction logs of the RDW example organization
- http://txlog.nlx-dev-brp.minikube/                Transaction logs of the BRP example organization
- http://insight.nlx-dev-directory.minikube/        Insight in logs concerning a specific person

To test a full request through outway>inway, use the PostmanEcho service through the exampleorg outway: `curl http://outway.nlx-dev-haarlem.minikube/DemoProviderOrganization/PostmanEcho/get?foo1=bar1&foo2=bar2`

If you want to connect over IP instead of using a hostname, the ingress controller cannot route the request properly. Therefore you must setup a port-forward directly to the application you want to expose. This is useful, for example, when testing IRMA using a phone on the same WiFi network as your host machine.

```bash
kubectl --namespace nlx-dev-rdw port-forward deployment/irma-api-server 2222:8080
socat tcp-listen:3333,fork tcp:127.0.0.1:2222
```

You can now let your phone connect to the IRMA api server of RDW on `your.host.machine.ip:3333`

## Troubleshooting

If you are running into other issues, please [Post an Issue on GitLab](https://gitlab.com/commonground/nlx/nlx/issues).

## Deploying and releasing

**NOTE:** Automated releases are currently not available.

The [CI system of GitLab](https://gitlab.com/commonground/nlx/nlx/pipelines) builds every push to the master branch and creates a release to Docker, tagging it with the short git commit hash.
When a release is successful, it also gets deployed to the test environment.

When a git tag is pushed, GitLab builds and deploys it to the test and staging environments.

## Live environments

There are multiple live environments for NLX

- `acc`: follows the master branch automatically
- `demo`, `preprod` and `prod`: updated after manually triggering a release

## License

Copyright © VNG Realisatie 2017

[Licensed under the EUPL](LICENCE.md)

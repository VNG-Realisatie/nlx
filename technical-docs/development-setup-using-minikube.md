# Development setup using Minikube

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

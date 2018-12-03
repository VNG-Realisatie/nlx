# Helm charts for NLX

## Development environment

### Setup dependencies

Install the following dependencies:

- `traefik` for web and rest-api requests.
- `nginx-ingress` for grpc and mutual-tls connections. Latest version is currently(2018-09-06) broken, so needs `--version 0.17.1`
- `postgres` for directory-db and txlog-db.

```bash
helm install stable/traefik --name traefik --namespace traefik --values helm/traefik-values.yaml
helm install stable/nginx-ingress --version 0.17.1 --name nginx-ingress --namespace=nginx-ingress --values helm/nginx-ingress-values.yaml
helm install stable/postgresql --name postgresql --namespace=postgresql --values helm/postgresql-values.yaml
```

### Execute skaffold

In a local development environment it's best to use skaffold for building containers and executing helm.
Execute skaffold in the following way: `MINIKUBE_IP=$(minikube ip) skaffold dev`.

The minikube IP is required to let inway/outway/directory communicate with eachother via the ingresses, instead of internally.
Internally doesn't work because the internal hostnames for services (e.g. `directory-api.nlx-directory-dev`) do not match the TLS certificates (e.g. `directory-api.minikube`).

### Domains

The NLX demo simulation (used in environments `test`, `acc` and `demo`) is based on fictional communications between Haarlem, RDW and BRP. Ofcourse, this is just an example and the organizations themselves are not involved, so we have dedicated three domains to this simulation.

- `voorbeeld-haarlem.nl`
- `voorbeeld-rdw.nl`
- `voorbeeld-brp.nl`

If an update is required to one of these domains, please only modify voorbeeld-haarlem.nl, then copy the changes using TransIP's bulk copy feature to `voorbeeld-rdw.nl` and `voorbeeld-brp.nl`. This means that all three domains have exactly the same subdomains, which makes it easy to maintain them and keep them all in sync. For the simulation, `vorbeeld-haarlem.nl` doesn't run an inway, but it still has that record so it is copied to brp and rdw as well.

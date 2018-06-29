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

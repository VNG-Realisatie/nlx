# irma-server

This is the Chart for the irma server.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-irma-server`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the irma-server Chart
$ helm install --name my-irma-server nlx/irma-server
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-irma-server` deployment:

```console
$ helm delete my-irma-server
```

## Configuration

The following table lists the configurable parameters of the irma-server Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `registry.gitlab.com` |
| `image.repository` | Image repository | `commonground/nlx/irma` |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). | `0.4.1` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of directory replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | 
| `config.verbosity` | Set verbosity of logs (`1` include `DEBUG` messages, `2` include `TRACE` messages) | `""` |
| `config.emailAddress` | Email address that will receive notifications about changes in the IRMA software or ecosystem | `""` |
| `config.jwtKeyPEM` | Key which is used to sign IRMA sessions | `""` |
| `authentication.enabled` | If `true`, only request from authorized requestors are accepted | `true` |
| `authentication.requestors` | Requestor authentication settings. For more info see [IRMA documentation](https://irma.app/docs/irma-server/)| `{}`|
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.sessionPort` | Port exposed by the service for session requests | `8080` |
| `service.irmaPort` | Port exposed by the service for irma requests | `8080` |
| `podSecuritiyContext` | SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty. | `{}` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `` |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `ingress.enabled` | Enable Ingress | `false` |
| `ingress.class` | Ingress class | `""` |
| `ingress.annotations` | Ingress annotations | `{}` 
| `ingress.hosts` | Ingress accepted hostname | `chart-example.local` |
| `ingress.tls` | Ingress TLS configuration | `[]` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-irma-server -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/irma-server/values.yaml)

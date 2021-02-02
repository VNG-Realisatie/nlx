# review-page

This is the Chart for the review page.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-review-page`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-docs Chart
$ helm install --name my-review-page nlx/review-page
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-review-page` deployment:

```console
$ helm delete my-review-page
```

## Configuration

The following table lists the configurable parameters of the review-page Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/review-page` |
| `image.tag` | Image tag. When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of directory replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` |
| `config.environmentSlugWithDomain` | The enviroment slug with domain is added as a suffix to the URLs of the websites linked on the review page | `""` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.port` | Port exposed by the service for the review-page | `80` |
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
$ helm install --name my-review-page -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/review-page/values.yaml)

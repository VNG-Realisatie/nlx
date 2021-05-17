# ca-certportal 

This is the Chart for the NLX ca-certportal. It provides a 
web interface that allows requesting a testing certificate easily.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `ca-certportal`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-ca-certportal helm Chart
$ helm install ca-certportal nlx/nlx-ca-certportal
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `ca-certportal` deployment:

```console
$ helm delete ca-certportal
```

## Parameters

The following table lists the configurable parameters of the ca-certportal Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | x |
| `global.imageTag` | Global Docker Image tag | `true` | x |

### Common parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | x | 
| `fullnameOverride` | Override full deployment name | `""` | x |

### Deployment parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | x | 
| `image.repository` | Image repository | `nlxio/ca-certportal` | x |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` | x |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` | x |
| `image.pullSecrets` | Secrets for the image repository | `[]` | x |
| `affinity` | Node affinity for pod assignment | `{}` | x |
| `nodeSelector` | Node labels for pod assignment | `{}` | x |
| `replicaCount` | Number of management replicas | `1` | x |
| `resources` | Pod resource requests & limits | `{}` | x |
| `tolerations` | Node tolerations for pod assignment | `[]` | x |
| `serviceAccount.create` | If `true`, create a new service account | `true` | x |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` | x |
| `serviceAccount.annotations` | Annotations to add to the service account | x |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` | x |

### CA Cert Portal parameters
| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live | x |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `""` | x |
| `config.caHost` | The host of the Certificate Authority. | `""` | x |

### Exposure parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- | 
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | x |
| `service.httpPort` | Port exposed by the service | `8090` | x |
| `ingress.enabled` | Enable Ingress | `false` | x |
| `ingress.class` | Ingress class | `""` | x |
| `ingress.annotations` | Ingress annotations | `{}` | x |
| `ingress.hosts` | Ingress accepted hostnames | `[]` | x |
| `ingress.tls` | Ingres TLS configuration | `[]` | x |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install ca-certportal -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/ca-certportal/values.yaml)

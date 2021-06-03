# ca-cfssl-unsafe 

This is the Chart for the NLX ca-cfssl-unsafe. It is used in 
non-production environments to generate certificates on the fly from 
a central root CA.

Unsafe-ca is based on [cfssl](https://github.com/cloudflare/cfssl).

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `ca-cfssl-unsafe`:

```console
## add the Common Ground Helm repository
$ helm repo add commonground https://charts.commonground.nl

## Install the nlx-ca-cfssl-unsafe helm Chart
$ helm install ca-cfssl-unsafe nlx/nlx-ca-cfssl-unsafe
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `ca-cfssl-unsafe` deployment:

```console
$ helm delete ca-cfssl-unsafe
```

## Parameters

The following table lists the configurable parameters of the ca-cfssl-unsafe Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | no |
| `global.imageTag` | Global Docker Image tag | `true` | no |

### Common parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | no | 
| `fullnameOverride` | Override full deployment name | `""` | no |

### Deployment parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | no | 
| `image.repository` | Image repository | `nlxio/ca-cfssl-unsafe` | no |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` | no |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` | no |
| `image.pullSecrets` | Secrets for the image repository | `[]` | no |
| `affinity` | Node affinity for pod assignment | `{}` | no |
| `nodeSelector` | Node labels for pod assignment | `{}` | no |
| `replicaCount` | Number of management replicas | `1` | no |
| `resources` | Pod resource requests & limits | `{}` | no |
| `tolerations` | Node tolerations for pod assignment | `[]` | no |
| `serviceAccount.create` | If `true`, create a new service account | `true` | no |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` | no |
| `serviceAccount.annotations` | Annotations to add to the service account | `{}` | no | 
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` | no |
| `podSecuritiyContext` | SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty. | `{}` | no |

### CA cfssl unsafe parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `config.existingSecret` | Secret containing the root certificate and key of the CA | `""` | yes |

### Exposure parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- 
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | no |
| `service.port` | Port exposed by service | `8888` | no |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install ca-cfssl-unsafe -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/ca-cfssl-unsafe/values.yaml)

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
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

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
| `image.repository` | Image repository | `nlxio/ca-cfssl-unsafe` | x |
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
| `serviceAccount.annotations` | Annotations to add to the service account | `{}` | x | 
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` | x |
| `podSecuritiyContext` | SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty. | `{}` | x |

### CA cfssl unsafe parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `config.existingSecret` | Secret containing the root certificate and key of the CA | `""` | ✓ |

### Exposure parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- 
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | x |
| `service.port` | Port exposed by service | `8888` | x |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install ca-cfssl-unsafe -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/ca-cfssl-unsafe/values.yaml)

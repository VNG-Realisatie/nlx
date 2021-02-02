# ca-cfssl-unsafe 

This is the Chart for the NLX ca-cfssl-unsafe. It is used in 
non-production environments to generate certificates on the fly from 
a central root CA.

Unsafe-ca is based on [cfssl](https://github.com/cloudflare/cfssl).

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-ca-cfssl-unsafe`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-ca-cfssl-unsafe helm Chart
$ helm install --name my-ca-cfssl-unsafe nlx/nlx-ca-cfssl-unsafe
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-ca-cfssl-unsafe` deployment:

```console
$ helm delete my-ca-cfssl-unsafe
```

## Configuration

The following table lists the configurable parameters of the ca-cfssl-unsafe Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/ca-cfssl-unsafe` |
| `image.tag` | Image tag | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of ca-cfssl-unsafe replicas  | `1` |
| `config.existingSecret` | Secret containing the root certificate and key of the CA | `""` |
| `nameOverride`  | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | #TODO fullname -> fullName
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |  
| `podSecuritiyContext` | SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty. | `{}` |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.port` | Port exposed by service | `8888` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-ca-cfssl-unsafe -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/ca-cfssl-unsafe/values.yaml)

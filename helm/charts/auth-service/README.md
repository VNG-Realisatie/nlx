# auth-service 

This is the Chart for the NLX auth-service. The auth-service  can be used by an Outway 
to authorize all requests for a service, before routing the request to the targeted API in the NLX network.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-auth-service`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-auth-service helm Chart
$ helm install --name my-auth-service nlx/nlx-auth-service
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-auth-service` deployment:

```console
$ helm delete my-auth-service
```

## Configuration

The following table lists the configurable parameters of the nlx-auth-service Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/inway` |
| `image.tag` | Image tag | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount`  | Number of inway replicas  | `1` |
| `tls.certificatePEM` | Organization certificate | `""` |
| `tls.keyPEM` | The private key of `tls.certificatePEM` | `""` |
| `tls.existingSecret` | If you have an exisisting secret with your keypair you can use it instead of `tls.certificatePEM` and `tls.keyPEM` | `""` |
| `nameOverride`  | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | #TODO fullname -> fullName
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |  
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.port` | Port exposed by the inway service | `443` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-auth-service -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-auth-service/values.yaml)

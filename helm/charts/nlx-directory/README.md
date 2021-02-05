# nlx-directory 

This is the Chart for the NLX Directory.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `nlx-directory`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-nlx-directory helm Chart
$ helm install nlx-directory nlx/nlx-nlx-directory
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `nlx-directory` deployment:

```console
$ helm delete nlx-directory
```

## Configuration

The following table lists the configurable parameters of the nlx-directory Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `global.tls.rootCertificatePEM` | Root certificate of your internal PKI to be used by all NLX charts. If not set the value of `tls.certificate.rootCertificatePEM` is used | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.databaseRepository` | Image repository | `nlxio/directory-db` |
| `image.inspectionRepository` | Image repository | `nlxio/directory-inspection-api` |
| `image.registrationRepository` | Image repository | `nlxio/directory-registration-api` |
| `image.monitorRepository` | Image repository | `nlxio/directory-monitor` |
| `image.uiRepository` | Image repository | `nlxio/directory-ui` |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of directory replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `""` |
| `config.monitorOfflineServiceTTL` | Time, in seconds, a service can be offline before being removed from the directory | `86400` |
| `config.resetDatabase` | If `true` the database will be cleared after installing or upgrading | `false` |
| `tls.rootCertificatePEM` | The NLX root certificate | `""` |
| `tls.certificatePEM` | Organization certificate | `""` |
| `tls.keyPEM` | Private key of `tls.certificatePEM` | `""` |
| `tls.existingSecret` | TIf you have an existing secret with your NLX keypair you can use it instead of `tls.certificatePEM` and `tls.keyPEM` | `""` |
| `ui.enabled` | Enable the Directory UI | `true` |
| `ui.port` | Port exposed by the directory UI service | `80` |
| `ui.ingress.enabled` | Enable Ingress | `false` |
| `ui.ingress.class` | Ingress class | `""` |
| `ui.ingress.annotations` | Ingress annotations | `{}` 
| `ui.ingress.hosts` | Ingress accepted hostname | `chart-example.local` |
| `ui.ingress.tls` | Ingress TLS configuration | `[]` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.inspectionPort` | Port exposed by the service for the inspection API | `443` |
| `service.inspectionPlainPort` | Port exposed by the plain service for inspection API | `80` |
| `service.registrationPort` | Port exposed by the service for directory registration API | `443` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install nlx-directory -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-directory/values.yaml)

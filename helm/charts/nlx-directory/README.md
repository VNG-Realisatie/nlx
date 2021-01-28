# nlx-directory 

This is the Chart for the NLX Directory.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-nlx-directory`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-nlx-directory helm Chart
$ helm install --name my-nlx-directory nlx/nlx-nlx-directory
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-nlx-directory` deployment:

```console
$ helm delete my-nlx-directory
```

## Configuration

The following table lists the configurable parameters of the nlx-directory Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `global.tls.rootCertificatePEM` | Root certificate of your internal PKI to be used by all NLX charts. If not set the value of `tls.certificate.rootCertificatePEM` is used | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.databaseRepository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/directory-db` |
| `image.inspectionRepository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/directory-inspection-api` |
| `image.registrationRepository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/directory-registration-api` |
| `image.monitorRepository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/directory-monitor` |
| `image.uiRepository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/directory-ui` |
| `image.tag` | Image tag. When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of directory replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | #TODO fullname -> fullName
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `""` |
| `config.monitorOfflineServiceTTL` | Time, in seconds, a service can be offline before being removed from the directory | `86400` |
| `config.resetDatabase` | TODO | `false` |
| `tls.rootCertificatePEM` | TODO | `""` |
| `tls.certificatePEM` | TODO | `""` |
| `tls.keyPEM` | TODO | `""` |
| `tls.existingSecret` | TODO | `""` |
| `ui.enabled` | Enable the Directory UI | `true` |
| `ui.port` | TODO | `80` |
| `ui.ingress.enabled` | TODO | `false` |
| `ui.ingress.class` | TODO | `""` |
| `ui.ingress.annotations` | TODO | `{}` |
| `ui.ingress.hosts` | TODO | `[]` |
| `ui.ingress.tls` | TODO | `[]` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |
| `service.type` | TODO | `ClusterIP` |
| `service.inspectionPort` | TODO | `443` |
| `service.inspectionPlainPort` | TODO | `80` |
| `service.registrationPort` | TODO | `443` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-nlx-directory -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-nlx-directory/values.yaml)

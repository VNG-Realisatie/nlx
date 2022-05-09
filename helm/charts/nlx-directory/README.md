# nlx-directory

NLX is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API ecosystem with many organizations. the NLX directory provides a centralized overview of every participant in the NLX ecosystem. It is used by the NLX management API to identify which API's are available and at which address. It is only used for this insight. Once a connection (or even a access request to the API) is established, the directory has no role for that connection anymore.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `nlx-directory`:

```console
## add the Common Ground Helm repository
$ helm repo add commonground https://charts.commonground.nl

## Install the nlx-nlx-directory helm Chart
$ helm install nlx-directory commonground/nlx-nlx-directory
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

## Parameters

The following table lists the configurable parameters of the nlx-directory Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | no |
| `global.imageTag` | Global Docker Image tag | `true` | no |
| `global.tls.organization.rootCertificatePEM`| Global NLX root certificate. If not set the value of `tls.organization.rootCertificatePEM` is used | `nil` | no |

### Common parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | no |
| `fullnameOverride` | Override full deployment name | `""` | no |

### NLX directory parameters

| Parameter                             | Description                                                                                                                                                        | Default | Required (yes/no) |
|---------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------| -------- |
| `config.logType`                      | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.                    | `live`  | no |
| `config.logLevel`                     | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType`                                                              | `info`  | no |
| `config.monitor.offlineServiceTTL`    | The offline Time to Live (TTL) for a service in seconds. If the offline time of a service exceeds the time to live, the service will be removed from the directory | `86400` | no |
| `config.monitor.dbConnectionTimeout`  | The database connection timeout in seconds                                                                                                                         | `300`   | no |
| `config.monitor.dbMaxIdleConnections` | The maximum number of idle connections allowed to the database                                                                                                     | `5`     | no |
| `config.monitor.dbMaxOpenConnections` | The maximum number of open connections allowed to the database                                                                                                     | `25`    | no |
| `config.termsOfServiceURL`            | If a terms of service URL is specified, participants of this NLX network need to agree to the terms of service before they can use the network                     | `""`    | no |

### Deployment parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | no |
| `image.apiRepository` | Image repository | `nlxio/directory-api` | no |
| `image.monitorRepository` | Image repository | `nlxio/directory-monitor` | no |
| `image.uiRepository` | Image repository | `nlxio/directory-ui` | no |
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
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` | no |

## Postgres parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `postgresql.hostname` | PostgreSQL hostname | `postgresql` | no |
| `postgresql.connectTimeout` | The connection timeout for PostgreSQL | `10` | no |
| `postgresql.port` | PostgreSQL port | `5432` | yes |
| `postgresql.sslMode` | PostgreSQL SSL mode | `require` | yes |
| `postgresql.database` | PostgreSQL database  | `nlx-directory` | no |
| `postgresql.username` | PostgreSQL username. Will be stored in a kubernetes secret | `""` | no |
| `postgresql.password` | PostgreSQL password. Will be stored in a kubernetes secret | `""` | no |
| `postgresql.existingSecret.name` | Use existing secret for password details (`postgresql.username` and `postgresql.password` will be ignored and picked up from this secret)  | `""` | no |
| `postgresql.existingSecret.usernameKey` | Key for username value in aforementioned existingSecret | `username` | no |
| `postgresql.existingSecret.passwordKey` | Key for password value in aforementioned existingSecret | `password` | no |


### NLX TLS parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `tls.organization.rootCertificatePEM` | The NLX root certificate | `""` | yes (if global value is not set) |
| `tls.organization.certificatePEM` | The organization certificate | `""` | yes (if global value is not set) |
| `tls.organization.keyPEM` | Private key of `tls.organization.certificatePEM` | `""` | yes (if global value is not set) |
| `tls.organization.existingSecret` | Use existing secret with your NLX keypair (`tls.organization.certificatePEM` and `tls.organization.keyPEM` will be ignored and picked up from the secret) | `""` |  x |

### Exposure parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | no |
| `service.port` | Port exposed by the service for the directory API | `443` | no |
| `service.plainPort` | Port exposed by the plain service for directory API | `80` | no |
| `service.annotations` | Annotations for directory API | `{}` | no |
| `ui.enabled` | Enable the Directory UI | `true` | no |
| `ui.port` | Port exposed by the directory UI service | `80` | no |
| `ui.ingress.enabled` | Enable Ingress | `false` | no |
| `ui.ingress.class` | Ingress class | `""` | no |
| `ui.ingress.annotations` | Ingress annotations | `{}` | no |
| `ui.ingress.hosts` | Ingress accepted hostname | `chart-example.local` | no |
| `ui.ingress.tls` | Ingress TLS configuration | `[]` | no |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart.

```console
$ helm install nlx-directory -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-directory/values.yaml)

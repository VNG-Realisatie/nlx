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

### Deployment parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | no | 
| `image.databaseRepository` | Image repository | `nlxio/directory-db` | no |
| `image.inspectionRepository` | Image repository | `nlxio/directory-inspection-api` | no |
| `image.registrationRepository` | Image repository | `nlxio/directory-registration-api` | no |
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

### NLX Directory parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | `live` | no |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` | no |
| `config.directoryInspectionHostname` | Used to retrieve information about services from the directory. | `""` | yes | 
| `config.directoryRegistrationHostname` | Address of the NLX directory where this inway will register its services. | `""` | yes |
| `config.oidc.clientID` | The OIDC client ID | `"nlx-management"` | no |
| `config.oidc.clientSecret` | The OIDC client secret | `""` | yes |
| `config.oidc.discoveryURL` | The OIDC discovery URL | `""` | yes |
| `config.oidc.redirectURL` | The OIDC redirect URL | `""` | yes |
| `config.oidc.sessionSignKey` | The OIDC session sign key | `""` | yes |
| `config.monitorOfflineServiceTTL` | Time, in seconds, a service can be offline before being removed from the directory | `86400` | no |
| `config.resetDatabase` | If `true` the database will be cleared after installing or upgrading | `false` | no |

## Postgres parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `postgresql.hostname` | PostgreSQL hostname | `postgresql` | no |
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
| `service.inspection.port` | Port exposed by the service for the inspection API | `443` | no |
| `service.inspection.plainPort` | Port exposed by the plain service for inspection API | `80` | no |
| `service.inspection.annotations` | Annotations for inspection API | `{}` | no |
| `service.registration.port` | Port exposed by the service for directory registration API | `443` | no |
| `service.registration.annotations` | Annotations for registration API | `{}` | no |
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

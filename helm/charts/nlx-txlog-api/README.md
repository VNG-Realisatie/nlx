# TxlogAPI

NLX is an open source peer-to-peer system facilitating federated authentication, secure connecting and protocolling in a large-scale, dynamic API ecosystem with many organizations.
The transaction log API is a optional component which enables users expose the content of the transaction log database through an API.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `txlog-api`:

```console
## add the Common Ground Helm repository
$ helm repo add commonground https://charts.commonground.nl

## Install the nlx-txlog-api helm Chart
$ helm install txlog-api commonground/nlx-txlog-api
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version.
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md)
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `txlog-api` deployment:

```console
$ helm delete txlog-api
```

## Parameters

The following table lists the configurable parameters of the nlx-txlog-api Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | no |
| `global.imageTag` | Global Docker Image tag | `true` | no |
| `global.tls.internal.rootCertificatePEM` | Global root certificate of your internal PKI. If not set the value of `tls.rootCertificatePEM` is used | `nil` | no |

### Common parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | no |
| `fullnameOverride` | Override full deployment name | `""` | no |


### Deployment parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | no |
| `image.repository` | Image repository for the txlog API | `nlxio/txlog-api` | no |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` | no |
| `image.pullPolicy` | Image pull policy | `Always` | no |
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

### NLX TxLog API parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | `live` | no |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` | no |
| `config.name` | Unique identifier of this txlog API. | `""` | yes |

### TLS parameters

TLS certificates used by NLX components for internal communication.

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `tls.rootCertificatePEM` | The root certificate of your internal PKI | `""` | yes (if global value is not set) |
| `tls.certificatePEM` | The certificate signed by your internal PKI | `""` | yes |
| `tls.keyPEM` | The private key of `tls.certificatePEM` | `""` | yes |
| `tls.existingSecret` | Use existing secret with your NLX keypair (`tls.certificatePEM` and `tls.keyPEM` will be ignored and picked up from this secret) | `""` | no |

### Transaction log database parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `txlogdb.connectTimeout` | The connection timeout for PostgreSQL | `10` | no |
| `txlogdb.hostname` | PostgreSQL hostname | `""` | no |
| `txlogdb.port` | PostgreSQL port | `5432` | yes |
| `txlogdb.sslMode` | PostgreSQL SSL mode | `require` | yes |
| `txlogdb.database` | PostgreSQL database  | `""` | no |
| `txlogdb.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | no |
| `txlogdb.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | no |
| `txlogdb.existingSecret.name` | Use existing secret for password details (`transactionLog.username` and `transactionLog.password` will be ignored and picked up from this secret)  | `""` | no |
| `txlogdb.existingSecret.usernameKey` | Key for username value in aforementioned existingSecret | `username` | no |
| `txlogdb.existingSecret.passwordKey` | Key for password value in aforementioned existingSecret | `password` | no |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart.

```console
$ helm install txlog-api -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-txlog-api/values.yaml)

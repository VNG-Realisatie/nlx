# outway 

This is the Chart for the NLX outway. An outway is needed to consume an API on the NLX network.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `outway`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-outway helm Chart
$ helm install outway nlx/nlx-outway
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `outway` deployment:

```console
$ helm delete outway
```

## Parameters

The following table lists the configurable parameters of the nlx-outway Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | no |
| `global.imageTag` | Global Docker Image tag | `true` | no |
| `global.tls.organization.rootCertificatePEM`| Global NLX root certificate. If not set the value of `tls.organization.rootCertificatePEM` is used | `nil` | no |
| `global.tls.internal.rootCertificatePEM` | Global root certificate of your internal PKI. If not set the value of `tls.internal.rootCertificatePEM` is used | `nil` | no |

### Common parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | no | 
| `fullnameOverride` | Override full deployment name | `""` | no |


### Deployment parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | no | 
| `image.repository` | Image repository for the management API | `nlxio/outway` | no |
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

### NLX Management parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | `live` | no |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` | no |
| `config.directoryInspectionHostname` | Used to retrieve information about services from the directory. | `""` | yes |
| `config.name` | Unique identifier of this outway. | `""` | no |
| `config.authorizationService.enabled` | If `true`, the outway will use the authorization service | `false` | no |
| `config.authorizationService.url` | URL of the authorization service to use | `""` | no |
| `config.managementAPI.enabled` | If `true` the outway will use a management API to retrieve the service it will offer to the NLX network instead of using `config.ServiceConfig` | true | no |
| `config.managementAPI.address` | The config address of the management API. Normally this would be: `hostname:443` where `hostname` is the hostname of the Management API | `""` | no |

### TLS parameters

TLS certificate of your organization (used to communicate on the NLX Network).

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `tls.organization.rootCertificatePEM` | The NLX root certificate | `""` | yes (if global value is not set) |
| `tls.organization.certificatePEM` | Your NLX certificate | `""` | yes |
| `tls.organization.keyPEM` | The private key of `tls.organization.certificatePEM` | `""` | yes |
| `tls.organization.existingSecret` | Use existing secret with your NLX keypair (`tls.organization.certificatePEM` and `tls.organization.keyPEM` will be ignored and picked up from the secret) | `""` |  x |

TLS certificates used by NLX components for internal communication.

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `tls.internal.rootCertificatePEM` | The root certificate of your internal PKI | `""` | yes (if global value is not set) |
| `tls.internal.certificatePEM` | The certificate signed by your internal PKI | `""` | yes |
| `tls.internal.keyPEM` | The private key of `tls.internal.certificatePEM` | `""` | yes |
| `tls.internal.existingSecret` | Use existing secret with your NLX keypair (`tls.internal.certificatePEM` and `tls.internal.keyPEM` will be ignored and picked up from this secret) | `""` | no |

### Transaction Log parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `transactionLog.enabled` | If `true` the outway will write log records into the transaction log | `false` | no |
| `transactionLog.hostname` | PostgreSQL hostname | `""` | no |
| `transactionLog.port` | PostgreSQL port | `5432` | yes |
| `transactionLog.sslMode` | PostgreSQL SSL mode | `require` | yes |
| `transactionLog.database` | PostgreSQL database  | `""` | no |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | no |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | no |
| `transactionLog.existingSecret` | Use existing secret for password details (`transactionLog.username` and `transactionLog.password` will be ignored and picked up from this secret)  | `""` | no |

### Exposure parameters

| Parameter | Description | Default | Required (yes/no) |
| --------- | ----------- | ------- | -------- |
| `https.enabled` | If `true`, HTTPs will be enabled | `false` | no |
| `https.keyPEM` | Private key of `https.certificatePEM` as PEM. Required if `https.enabled` is `true` | `""` | no |
| `https.certificatePEM` | TLS certificate as PEM. Required if `https.enabled` is `true` | `""` | no |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | no |
| `service.apiPort` | Port exposed by the management API service | `80` | no |
| `service.apiConfigPort` | Port exposed by the management API service for the config endpoints | `443` | no |
| `service.uiPort` | Port exposed by the management UI service | `8080` | no |
| `ingress.enabled` | Enable Ingress | `false` | no |
| `ingress.annotations` | Ingress annotations | `{}` | no |
| `ingress.hosts.host` | Ingress accepted hostname | `chart-example.local` | no |
| `ingress.hosts.paths` | Ingress accepted paths | `[]` | no |
| `ingress.tls` | Ingress TLS configuration | `[]` | no |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install outway -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-outway/values.yaml)

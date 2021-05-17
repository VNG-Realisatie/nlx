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

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | x |
| `global.imageTag` | Global Docker Image tag | `true` | x |
| `global.tls.organization.rootCertificatePEM`| Global NLX root certificate. If not set the value of `tls.organization.rootCertificatePEM` is used | `nil` | x |
| `global.tls.internal.rootCertificatePEM` | Global root certificate of your internal PKI. If not set the value of `tls.internal.rootCertificatePEM` is used | `nil` | x |

### Common parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | x | 
| `fullnameOverride` | Override full deployment name | `""` | x |


### Deployment parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | x | 
| `image.repository` | Image repository for the management API | `nlxio/outway` | x |
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
| `serviceAccount.annotations` | Annotations to add to the service account | x |
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` | x |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` | x |

### NLX Management parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | `live` | x |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` | x |
| `config.directoryInspectionHostname` | Used to retrieve information about services from the directory. | `""` | ✓ |
| `config.name` | Unique identifier of this outway. | `""` | x |
| `config.authorizationService.enabled` | If `true`, the outway will use the authorization service | `false` | x |
| `config.authorizationService.url` | URL of the authorization service to use | `""` | x |
| `config.managementAPI.enabled` | If `true` the outway will use a management API to retrieve the service it will offer to the NLX network instead of using `config.ServiceConfig` | true | x |
| `config.managementAPI.address` | The config address of the management API. Normally this would be: `hostname:443` where `hostname` is the hostname of the Management API | `""` | x |

### NLX Management TLS parameters

TLS certificate of your organization (used to communicate on the NLX Network).

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `tls.organization.rootCertificatePEM` | The NLX root certificate | `""` | ✓ (if global value is not set) |
| `tls.organization.certificatePEM` | Your NLX certificate | `""` | ✓ |
| `tls.organization.keyPEM` | The private key of `tls.organization.certificatePEM` | `""` | ✓ |
| `tls.organization.existingSecret` | Use existing secret with your NLX keypair (`tls.organization.certificatePEM` and `tls.organization.keyPEM` will be ignored and picked up from the secret) | `""` |  x |

TLS certificates used by NLX components for internal communication.

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `tls.internal.rootCertificatePEM` | The root certificate of your internal PKI | `""` | ✓ (if global value is not set) |
| `tls.internal.certificatePEM` | The certificate signed by your internal PKI | `""` | ✓ |
| `tls.internal.keyPEM` | The private key of `tls.internal.certificatePEM` | `""` | ✓ |
| `tls.internal.existingSecret` | Use existing secret with your NLX keypair (`tls.internal.certificatePEM` and `tls.internal.keyPEM` will be ignored and picked up from this secret) | `""` | x |

### NLX Management Transaction Log parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `transactionLog.enabled` | If `true` the outway will write log records into the transaction log | `false` | x |
| `transactionLog.hostname` | PostgreSQL hostname | `""` | x |
| `transactionLog.port` | PostgreSQL port | `5432` | ✓ |
| `transactionLog.sslMode` | PostgreSQL SSL mode | `required` | ✓ |
| `transactionLog.database` | PostgreSQL database  | `""` | x |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | x |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | x |
| `transactionLog.existingSecret` | Use existing secret for password details (`transactionLog.username` and `transactionLog.password` will be ignored and picked up from this secret)  | `""` | x |

### Exposure parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `https.enabled` | If `true`, HTTPs will be enabled | `false` | x |
| `https.keyPEM` | Private key of `https.certificatePEM` as PEM. Required if `https.enabled` is `true` | `""` | x |
| `https.certificatePEM` | TLS certificate as PEM. Required if `https.enabled` is `true` | `""` | x |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | x |
| `service.apiPort` | Port exposed by the management API service | `80` | x |
| `service.apiConfigPort` | Port exposed by the management API service for the config endpoints | `443` | x |
| `service.uiPort` | Port exposed by the management UI service | `8080` | x |
| `ingress.enabled` | Enable Ingress | `false` | x |
| `ingress.annotations` | Ingress annotations | `{}` | x |
| `ingress.hosts.host` | Ingress accepted hostname | `chart-example.local` | x |
| `ingress.hosts.paths` | Ingress accepted paths | `[]` | x |
| `ingress.tls` | Ingress TLS configuration | `[]` | x |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install outway -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-outway/values.yaml)

# management 

This is the Chart for the NLX Management. NLX Management can be used to configure your NLX setup.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `management`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-management helm Chart
$ helm install management nlx/nlx-management
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `management` deployment:

```console
$ helm delete management
```

## Parameters

The following table lists the configurable parameters of the nlx-management Chart and its default values.

### Global parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `nil` | x |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` | x |
| `global.tls.organizationRootCertificatePEM`| NLX root certificate to be used by all NLX charts. If not set the value of `tls.organizationCertificate.rootCertificatePEM` is used | `nil` | x |
| `global.tls.rootCertificatePEM` | Root certificate of your internal PKI to be used by all NLX charts. If not set the value of `tls.certificate.rootCertificatePEM` is used | `nil` | x |

### Common parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | x | 
| `fullnameOverride` | Override full deployment name | `""` | x |

### Deployment parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | x | 
| `image.apiRepository` | Image repository for the management API | `nlxio/management-api` | x |
| `image.uiRepository` | Image repository for the management UI | `nlxio/management-ui` | x |
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
| `config.directoryRegistrationHostname` | Address of the NLX directory where this inway will register its services. | `""` | ✓ |
| `config.sessionCookieSecure` | If `true`, the API will use 'secure' cookies. | `"false"` | x |
| `config.oidc.clientID` | The OIDC client ID | `"nlx-management"` | x |
| `config.oidc.clientSecret` | The OIDC client secret | `""` | ✓ |
| `config.oidc.discoveryURL` | The OIDC discovery URL | `""` | ✓ |
| `config.oidc.redirectURL` | The OIDC redirect URL | `""` | ✓ |
| `config.oidc.sessionSignKey` | The OIDC session sign key | `""` | ✓ |

### NLX Management TLS parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `tls.organizationCertificate.rootCertificatePEM` | The NLX root certificate | `""` | ✓ (if global value is not set) |
| `tls.organizationCertificate.certificatePEM` | Your NLX certificate | `""` | ✓ |
| `tls.organizationCertificate.keyPEM` | The private key of `tls.organizationCertificate.certificatePEM` | `""` | ✓ |
| `tls.organizationCertificate.existingSecret` | Use existing secret with your NLX keypair (`tls.organizationCertificate.certificatePEM` and `tls.organizationCertificate.keyPEM` will be ignored and picked up from the secret) | `""` |  x |
| `tls.certificate.rootCertificatePEM` | The root certificate of your internal PKI | `""` | ✓ (if global value is not set) |
| `tls.certificate.certificatePEM` | The certificate signed by your internal PKI | `""` | ✓ |
| `tls.certificate.keyPEM` | The private key of `tls.certificate.certificatePEM` | `""` | ✓ |
| `tls.certificate.existingSecret` | Use existing secret with your NLX keypair (`tls.certificate.certificatePEM` and `tls.certificate.keyPEM` will be ignored and picked up from this secret) | `""` | x |

### NLX Management Transaction Log parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `transactionLog.enabled` | If `true` the outway will write log records into the transaction log | `false` | x |
| `transactionLog.hostname` | Hostname of the transaction log database | `""` | x |
| `transactionLog.database` | Database name of the transaction log | `""` | x |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | x |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` | x |
| `transactionLog.existingSecret` | Use existing secret for password details (`transactionLog.username` and `transactionLog.password` will be ignored and picked up from this secret)  | `""` | x |

### NLX Management PostgreSQL parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `postgresql.hostname` | PostgreSQL hostname | `"postgresql"` | ✓ |
| `postgresql.port` | PostgreSQL port | `5432` | ✓ |
| `postgresql.database` | PostgreSQL database | `"nlx_management"` | ✓ |
| `postgresql.username` | PostgreSQL username. Will be stored in a Kubernetes secret | `""` | ✓ (if not using `postgresql.existingSecret`) |
| `postgresql.password` | PostgreSQL password. Will be stored in a Kubernetes secret | `""` | ✓ (if not using `postgresql.existingSecret`) |
| `postgresql.existingSecret` | Use existing secret for password details (`postgresql.username` and `postgresql.password` will be ignored and picked up from this secret)  | `""` | x |

### Exposure parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
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
$ helm install management -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-management/values.yaml)

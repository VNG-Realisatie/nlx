# outway 

This is the Chart for the NLX outway. An outway is needed to consume an API on the NLX network.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-outway`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-outway helm Chart
$ helm install --name my-outway nlx/nlx-outway
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-outway` deployment:

```console
$ helm delete my-outway
```

## Configuration

The following table lists the configurable parameters of the nlx-outway Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `global.tls.organizationRootCertificatePEM`| NLX root certificate to be used by all NLX charts. If not set the value of `tls.organizationCertificate.rootCertificatePEM` is used | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/outway` |
| `image.tag` | Image tag. When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of outway replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | #TODO fullname -> fullName
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` | 
| `config.name` | Unique identifier of this outway. | `""` |
| `config.directoryInspectionHostname` | Used to retrieve information about services from the directory. | `""` |
| `config.authorizationService.enabled` | If `true`, the outway will use the authorization service | `false` |
| `config.authorizationService.url` | URL of the authorization service to use | `""` |
| `config.managementAPI.enabled` | If `true` the outway will use a management API to retrieve the service it will offer to the NLX network instead of using `config.ServiceConfig` | true |
| `config.managementAPI.address` | The address of the management API | `""` |
| `transactionLog.enabled` | If `true` the outway will write log records into the transaction log | `true` |
| `transactionLog.hostname` | Hostname of the transaction log database | `""` |
| `transactionLog.database` | Database name of the transaction log | `""` |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.existingSecret` | If you have an existing secret with PostgreSQL credentials you can use it instead of `transactionLog.username` and `transaction.password` | `""` |
| `tls.organizationCertificate.rootCertificatePEM` | The NLX root certificate | `""` |
| `tls.organizationCertificate.certificatePEM` | Your NLX certificate | `""` |
| `tls.organizationCertificate.keyPEM` | The private key of `tls.organizationCertificate.certificatePEM` | `""` |
| `tls.organizationCertificate.existingSecret` | If you have an existing secret with your NLX keypair you can use it instead of `tls.organizationCertificate.certificatePEM` and `tls.organizationCertificate.keyPEM` | `""` |
| `https.enabled` | If `true`, HTTPs will be enabled | `false` |
| `https.keyPEM` | Private key of `https.certificatePEM` as PEM. Required if `https.enabled` is `true` | `""` |
| `https.certificatePEM` | TLS certificate as PEM. Required if `https.enabled` is `true` | `""` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |  
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.httpPort` | Port exposed by the outway service when `https.enabled` is `false` | `80` |
| `service.httpsPort` | Port exposed by the outway service when `https.enabled` is `true`  | `443` |
| `ingress.enabled` | Enable Ingress | `false` |
| `ingress.annotations` | Ingress annotations | `{}` |
| `ingress.hosts.host` | Ingress accepted hostname | `chart-example.local` |
| `ingress.hosts.paths` | Ingress accepted paths | `[]` |
| `ingress.tls` | Ingress TLS configuration | `[]` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-outway -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-outway/values.yaml)

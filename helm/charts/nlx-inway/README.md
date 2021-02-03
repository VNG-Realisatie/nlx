# inway 

This is the Chart for the NLX inway. An inway is needed to offer an API on the NLX network.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-inway`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-inway helm Chart
$ helm install --name my-inway nlx/nlx-inway
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-inway` deployment:

```console
$ helm delete my-inway
```

## Configuration

The following table lists the configurable parameters of the nlx-inway Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `global.tls.organizationRootCertificatePEM`| NLX root certificate to be used by all NLX charts. If not set the value of `tls.organizationCertificate.rootCertificatePEM` is used | `""` |
| `global.tls.rootCertificatePEM` | Root certificate of your internal PKI to be used by all NLX charts. If not set the value of `tls.certificate.rootCertificatePEM` is used | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository | `nlxio/inway` |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of inway replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `""` | 
| `config.name` | Unique identifier of this inway. | `""` |
| `config.selfAddress` | The address that can be used by the NLX network to reach this inway | `""` |
| `config.directoryRegistrationHostname` | Address of the NLX directory where this inway can register its services | `""` |
| `config.serviceConfig` |**deprecated** The services this inway will offer to the NLX network. For more info see: https://docs.nlx.io/reference-information/service-configuration | `{}` |
| `config.managementAPI.enabled` | If `true` the inway will use a management API to retrieve the service it will offer to the NLX network instead of using `config.ServiceConfig` | true |
| `config.managementAPI.address` | The address of the management API | `""` |
| `transactionLog.enabled` | If `true` the inway will write log records into the transaction log | `true` |
| `transactionLog.hostname` | Hostname of the transaction log database | `""` |
| `transactionLog.database` | Database name of the transaction log | `""` |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.existingSecret` | If you have an existing secret with PostgreSQL credentials you can use it instead of `transactionLog.username` and `transaction.password` | `""` |
| `tls.organizationCertificate.rootCertificatePEM` | The NLX root certificate | `""` |
| `tls.organizationCertificate.certificatePEM` | Your NLX certificate | `""` |
| `tls.organizationCertificate.keyPEM` | The private key of `tls.organizationCertificate.certificatePEM` | `""` |
| `tls.organizationCertificate.existingSecret` | If you have an existing secret with your NLX keypair you can use it instead of `tls.organizationCertificate.certificatePEM` and `tls.organizationCertificate.keyPEM` | `""` |
| `tls.certificate.rootCertificatePEM` | The root certificate of your internal PKI | `""` |
| `tls.certificate.certificatePEM` | The certificate signed by your internal PKI | `""` |
| `tls.certificate.keyPEM` | The private key of `tls.certificate.certificatePEM` | `""` |
| `tls.certificate.existingSecret` | If you have an existing secret with your NLX keypair you can use it instead of `tls.organizationCertificate.certificatePEM` and `tls.organizationCertificate.keyPEM` | `""` |
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
$ helm install --name my-inway -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-inway/values.yaml)

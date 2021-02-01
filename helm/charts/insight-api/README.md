# insight-api 

This is the Chart for the NLX insight-api. An insight-api is needed to offer an API on the NLX network.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `my-insight-api`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-insight-api helm Chart
$ helm install --name my-insight-api nlx/nlx-insight-api
```

> **Tip**: List all releases using `helm list`

## Upgrading the Chart

Currently, our Helm charts use the same release version as the NLX release version. 
To know what has changed for the Helm charts, look at the changes in our [CHANGELOG](https://gitlab.com/commonground/nlx/nlx/-/blob/master/CHANGELOG.md) 
that are prefixed with 'Helm'.

## Uninstalling the Chart

To uninstall or delete the `my-insight-api` deployment:

```console
$ helm delete my-insight-api
```

## Configuration

The following table lists the configurable parameters of the insight-api Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.repository` | Image repository (ignored if `global.imageTag` is set) | `nlxio/insight-api` |
| `image.tag` | Image tag. When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of insight-api replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` | #TODO fullname -> fullName
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `""` | 
| `config.configurationFile` | Content for the insight-config.toml file | `""` |
| `config.jwt.signPrivateKeyPEM` | PEM RSA private key to sign requests for IRMA server | `""` |
| `config.jwt.irmaPublicKeyPEM` | PEM RSA public key to verify results from IRMA server | `""` |
| `transactionLog.enabled` | If `true` the insight-api will write log records into the transaction log | `true` |
| `transactionLog.hostname` | Hostname of the transaction log database | `""` |
| `transactionLog.database` | Database name of the transaction log | `""` |
| `transactionLog.username` | Username of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.password` | Password of the PostgreSQL user for the transaction log database. Will be stored in a kubernetes secret | `""` |
| `transactionLog.existingSecret` | If you have an existing secret with PostgreSQL credentials you can use it instead of `transactionLog.username` and `transaction.password` | `""` |
| `serviceAccount.create` | If `true`, create a new service account | `true` |
| `serviceAccount.name` | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template | `""` |
| `serviceAccount.annotations` | Annotations to add to the service account |  
| `securityContext` | Optional security context. The YAML block should adhere to the [SecurityContext spec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#securitycontext-v1-core) | `{}` |
| `podSecuritiyContext.fsGroup` | Group ID under which the pod should be started | `1001` |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` |
| `service.port` | Port exposed by service | `8080` |
| `ingress.enabled` | Enable Ingress | `false` |
| `ingress.annotations` | Ingress annotations | `{}` |
| `ingress.hosts.host` | Ingress accepted hostnames | `chart-example.local` |
| `ingress.hosts.paths` | Ingress paths | `[]` |
| `ingress.tls` | Ingress TLS configuration | `[]` |
| `resources` | Pod resource requests & limits | `{}` |
| `nodeSelector` | Node labels for pod assignment | `{}` |
| `affinity` | Node affinity for pod assignment | `{}` |
| `tolerations` | Node tolerations for pod assignment | `[]` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install --name my-insight-api -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-insight-api/values.yaml)

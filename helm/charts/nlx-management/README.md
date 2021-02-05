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

## Configuration

The following table lists the configurable parameters of the nlx-management Chart and its default values.

| Parameter | Description | Default |
| --------- | ----------- | ------- |
| `global.imageRegistry` | Image registry to be used by all NLX charts | `""` |
| `global.imageTag` | Image tag to be used by all NLX charts | `true` |
| `global.tls.organizationRootCertificatePEM`| NLX root certificate to be used by all NLX charts. If not set the value of `tls.organizationCertificate.rootCertificatePEM` is used | `""` |
| `global.tls.rootCertificatePEM` | Root certificate of your internal PKI to be used by all NLX charts. If not set the value of `tls.certificate.rootCertificatePEM` is used | `""` |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` |
| `image.apiRepository` | Image repository for the management API | `nlxio/management-api` |
| `image.uiRepository` | Image repository for the management UI | `nlxio/management-ui` |
| `image.tag` | Image tag (ignored if `global.imageTag` is set). When set to null, the AppVersion from the Chart is used | `The appVersion from the chart` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `image.pullSecrets` | Secrets for the image repository | `[]` |
| `replicaCount` | Number of management replicas | `1` |
| `nameOverride` | Override deployment name | `""` |
| `fullnameOverride` | Override full deployment name | `""` |
| `config.logType` | Possible values: **live**, **local**. Affects the log output. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger. | live |
| `config.logLevel` | Possible values: **debug**, **warn**, **info**. Override the default loglevel set by `config.logType` | `info` |
| `config.directoryInspectionHostname` | Used to retrieve information about services from the directory. | `""` |
| `config.directoryRegistrationHostname` | Address of the NLX directory where this inway can register its services. | `""` |
| `config.sessionCookieSecure` | If `true`, the API will use 'secure' cookies. | `"false"` |
| `config.secretKey` | Secret key that is used for signing sessions | `""` |
| `config.oidc.clientID` | The OIDC client ID | `"nlx-management"` |
| `config.oidc.clientSecret` | The OIDC client secret | `""` |
| `config.oidc.discoveryURL` | The OIDC discovery URL | `""` |
| `config.oidc.redirectURL` | The OIDC redirect URL | `""` |
| `config.oidc.sessionSignKey` | The OIDC session sign key | `""` |
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
| `service.apiPort` | Port exposed by the management API service | `80` |
| `service.apiConfigPort` | Port exposed by the management API service for the config endpoints | `443` |
| `service.uiPort` | Port exposed by the management UI service | `8080` |
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
$ helm install management -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-management/values.yaml)

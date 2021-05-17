# nlx-directory 

This is the Chart for the NLX Directory.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the Chart with the release name `nlx-directory`:

```console
## Add the NLX Helm repository
$ helm repo add nlx https://charts.nlx.io

## Install the nlx-nlx-directory helm Chart
$ helm install nlx-directory nlx/nlx-nlx-directory
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

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `global.imageRegistry` | Global Docker Image registry | `nil` | x |
| `global.imageTag` | Global Docker Image tag | `true` | x |
| `global.tls.rootCertificatePEM` | Global root certificate of your internal PKI. If not set the value of `tls.internal.rootCertificatePEM` is used | `nil` | x |

### Common parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `nameOverride` | Override deployment name | `""` | x | 
| `fullnameOverride` | Override full deployment name | `""` | x |

### Deployment parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `image.registry` | Image registry (ignored if `global.imageRegistry` is set) | `docker.io` | x | 
| `image.databaseRepository` | Image repository | `nlxio/directory-db` | x |
| `image.inspectionRepository` | Image repository | `nlxio/directory-inspection-api` | x |
| `image.registrationRepository` | Image repository | `nlxio/directory-registration-api` | x |
| `image.monitorRepository` | Image repository | `nlxio/directory-monitor` | x |
| `image.uiRepository` | Image repository | `nlxio/directory-ui` | x |
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

### NLX Directory parameters

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
| `config.monitorOfflineServiceTTL` | Time, in seconds, a service can be offline before being removed from the directory | `86400` | x |
| `config.resetDatabase` | If `true` the database will be cleared after installing or upgrading | `false` | x |

### NLX Directory TLS parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `tls.organization.rootCertificatePEM` | The NLX root certificate | `""` | ✓ (if global value is not set) |
| `tls.organization.certificatePEM` | The organization certificate | `""` | ✓ (if global value is not set) |
| `tls.organization.keyPEM` | Private key of `tls.organization.certificatePEM` | `""` | ✓ (if global value is not set) |
| `tls.organization.existingSecret` | Use existing secret with your NLX keypair (`tls.organization.certificatePEM` and `tls.organization.keyPEM` will be ignored and picked up from the secret) | `""` |  x |

### Exposure parameters

| Parameter | Description | Default | Required |
| --------- | ----------- | ------- | -------- |
| `service.type` | Service type (ClusterIP, NodePort or LoadBalancer) | `ClusterIP` | x |
| `service.inspection.port` | Port exposed by the service for the inspection API | `443` | x |
| `service.inspection.plainPort` | Port exposed by the plain service for inspection API | `80` | x |
| `service.inspection.annotations` | Annotations for inspection API | `{}` | x |
| `service.registration.port` | Port exposed by the service for directory registration API | `443` | x |
| `service.registration.annotations` | Annotations for registration API | `{}` | x |
| `ui.enabled` | Enable the Directory UI | `true` | x |
| `ui.port` | Port exposed by the directory UI service | `80` | x |
| `ui.ingress.enabled` | Enable Ingress | `false` | x |
| `ui.ingress.class` | Ingress class | `""` | x |
| `ui.ingress.annotations` | Ingress annotations | `{}` | x |
| `ui.ingress.hosts` | Ingress accepted hostname | `chart-example.local` | x |
| `ui.ingress.tls` | Ingress TLS configuration | `[]` | x |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. 

```console
$ helm install nlx-directory -f values.yaml .
```
> **Tip**: You can use the default [values.yaml](https://gitlab.com/commonground/nlx/nlx/blob/master/helm/charts/nlx-directory/values.yaml)

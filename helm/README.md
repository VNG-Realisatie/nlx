# Helm charts for NLX

This directory contains Helm charts to deploy various NLX components to a Kubernetes cluster.

- `charts` charts for each individual component
- `deploy` meta charts used for deployment, uses charts from the `charts` directory

The charts `version` and `appVersion` follows the version of this repository.


## Deploy charts

| Chart                 | Description                                                                             |
| --------------------- | --------------------------------------------------------------------------------------- |
| rvrd                  | Demo organization: Rijksdienst voor Realistische Demo's, provides an API                |
| directory             | Used for deployment to pre-production and production                                    |
| gemeente-stijns       | Demo municipality, consumes APIs from RvRD                                              |
| vergunningsoftware-bv | Demo Saas organization, consumes a service in the name of municipality Gemeente Stijns  |
| shared                | Shared components that the RvRD and Gemeente Stijns use                                 |

> *Note: The NLX demo simulation is based on fictional communications between Gemeente Stijns, RvRD and Vergunningsoftware BV. This is just an example with made up the organizations.*


## Environments

Charts in the `deploy` directory are used to deploy to different environments. The default values in `values.yaml` are suited for development only. To deploy to other environments use the `values-<environment>.yaml` file.


### Review environment

Used for testing code changes. No NLX directory or Inway is exposed outside the Kubernetes cluster. Before deployment the variables in the `values-review.yaml.tpl` file are replaced with their actual values. This is because we use dynamic domain names in this environment.

*Charts: rvrd, gemeente-stijns, vergunningsoftware-bv, shared*


### Acceptance environment

All changes pushed to the master branch are deployed to this environment. No NLX directory or Inway is exposed outside the Kubernetes cluster.

*Charts: rvrd, gemeente-stijns, vergunningsoftware-bv, shared*


### Demo environment

Used for demonstrating the NLX ecosystem. The NLX directory and Inways are publicly exposed and can be accessed from outside the Kubernetes cluster.

*Charts: rvrd, gemeente-stijns, vergunningsoftware-bv, shared*


### Producion and pre-production environments

Only a NLX directory is deployed. PKIOverheid certificates are used and are already present on the Kubernetes cluster.

*Charts: directory*

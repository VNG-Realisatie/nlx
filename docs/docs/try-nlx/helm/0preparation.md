---
id: 0preparation
title: 0. Preparation
---

# 0. Preparation

Be sure to have all mentioned requirements in place before continuing to step 1.

## Cluster

Please have a Kubernetes cluster available including sufficient authorization. The cluster must be at least version v1.18.0.

_Make sure the default kubectl cluster is set to the correct cluster for the duration of this guide._

## Domain names

Make sure you have a domain name ready that can be used to make the NLX inway publicly available. The inway, which you will install, will create a load balancer on your Kubernetes cluster. Traffic must be allowed to this load balancer (maybe there are firewall rules?). The domain you want to use for your inway must point to the IP of this load balancer. In addition, a domain is also required for NLX management. This domain should route traffic to your Kubernetes cluster.

## Working directory

While following this installation guide you will be asked several times to download or change files. In this guide, we assume that you save all these files in the same directory and execute the commands from this directory with the terminal.

```
mkdir your-directory
cd your-directory
```

## Tooling

* [Kubectl, version v1.19.0 minimum](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-kubectl/)
* [Helm, minimum v3.2.0](https://helm.sh/docs/intro/install/)
* [Homebrew package manager](https://brew.sh) **(only needed for MacOS users)**
* [OpenSSL](https://www.openssl.org/source/)
**MacOS users**
Openssl on Mac OS is not suitable for creating V3 CA certificates. Therefore only for Mac OS users:` brew install openssl@1.1`
**Windows users**
Install OpenSSL with the following command: `choco install OpenSSL.Light`
* [Git](https://git-scm.com/docs/git-archive)

## Download base files

Now get the necessary base files with:
```
curl --location \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/internal-issuer.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/inway-internal-tls.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/job-create-administrator.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/management-internal-tls.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/nlx-inway-values.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/nlx-management-values.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/nlx-outway-values.yaml \
    --remote-name https://gitlab.com/commonground/nlx/nlx/-/raw/master/technical-docs/nlx-helm-installation-guide/outway-internal-tls.yaml
```

Check the files are now in your working directory with:

```
ls
```

The following files must be available:

```
internal-issuer.yaml
inway-internal-tls.yaml 
job-create-administrator.yaml
management-internal-tls.yaml 
nlx-inway-values.yaml
nlx-management-values.yaml
nlx-outway-values.yaml 
outway-internal-tls.yaml
```


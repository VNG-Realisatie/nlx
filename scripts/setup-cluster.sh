#!/bin/bash

set -e

echo ">> setup traefik"
helm repo add traefik https://containous.github.io/traefik-helm-chart && \
    helm repo update && \
    kubectl create namespace traefik && \
    helm install traefik traefik/traefik --namespace traefik --values helm/traefik-values-minikube.yaml

echo ">> setup kubedb"
helm install kubedb-operator appscode/kubedb --version 0.12.0 --namespace kube-system && \
    helm install kubedb-catalog appscode/kubedb-catalog --version 0.12.0 --namespace kube-system

echo ">> setup certmanager"
helm repo add jetstack https://charts.jetstack.io && \
    helm repo update && \
    kubectl create namespace cert-manager && \
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.16.1/cert-manager.crds.yaml && \
    helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v0.16.1

echo ">> building helm charts"
helm repo add stable https://charts.helm.sh/stable && \
    helm dependency build ./helm/deploy/gemeente-stijns && \
    helm dependency build ./helm/deploy/rvrd

echo ">> installing shared helm chart"
helm upgrade --install shared ./helm/deploy/shared

echo ">> installing rvrd helm chart"
helm upgrade --install rvrd ./helm/deploy/rvrd

echo ">> installing gemeente-stijns helm chart"
helm upgrade --install gemeente-stijns ./helm/deploy/gemeente-stijns

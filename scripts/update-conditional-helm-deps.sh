#!/bin/sh

helm repo add stable https://charts.helm.sh/stable && \
helm repo update && \
helm dependency update helm/deploy/directory && \
helm dependency update helm/deploy/gemeente-riemer && \
helm dependency update helm/deploy/gemeente-stijns && \
helm dependency update helm/deploy/vergunningsoftware-bv && \
helm dependency update helm/deploy/rvrd && \
helm dependency update helm/deploy/shared


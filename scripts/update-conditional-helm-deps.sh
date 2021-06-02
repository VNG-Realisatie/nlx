#!/bin/sh

helm repo add stable https://charts.helm.sh/stable && \
  helm repo update && \
  helm dependency update helm/deploy/gemeente-stijns
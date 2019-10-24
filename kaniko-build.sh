#!/bin/bash

set -e

docker run \
    -v $(pwd):/workspace \
    gcr.io/kaniko-project/warmer:latest \
    --cache-dir=/workspace/cache \
    --image=nginx:alpine \
    --image=node:10.16.3 \
    --image=golang:1.13.3-alpine \
    --image=alpine:latest

# auth-service
docker run \
    -v $(pwd):/workspace \
    gcr.io/kaniko-project/executor:latest \
    --cache-dir=/workspace/cache \
    --context dir:///workspace \
    --dockerfile auth-service/Dockerfile \
    --destination minikube:5000/nlxio/auth-service \
    --insecure \
    --no-push

# ca-cfssl-unsafe
docker run \
    -v $(pwd):/workspace \
    gcr.io/kaniko-project/executor:latest \
    --cache-dir=/workspace/cache \
    --context dir:///workspace/ca-cfssl-unsafe \
    --dockerfile Dockerfile \
    --destination minikube:5000/nlxio/ca-cfssl-unsafe \
    --insecure \
    --no-push

# docs
docker run \
    -v $(pwd):/workspace \
    gcr.io/kaniko-project/executor:latest \
    --cache-dir=/workspace/cache \
    --context dir:///workspace/docs \
    --dockerfile Dockerfile \
    --destination minikube:5000/nlxio/docs \
    --insecure \
    --no-push

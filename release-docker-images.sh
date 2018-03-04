#!/bin/bash

set -e # exit on error
set -x # echo commands

RELEASE_TAG=${RELEASE_TAG:-latest}

docker build \
	-t nlxio/docs:${RELEASE_TAG} \
	-f docs/Dockerfile .

docker build \
	-t nlxio/unsafe-ca:${RELEASE_TAG} \
	-f unsafe-ca/Dockerfile .

docker build \
	-t nlxio/directory:${RELEASE_TAG} \
	-f directory/Dockerfile .

docker build \
	-t nlxio/inway:${RELEASE_TAG} \
	-f inway/Dockerfile .

docker build \
	-t nlxio/outway:${RELEASE_TAG} \
	-f outway/Dockerfile .

docker build \
	-t nlxio/directory-ui:${RELEASE_TAG} \
	-f directory-ui/Dockerfile .

# TODO: only push the image when this script is ran in CI/CD or forced using env var (backup-plan for when CI/CD is down/unavailable)
docker push nlxio/docs:${RELEASE_TAG}
docker push nlxio/unsafe-ca:${RELEASE_TAG}
docker push nlxio/directory:${RELEASE_TAG}
docker push nlxio/inway:${RELEASE_TAG}
docker push nlxio/outway:${RELEASE_TAG}
docker push nlxio/certportal:${RELEASE_TAG}
docker push nlxio/directory-ui:${RELEASE_TAG}

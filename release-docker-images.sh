#!/bin/bash

set -e # exit on error
set -x # echo commands

RELEASE_TAG='latest' # Local builds use latest
if [[ ! -z "$CI_COMMIT_TAG" ]]; then
	# Build tagged releases using the version tag
	RELEASE_TAG=${CI_COMMIT_TAG}
elif [[ ! -z "$CI_COMMIT_SHA" ]]; then
	# When not tagged but still on CI, use the commit SHA.
	RELEASE_TAG=${CI_COMMIT_SHA}
fi

docker build \
	-t nlxio/docs:latest -t nlxio/docs:${RELEASE_TAG} \
	-f docs/Dockerfile .

docker build \
	-t nlxio/unsafe-ca:latest -t nlxio/unsafe-ca:${RELEASE_TAG} \
	-f unsafe-ca/Dockerfile .

docker build \
	-t nlxio/db:latest -t nlxio/db:${RELEASE_TAG} \
	-f db/Dockerfile .

docker build \
	-t nlxio/directory:latest -t nlxio/directory:${RELEASE_TAG} \
	-f directory/Dockerfile .

docker build \
	-t nlxio/inway:latest -t nlxio/inway:${RELEASE_TAG} \
	-f inway/Dockerfile .

docker build \
	-t nlxio/outway:latest -t nlxio/outway:${RELEASE_TAG} \
	-f outway/Dockerfile .

docker build \
	-t nlxio/directory-ui:latest -t nlxio/directory-ui:${RELEASE_TAG} \
	-f directory-ui/Dockerfile .

docker build \
	-t nlxio/certportal:latest -t nlxio/certportal:${RELEASE_TAG} \
	-f certportal/Dockerfile .

docker build \
	-t nlxio/monitor:latest -t nlxio/monitor:${RELEASE_TAG} \
	-f monitor/Dockerfile .
	
docker build \
	-t nlxio/logdb:latest -t nlxio/logdb:${RELEASE_TAG} \
	-f logdb/Dockerfile .
	
docker build \
	-t nlxio/logdb-api:latest -t nlxio/logdb-api:${RELEASE_TAG} \
	-f logdb-api/Dockerfile .
	
docker build \
	-t nlxio/logdb-ui:latest -t nlxio/logdb-ui:${RELEASE_TAG} \
	-f logdb-ui/Dockerfile .

# Only push the image when this script is ran in CI/CD or forced using env var (backup-plan for when CI/CD is down/unavailable)
if [ "${RELEASE_TAG}" != "latest" ]; then
	docker push nlxio/docs:latest
	docker push nlxio/docs:${RELEASE_TAG}
	docker push nlxio/unsafe-ca:latest
	docker push nlxio/unsafe-ca:${RELEASE_TAG}
	docker push nlxio/db:latest
	docker push nlxio/db:${RELEASE_TAG}
	docker push nlxio/directory:latest
	docker push nlxio/directory:${RELEASE_TAG}
	docker push nlxio/inway:latest
	docker push nlxio/inway:${RELEASE_TAG}
	docker push nlxio/outway:latest
	docker push nlxio/outway:${RELEASE_TAG}
	docker push nlxio/certportal:latest
	docker push nlxio/certportal:${RELEASE_TAG}
	docker push nlxio/directory-ui:latest
	docker push nlxio/directory-ui:${RELEASE_TAG}
	docker push nlxio/monitor:latest
	docker push nlxio/monitor:${RELEASE_TAG}
	docker push nlxio/logdb:latest
	docker push nlxio/logdb:${RELEASE_TAG}
	docker push nlxio/logdb-api:latest
	docker push nlxio/logdb-api:${RELEASE_TAG}
	docker push nlxio/logdb-ui:latest
	docker push nlxio/logdb-ui:${RELEASE_TAG}
fi

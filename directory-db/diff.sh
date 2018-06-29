#!/bin/bash

## This diff script helps in development of the database. It starts a container which runs diff/modd.sh
## The diff/modd.sh file watches for changes to the model or migrations, and verifies that the migrations match the model.

set -e

dockerCmd=''
dockerRunArgs='-ti'
if [ $1 = 'ci-once' ]
then
    # When in CI/CD, run the test only once and exit the status
    dockerCmd='./diff/calc-model-diff.sh'
    dockerRunArgs=''
fi

NLXROOT=$(git rev-parse --show-toplevel)

# Not using nlxio/ prefix in --tag image name because this image isn't meant to be released; only for local/development use.
docker build \
    --tag nlx-directory-db-diff:latest \
    --file ${NLXROOT}/directory-db/diff/Dockerfile \
    ${NLXROOT}/directory-db

docker run ${dockerRunArgs} \
    --volume ${NLXROOT}:/go/src/go.nlx.io/nlx \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    nlx-directory-db-diff:latest ${dockerCmd}

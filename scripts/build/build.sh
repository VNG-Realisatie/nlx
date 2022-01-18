#!/bin/bash -eu

DOCKER_BUILDKIT=1
BUILDKIT_INLINE_CACHE=1
export DOCKER_BUILDKIT

jq_bin=jq
if [ -n "$JQ_BIN_PATH" ]; then jq_bin=$JQ_BIN_PATH; fi

pids=()
images=$($jq_bin -r '.images' "./scripts/build/images.json")

wait_for_pids() {
    for pid in "${pids[@]}"; do
    if wait -n; then
        :
    else
        status=$?
        echo "One of the subprocesses exited with nonzero status $status. Aborting."
        for pid in "${pids[@]}"; do
        # Send a termination signal to all the children, and ignore errors
        # due to children that no longer exist.
        kill "$pid" 2> /dev/null || :
        done
        exit "$status"
    fi
    done
}

_term() { 
    echo "Killing leftover child processes"
    for pid in "${pids[@]}"; do
        # Send a termination signal to all the children, and ignore errors
        # due to children that no longer exist.
        kill "$pid" 2> /dev/null || :
    done 
}

# Trap EXIT and SIGKILL signal so child processes can be cleaned up
trap _term EXIT
trap _term SIGKILL

# Build all images
for row in $(echo "${images}" | $jq_bin -r '.[] | @base64'); do
    _jq()
    {
        echo ${row} | base64 --decode | $jq_bin -r ${1}
    }

    cmd="docker build $(_jq '.build.context') -f $(_jq '.build.dockerfile') --build-arg BUILDKIT_INLINE_CACHE=${BUILDKIT_INLINE_CACHE}"

    for tag in $(echo "$(_jq '.tags')" | $jq_bin -r '.[]'); do
        cmd+=" -t ${tag}"
    done;

    for arg in $(echo "$(_jq '.build.args')" | $jq_bin -r '.[]?'); do
        cmd+=" --build-arg ${arg}"
    done;

    for cache in $(echo "$(_jq '.build.cache_from')" | $jq_bin -r '.[]?'); do
        cmd+=" --cache-from ${cache}"
    done;
    
    echo "${cmd}" | envsubst
    (eval $(echo "${cmd}" | envsubst)) &
    pids+=($!)
done;
wait_for_pids

echo "Building images done"

pids=()

# Push all images
for row in $(echo "${images}" | $jq_bin -r '.[] | @base64'); do
    _jq()
    {
        echo ${row} | base64 --decode | $jq_bin -r ${1}
    }

    for tag in $(echo "$(_jq '.tags')" | $jq_bin -r '.[]'); do
        (eval $(echo "docker push ${tag}" | envsubst)) &
        pids+=($!)
    done;
done;
wait_for_pids

echo "Pushing images done"

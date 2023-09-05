#!/bin/bash

set -eo pipefail

cd "$(dirname "$0")"

docker_build_arg=""
if [ ! -z "${DOCKER_FROM}" ]; then
    docker_build_arg=" --build-arg DOCKER_FROM=${DOCKER_FROM} "
fi

export DOCKER_BUILDKIT=1

sudo docker build -t pcre2-build-linux $docker_build_arg -f Dockerfile .

sudo docker run --user "$(id -u)" --rm -v "$(pwd):/project" -w /project -v "$(pwd)/../build:/build" \
    -e "PREFIX_BASE=/build" -e "ARCH=${ARCH}" pcre2-build-linux ./build.sh

#!/bin/bash -e
PLATFORM_DOCKER_DIR=${PLATFORM_DOCKER_DIR:-../decentage-platform-docker}
source "${PLATFORM_DOCKER_DIR}/scripts/base-build-scripts.sh"
run_golang_in_docker "decentage-gatekeeper" "${ROOT_DIR}" "make"

#!/bin/bash

set -e

COMMIT8=${TRAVIS_COMMIT::8}

docker build -t ed201971/godaddy:$COMMIT8 devops/.
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin

docker push ed201971/godaddy:$COMMIT8

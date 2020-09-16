#!/bin/bash

set -e

COMMIT8=${TRAVIS_COMMIT::8}
COMMIT=${TRAVIS_COMMIT}

#npm install -g npm
#npm install
#npm list
#npm run deploy

docker build -t ed201971/godaddy devops/.
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
#npx semantic-release

echo ¢COMMIT8
echo ¢COMMIT

#docker push ed201971/godaddy

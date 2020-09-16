#!/bin/bash

set -e

npm install -g npm
npm install
#npm list
npm run deploy

docker build -t ed201971/godaddy devops/.
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin
npx semantic-release

#docker push ed201971/godaddy

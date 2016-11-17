#!/usr/bin/env bash

set -e

# first argument passed in is used for the image name
dockerTag=$1
echo $dockerTag
containerName=tmp_watcher
docker rm $containerName || true
docker build -t $containerName .
# get build binary from container
id=$(docker create $containerName)
docker cp $id:/go/bin/configmap-reload .
docker rm -v $id

docker build -t $dockerTag -f Dockerfile.deploy .
rm configmap-reload
docker push $dockerTag
#!/bin/bash

# Check if docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo 'Unable to find docker command, please install Docker (https://www.docker.com/) and retry' >&2
  exit 1
fi

echo "Creating Service Network (service) if ir doesn't not exist"
# Create network func_function if doesn't exists
[ ! "$(docker network ls | grep service)" ] && docker network create -d overlay --attachable service

echo "Deploying device stack"
# Deploy the docker stack
docker stack deploy --compose-file docker-compose.yml magic-home-cloud
echo "Device stack succesfully deployed"

#!/bin/sh

export dockerfile="Dockerfile"
export arch=$(uname -m)

export eTAG="0.0.1"

echo Building magichome/device-state:$eTAG

docker build -t magichome/device-state:$eTAG . -f $dockerfile

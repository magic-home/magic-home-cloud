#!/bin/sh

export dockerfile="Dockerfile"
export arch=$(uname -m)

export eTAG="0.0.1"

echo Building magichome/device-conf:$eTAG

docker build -t magichome/device-conf:$eTAG . -f $dockerfile

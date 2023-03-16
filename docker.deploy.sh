#!/usr/bin/bash

IMAGE_NAME=docker.happyelements.com/muggle/sh01-feishu-chatgpt
TAG=latest
DOCKER_FILE=Dockerfile
if [ "$1" == "debug" ]; then
    DOCKER_FILE=Dockerfile.debug
fi

docker build --platform linux/amd64 -t $IMAGE_NAME:$TAG -f $DOCKER_FILE .

docker image push $IMAGE_NAME:$TAG
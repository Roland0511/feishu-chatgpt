IMAGE_NAME:=docker.happyelements.com/muggle/sh01-feishu-chatgpt
TAG:=latest
all: docker docker.push

docker:
	docker build --platform linux/amd64 -t $(IMAGE_NAME):$(TAG) -f Dockerfile .

docker.debug:
	docker build --platform linux/amd64 -t $(IMAGE_NAME):$(TAG) -f Dockerfile.debug .

docker.push:
	docker push $(IMAGE_NAME):$(TAG)
#!/usr/bin/bash
IMAGE_NAME=docker.happyelements.com/muggle/sh01-feishu-chatgpt
TAG=latest

docker run -it --rm --name sh01-feishu-chatgpt -p 9000:9000 -p 40000:40000 \
--env APP_ID=cli_a49b26bf26fb500b \
--env APP_SECRET=ij59LPZsSvIBdZ3ws0gkH5FcxcyPjsHE \
--env APP_ENCRYPT_KEY=L0owdz6QqPXO1EinsC3bXbRveSjsCpeg \
--env APP_VERIFICATION_TOKEN=oFQAZD1qlB71qnSyEqMcFcI3fn84xVMW \
--env BOT_NAME=毛毛 \
--env OPENAI_KEY="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" \
$IMAGE_NAME:$TAG
#!/bin/sh

if [ "$1" = "" ]; then
  echo  "tagを指定してください .e.g. v1.0.0"
  exit 1
fi

imageTag="$1"
nginxRepositoryName="gcr.io/${GCP_PROJECT_ID}/portfolio-backend-nginx"
goLangRepositoryName="gcr.io/${GCP_PROJECT_ID}/portfolio-backend-go"

docker build --no-cache --rm -t ${nginxRepositoryName}:latest -f docker/nginx/Dockerfile .
docker tag ${nginxRepositoryName}:latest ${nginxRepositoryName}:${imageTag}
docker push ${nginxRepositoryName}:latest
docker push ${nginxRepositoryName}:${imageTag}

docker build --no-cache --rm -t ${goLangRepositoryName}:latest -f docker/go/Dockerfile .
docker tag ${goLangRepositoryName}:latest ${goLangRepositoryName}:${imageTag}
docker push ${goLangRepositoryName}:latest
docker push ${goLangRepositoryName}:${imageTag}

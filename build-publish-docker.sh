#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR

docker build -f WatchSecret.Dockerfile -t bmaynard/kubevol-watch-secret:latest .
docker build -f WatchConfigMap.Dockerfile -t bmaynard/kubevol-watch-configmap:latest .

docker push bmaynard/kubevol-watch-secret:latest
docker push bmaynard/kubevol-watch-configmap:latest
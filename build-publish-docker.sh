#!/bin/bash

if [ -z "$TAG" ]
then
      echo "\$TAG has not been supplied"
      exit 1
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR

docker build -t bmaynard/kubevol-watch:$TAG .
docker push bmaynard/kubevol-watch:$TAG
#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

kubectl delete ns kubevol-test-run
kubectl create ns kubevol-test-run

kubectl apply -n kubevol-test-run -f ${DIR}/yaml

while [[ $(kubectl -n kubevol-test-run get pods -l app=redis -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do echo "waiting for pod" && sleep 1; done

sleep 2

kubectl delete -n kubevol-test-run -f ${DIR}/yaml/redis-configmap-outdated.yml
kubectl apply -n kubevol-test-run -f ${DIR}/yaml/redis-configmap-outdated.yml


kubectl delete -n kubevol-test-run -f ${DIR}/yaml/redis-secret-outdated.yml
kubectl apply -n kubevol-test-run -f ${DIR}/yaml/redis-secret-outdated.yml
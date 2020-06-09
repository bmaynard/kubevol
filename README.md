# kubevol

This is a simple application that queries all pods for an attached volume or see all the volumes attached to each pod by specific type (eg: ConfigMap, Secret).

Features:

- Query for ConfigMaps and Secrets (future support coming for other types of volumes)
- Filter by namespace
- Filter by a specific object name
- See if attached volume is outdated
    - Limited support, can only detect if configmap was deleted after pod was created

## Install

Currently you need to build the binary yourself which you can accomplish with the following steps:

```
git clone git@github.com:bmaynard/kubevol.git
cd kubevol
go build
./kubevol --help
```

### Configuration

If your kubeconfig is not in the default location in your home directory, you can specify a custom kubeconfig file by creating the following file:

`~/.kubevol.yaml`
```
---
kubeconfig: /path/to/kube/config
```

## Sample Output

```
There are 1 pods in the cluster
Searching for pods that have a Secret attached

+------------------+----------+-----------------------+-----------------------+-------------+
| NAMESPACE        | POD NAME | SECRET NAME           | VOLUME NAME           | OUT OF DATE |
+------------------+----------+-----------------------+-----------------------+-------------+
| kubevol-test-run | redis    | redis-secret          | redis-secret          | Unknown     |
| kubevol-test-run | redis    | redis-secret-outdated | redis-secret-outdated | Yes         |
| kubevol-test-run | redis    | default-token-nd4wr   | default-token-nd4wr   | Unknown     |
+------------------+----------+-----------------------+-----------------------+-------------+
```

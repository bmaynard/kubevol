# kubevol

This is an application to audit all your Kubernetes pods for an attached volume or see all the volumes attached to each pod by specific type (eg: ConfigMap, Secret).

Features:

- Query for ConfigMaps and Secrets (future support coming for other types of volumes)
- Kubernetes controller to watch and record changes to ConfigMaps and Secrets
- Filter by namespace
- Filter by a specific object name
- See if attached volume has a stale version attached

## Installation

You can download the latest release from [Releases](https://github.com/bmaynard/kubevol/releases).

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

# kubevol

![Unit Tests](https://github.com/bmaynard/kubevol/workflows/Unit%20Tests/badge.svg) ![Build](https://github.com/bmaynard/kubevol/workflows/Build/badge.svg) ![Release](https://github.com/bmaynard/kubevol/workflows/Release/badge.svg)

Kubevol allows you to audit all your Kubernetes pods for an attached volume or see all the volumes attached to each pod by a specific type (eg: ConfigMap, Secret).

Features:

- Query for ConfigMaps and Secrets (future support coming for other types of volumes)
- Kubernetes controller to watch and record changes to ConfigMaps and Secrets
- Filter by namespace
- Filter by a specific object name
- See if attached volume has a stale version attached

## Installation

You can download the latest release from [Releases](https://github.com/bmaynard/kubevol/releases).

## Watch And Record Changes

Since Kubernetes doesn't keep track of when a `Secret` or `Configmap` was updated, `kubevol` has a Kubernetes controller that will watch for all changes and will record the last modified date. This then gives `kubevol` the ability to detect if an attached `Secret` or `Configmap` is outdated. 

To install the watch controller, run:

```bash
$ kubectl apply -f https://raw.githubusercontent.com/bmaynard/kubevol/master/deployment/manifest.yaml
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
$ kubevol secret
There are 12 pods in the cluster
Searching for pods that have a Secret attached

+------------------+----------+-----------------------+-----------------------+-------------+
| NAMESPACE        | POD NAME | SECRET NAME           | VOLUME NAME           | OUT OF DATE |
+------------------+----------+-----------------------+-----------------------+-------------+
| kubevol-test-run | redis    | redis-secret          | redis-secret          | No          |
| kubevol-test-run | redis    | redis-secret-outdated | redis-secret-outdated | Yes         |
+------------------+----------+-----------------------+-----------------------+-------------+
```

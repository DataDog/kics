---
title: "Encryption provider config is not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/encryption_provider_config_is_not_defined"
  id: "cbd2db69-0b21-4c14-8a40-7710a50571a9"
  display_name: "Encryption provider config is not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `cbd2db69-0b21-4c14-8a40-7710a50571a9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Encryption

### Description

 When using `kube-apiserver`, the `--encryption-provider-config` flag should be set, and encryption should be correctly configured in the encryption configuration file.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: command-demo
  labels:
    purpose: demonstrate-command
spec:
  containers:
    - name: command-demo-container
      image: gcr.io/google_containers/kube-apiserver-amd64:v1.6.0
      command: ["kube-apiserver"]
      args: ["--encryption-provider-config=/path/to/config/file.yaml"]
  restartPolicy: OnFailure

```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: command-demo
  labels:
    purpose: demonstrate-command
spec:
  containers:
    - name: command-demo-container
      image: gcr.io/google_containers/kube-apiserver-amd64:v1.6.0
      command: ["kube-apiserver","--encryption-provider-config=/path/to/config/file.yaml"]
      args: []
  restartPolicy: OnFailure

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: command-demo
  labels:
    purpose: demonstrate-command
spec:
  containers:
    - name: command-demo-container
      image: gcr.io/google_containers/kube-apiserver-amd64:v1.6.0
      command: ["kube-apiserver"]
      args: []
  restartPolicy: OnFailure

```
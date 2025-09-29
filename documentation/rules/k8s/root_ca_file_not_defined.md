---
title: "Root CA file not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/root_ca_file_not_defined"
  id: "05fb986f-ac73-4ebb-a5b2-7faafa93d882"
  display_name: "Root CA file not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `05fb986f-ac73-4ebb-a5b2-7faafa93d882`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Encryption

### Description

 When using `kube-controller-manager`, the `--root-ca-file` flag should be set.


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
      image: gcr.io/google_containers/kube-controller-manager-amd64:v1.6.0
      command: ["kube-controller-manager"]
      args: ["--root-ca-file=/path/to/ca/file.pem"]
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
      image: gcr.io/google_containers/kube-controller-manager-amd64:v1.6.0
      command: ["kube-controller-manager","--root-ca-file=/path/to/ca/file.pem"]
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
      image: gcr.io/google_containers/kube-controller-manager-amd64:v1.6.0
      command: ["kube-controller-manager"]
      args: []
  restartPolicy: OnFailure

```
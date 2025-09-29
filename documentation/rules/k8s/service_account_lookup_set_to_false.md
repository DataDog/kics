---
title: "Service account lookup set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_lookup_set_to_false"
  id: "a5530bd7-225a-48f9-91bb-f40b04200165"
  display_name: "Service account lookup set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `a5530bd7-225a-48f9-91bb-f40b04200165`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--service-account-lookup` flag should be set to `true`.


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
      args: ["--service-account-lookup=true"]
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
      command: ["kube-apiserver"]
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
      args: ["--service-account-lookup=false"]
  restartPolicy: OnFailure

```
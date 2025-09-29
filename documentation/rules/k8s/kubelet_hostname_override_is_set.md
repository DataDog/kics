---
title: "Kubelet hostname override is set"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_hostname_override_is_set"
  id: "bf36b900-b5ef-4828-adb7-70eb543b7cfb"
  display_name: "Kubelet hostname override is set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `bf36b900-b5ef-4828-adb7-70eb543b7cfb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 Hostnames should not be overridden.


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
      image: foo/bar
      command: ["kubelet"]
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
      image: foo/bar
      command: ["kubelet","--hostname-override=host"]
      args: []
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
      image: foo/bar
      command: ["kubelet"]
      args: ["--hostname-override=host"]
  restartPolicy: OnFailure

```
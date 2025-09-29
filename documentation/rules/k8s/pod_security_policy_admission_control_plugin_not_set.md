---
title: "Pod security policy admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/pod_security_policy_admission_control_plugin_not_set"
  id: "afa36afb-39fe-4d94-b9b6-afb236f7a03d"
  display_name: "Pod security policy admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Build Process"
---
## Metadata

**Id:** `afa36afb-39fe-4d94-b9b6-afb236f7a03d`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Build Process

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `PodSecurityPolicy`, and it should be correctly configured in the admission control config file.


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
      args: ["--enable-admission-plugins=PodSecurityPolicy", "--admission-control-config-file=path/to/plugin/config/file.yaml"]
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
      command: ["kube-apiserver","--enable-admission-plugins=PodSecurityPolicy", "--admission-control-config-file=path/to/plugin/config/file.yaml"]
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
      args: ["--enable-admission-plugins=AlwaysAdmit"]
  restartPolicy: OnFailure

```
---
title: "Container is privileged"
group_id: "rules/k8s"
meta:
  name: "k8s/container_is_privileged"
  id: "dd29336b-fe57-445b-a26e-e6aa867ae609"
  display_name: "Container is privileged"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `dd29336b-fe57-445b-a26e-e6aa867ae609`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/concepts/workloads/pods/#privileged-mode-for-containers)

### Description

 Privileged containers lack essential security restrictions and should be avoided by removing the `privileged` flag or setting it to `false`.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-4
spec:
  containers:
  - name: sec-ctx-4
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      privileged: false
      capabilities:
        add: ["NET_ADMIN", "SYS_TIME"]

```
## Non-Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deployment
  labels:
    app: test
spec:
  replicas: 3
  selector:
    matchLabels:
      app: test
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
        - name:  pause
          image: k8s.gcr.io/pause
          securityContext:
            privileged: true

```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-4
spec:
  containers:
  - name: sec-ctx-4
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      privileged: true
      capabilities:
        add: ["NET_ADMIN", "SYS_TIME"]
---
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-5
spec:
  initContainers:
  - name: sec-ctx-4
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      privileged: true
      capabilities:
        add: ["NET_ADMIN", "SYS_TIME"]
  containers:
  - name: sec-ctx-4
    image: gcr.io/google-samples/node-hello:1.0

```
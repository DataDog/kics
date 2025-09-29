---
title: "Auto TLS set to true"
group_id: "rules/k8s"
meta:
  name: "k8s/auto_tls_set_to_true"
  id: "98ce8b81-7707-4734-aa39-627c6db3d84b"
  display_name: "Auto TLS set to true"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `98ce8b81-7707-4734-aa39-627c6db3d84b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `etcd`, the `--auto-tls` flag should be set to `false`.


## Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-etcd-deployment
spec:
  selector:
    matchLabels:
      app: app
  replicas: 1
  template:
    metadata:
      labels:
        app: app
        version: v1
    spec:
      serviceAccountName: database
      containers:
      - name: database
        image: gcr.io/google_containers/etcd:v3.2.18
        imagePullPolicy: IfNotPresent
        command: ["etcd"]
        args: []
      nodeSelector:
        kubernetes.io/hostname: worker02  
    restartPolicy: OnFailure

```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-etcd-deployment
spec:
  selector:
    matchLabels:
      app: app
  replicas: 1
  template:
    metadata:
      labels:
        app: app
        version: v1
    spec:
      serviceAccountName: database
      containers:
      - name: database
        image: gcr.io/google_containers/etcd:v3.2.18
        imagePullPolicy: IfNotPresent
        command: ["etcd", "--auto-tls=false"]
        args: []
      nodeSelector:
        kubernetes.io/hostname: worker02  
    restartPolicy: OnFailure

```
## Non-Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-etcd-deployment
spec:
  selector:
    matchLabels:
      app: app
  replicas: 1
  template:
    metadata:
      labels:
        app: app
        version: v1
    spec:
      serviceAccountName: database
      containers:
      - name: database
        image: gcr.io/google_containers/etcd:v3.2.18
        imagePullPolicy: IfNotPresent
        command: ["etcd"]
        args: ["--auto-tls=true"]
      nodeSelector:
        kubernetes.io/hostname: worker02  
    restartPolicy: OnFailure

```
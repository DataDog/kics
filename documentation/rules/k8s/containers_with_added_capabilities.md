---
title: "Containers with added capabilities"
group_id: "rules/k8s"
meta:
  name: "k8s/containers_with_added_capabilities"
  id: "19ebaa28-fc86-4a58-bcfa-015c9e22fe40"
  display_name: "Containers with added capabilities"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `19ebaa28-fc86-4a58-bcfa-015c9e22fe40`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Containers should not allow extra capabilities.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod1
spec:
  containers:
  - name: app
    image: images.my-company.example/app:v4
    securityContext:
      allowPrivilegeEscalation: false
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"

  - name: log-aggregator
    image: images.my-company.example/log-aggregator:v6
    securityContext:
      allowPrivilegeEscalation: false
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"

```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod4
spec:
  containers:
  - name: app
    image: images.my-company.example/app:v4
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
        drop:
          - ALL
        add:
          - NET_BIND_SERVICE
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: log-aggregator
    image: images.my-company.example/log-aggregator:v6
    securityContext:
      allowPrivilegeEscalation: false
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod2
spec:
  containers:
  - name: app
    image: images.my-company.example/app:v4
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
           add: ["NET_ADMIN", "SYS_TIME"]
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: log-aggregator
    image: images.my-company.example/log-aggregator:v6
    securityContext:
      allowPrivilegeEscalation: false
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
---
apiVersion: v1
kind: Pod
metadata:
  name: pod3
spec:
  initContainers:
  - name: app
    image: images.my-company.example/app:v4
    securityContext:
      allowPrivilegeEscalation: false
      capabilities:
           add: ["NET_ADMIN", "SYS_TIME"]
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: log-aggregator
    image: images.my-company.example/log-aggregator:v6
    securityContext:
      allowPrivilegeEscalation: false
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  containers:
  - name: app
    image: images.my-company.example/app:v4

```
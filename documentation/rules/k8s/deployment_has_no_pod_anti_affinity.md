---
title: "Deployment without podAntiAffinity"
group_id: "rules/k8s"
meta:
  name: "k8s/deployment_has_no_pod_anti_affinity"
  id: "a31b7b82-d994-48c4-bd21-3bab6c31827a"
  display_name: "Deployment without podAntiAffinity"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Id:** `a31b7b82-d994-48c4-bd21-3bab6c31827a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Resource Management

### Description

 Deployments should define a `podAntiAffinity` policy to prevent multiple Pods from being scheduled on the same node.


## Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-server
spec:
  selector:
    matchLabels:
      app: web-store
  replicas: 3
  template:
    metadata:
      labels:
        app: web-store
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - web-store
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: web-app
        image: nginx:1.16-alpine
```
## Non-Compliant Code Examples
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: label-mismatch
spec:
  selector:
    matchLabels:
      app: web-store
  replicas: 3
  template:
    metadata:
      labels:
        app: web-shore
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchLabels:
                app: web-store
            topologyKey: "kubernetes.io/hostname"
      containers:
      - name: web-app
        image: nginx:1.16-alpine
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: no-affinity
spec:
  selector:
    matchLabels:
      app: web-store
  replicas: 3
  template:
    metadata:
      labels:
        app: web-store
    spec:
      containers:
      - name: web-app
        image: nginx:1.16-alpine

```
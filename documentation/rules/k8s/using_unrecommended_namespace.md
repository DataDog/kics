---
title: "Using unrecommended namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/using_unrecommended_namespace"
  id: "611ab018-c4aa-4ba2-b0f6-a448337509a6"
  display_name: "Using unrecommended namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `611ab018-c4aa-4ba2-b0f6-a448337509a6`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Namespaces such as `default`, `kube-system`, or `kube-public` should not be used.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
  namespace: cosmicPod
spec:
  securityContext:
    runAsUser: 1000
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

---
apiVersion: v1
kind: CustomResourceDefinition
metadata:
  name: mongo.db.collection.com

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend2
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
  name: mongo.db.collection.com
  namespace: kube-public

```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mongo.db.collection.com
  namespace: kube-system

```
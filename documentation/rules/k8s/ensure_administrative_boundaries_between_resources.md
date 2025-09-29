---
title: "Ensure administrative boundaries between resources"
group_id: "rules/k8s"
meta:
  name: "k8s/ensure_administrative_boundaries_between_resources"
  id: "e84eaf4d-2f45-47b2-abe8-e581b06deb66"
  display_name: "Ensure administrative boundaries between resources"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Access Control"
---
## Metadata

**Id:** `e84eaf4d-2f45-47b2-abe8-e581b06deb66`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Access Control

### Description

 As a best practice, ensure correct use of namespaces to administer resources. Kubernetes authorization plugins can also create policies that segregate user access to namespaces.

## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
  namespace: cosmic-namespace
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
---
apiVersion: v1
kind: Pod
metadata:
  name: frontend2
  namespace: cosmic-namespace
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
---
title: "Pod or container without security context"
group_id: "rules/k8s"
meta:
  name: "k8s/pod_or_container_without_security_context"
  id: "a97a340a-0063-418e-b3a1-3028941d0995"
  display_name: "Pod or container without security context"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `a97a340a-0063-418e-b3a1-3028941d0995`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 A security context defines privilege and access control settings for a Pod or container.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
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


```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: frontend
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
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
```
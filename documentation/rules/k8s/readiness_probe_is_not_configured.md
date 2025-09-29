---
title: "Readiness probe is not configured"
group_id: "rules/k8s"
meta:
  name: "k8s/readiness_probe_is_not_configured"
  id: "a659f3b5-9bf0-438a-bd9a-7d3a6427f1e3"
  display_name: "Readiness probe is not configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata

**Id:** `a659f3b5-9bf0-438a-bd9a-7d3a6427f1e3`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Availability

### Description

 A readiness probe should be configured.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: goproxy
  labels:
    app: goproxy
spec:
  containers:
  - name: goproxy
    image: k8s.gcr.io/goproxy:0.1
    ports:
    - containerPort: 8080
    readinessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 5
      periodSeconds: 10
    livenessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 15
      periodSeconds: 20

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: goproxy
  labels:
    app: goproxy
spec:
  containers:
  - name: goproxy
    image: k8s.gcr.io/goproxy:0.1
    ports:
    - containerPort: 8080
    livenessProbe:
      tcpSocket:
        port: 8080
      initialDelaySeconds: 15
      periodSeconds: 20

```
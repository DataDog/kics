---
title: "Invalid metadata label"
group_id: "rules/k8s"
meta:
  name: "k8s/metadata_label_is_invalid"
  id: "1123031a-f921-4c5b-bd86-ef354ecfd37a"
  display_name: "Invalid metadata label"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata

**Id:** `1123031a-f921-4c5b-bd86-ef354ecfd37a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)

### Description

 Metadata labels should be valid.


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
    app: g**dy.l+bel.
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
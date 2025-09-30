---
title: "Container running as root"
group_id: "rules/k8s"
meta:
  name: "k8s/containers_running_as_root"
  id: "cf34805e-3872-4c08-bf92-6ff7bb0cfadb"
  display_name: "Container running as root"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata

**Id:** `cf34805e-3872-4c08-bf92-6ff7bb0cfadb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)

### Description

 Containers should run only as a non-root user. This limits the exploitability of security misconfigurations and restricts an attacker's options in case of compromise.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-2
spec:
  securityContext:
    runAsUser: 10000
    runAsNonRoot: true
  containers:
  - name: sec-ctx-demo-2
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 10100
      allowPrivilegeEscalation: false
      runAsNonRoot: true
      
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: containers-runs-as-root
spec:
  securityContext:
    runAsUser: 0
    runAsNonRoot: false
  containers:
  - name: sec-ctx-demo-100
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 1000
      runAsNonRoot: false
      
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-1
spec:
  securityContext:
    runAsUser: 1000
    runAsNonRoot: true
  containers:
  - name: sec-ctx-demo-100
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 1000
      runAsNonRoot: false
  - name: sec-ctx-demo-200
    image: gcr.io/google-samples/node-hedwfwllo:1.0
    securityContext:
      runAsUser: 2000
      runAsNonRoot: true
      
```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-2
spec:
  securityContext:
    runAsUser: 1000
    runAsNonRoot: false
  containers:
  - name: sec-ctx-demo-2
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 0
      allowPrivilegeEscalation: false
      runAsNonRoot: false
---
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-3
spec:
  securityContext:
    runAsUser: 1000
    runAsNonRoot: false
  containers:
  - name: sec-ctx-demo-2
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      allowPrivilegeEscalation: false
      runAsNonRoot: false
---
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-4
spec:
  securityContext:
    runAsUser: 1000
    runAsNonRoot: true
  containers:
  - name: sec-ctx-demo-2
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 0
      allowPrivilegeEscalation: false
      runAsNonRoot: false

```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: containers-runs-as-root
spec:
  securityContext:
    runAsUser: 0
    runAsNonRoot: false
  containers:
  - name: sec-ctx-demo-100
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 0
      runAsNonRoot: false
      
      
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-2
spec:
  securityContext:
    runAsUser: 10
    runAsNonRoot: false
  containers:
  - name: sec-ctx-demo-100
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      runAsUser: 0
      runAsNonRoot: false
  - name: sec-ctx-demo-200
    image: gcr.io/google-samples/node-hedwfwllo:1.0
    securityContext:
      runAsUser: 0
      runAsNonRoot: false
      
```
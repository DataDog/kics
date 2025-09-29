---
title: "Ingress controller exposes workload"
group_id: "rules/k8s"
meta:
  name: "k8s/ingress_controller_exposes_workload"
  id: "69bbc5e3-0818-4150-89cc-1e989b48f23b"
  display_name: "Ingress controller exposes workload"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `69bbc5e3-0818-4150-89cc-1e989b48f23b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Ingress controllers should not expose workloads to avoid vulnerabilities and DoS attacks.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Service
metadata:
  name: app
  labels:
    app: app
spec:
  type: ClusterIP
  ports:
  - port: 3000
    targetPort: 3000
  selector:
    app: app


---

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: app-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
  labels:
    app: app
spec:
  rules:
  - host: app.acme.org
    http:
      paths:
      - backend:
          serviceName: app2
          servicePort: 3000

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Service
metadata:
  name: app
  labels:
    app: app
spec:
  type: ClusterIP
  ports:
  - port: 3000
    targetPort: 3000
  selector:
    app: app


---

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: app-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
  labels:
    app: app
spec:
  rules:
  - host: app.acme.org
    http:
      paths:
      - backend:
          serviceName: app
          servicePort: 3000

```
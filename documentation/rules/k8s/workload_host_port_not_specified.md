---
title: "Workload host port not specified"
group_id: "rules/k8s"
meta:
  name: "k8s/workload_host_port_not_specified"
  id: "2b1836f1-dcce-416e-8e16-da8c71920633"
  display_name: "Workload host port not specified"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `2b1836f1-dcce-416e-8e16-da8c71920633`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Networking and Firewall

### Description

 Verifies whether the Kubernetes workload's host port is specified.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: firstpod
spec:
  containers:
  - name: container
    image: nginx
    ports:
    - containerPort: 80
      hostIP: 10.0.0.1
```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: firstpod
spec:
  containers:
  - name: container
    image: nginx
    ports:
    - containerPort: 80
      hostIP: 10.0.0.1
      hostPort: 8080
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: secondpod
spec:
  template:
    spec:
      containers:
      - name: container2
        image: nginx
        ports:
        - containerPort: 81
          hostIP: 10.0.0.2
          hostPort: 8081
    metadata:
      labels:
        app: nginx
  selector:
    matchLabels:
      app: nginx

```
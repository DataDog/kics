---
title: "StatefulSet without service name"
group_id: "rules/k8s"
meta:
  name: "k8s/statefulset_without_service_name"
  id: "bb241e61-77c3-4b97-9575-c0f8a1e008d0"
  display_name: "StatefulSet without service name"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Availability"
---
## Metadata

**Id:** `bb241e61-77c3-4b97-9575-c0f8a1e008d0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Availability

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)

### Description

 StatefulSets should have an existing headless `serviceName`. The headless service labels should also be applied to StatefulSet labels.


## Compliant Code Examples
```yaml
#this is a problematic code where the query should report a result(s)
apiVersion: v1
kind: Service
metadata:
  name: nginx2
  namespace: nginx2
  labels:
    app: nginx2
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx2
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web2
  namespace: nginx2
spec:
  selector:
    matchLabels:
      app: nginx2
  serviceName: "nginx2"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx2
        foo: bar
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx2
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi

```
## Non-Compliant Code Examples
```yaml
#this is a problematic code where the query should report a result(s)
apiVersion: v1
kind: Service
metadata:
  name: nginx
  namespace: nginx
  labels:
    app: nginx2
spec:
  ports:
  - port: 80
    name: web
  clusterIP: All
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
  namespace: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx"
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      terminationGracePeriodSeconds: 10
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "my-storage-class"
      resources:
        requests:
          storage: 1Gi

```
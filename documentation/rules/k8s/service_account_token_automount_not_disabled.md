---
title: "Service account token auto-mount not disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_token_automount_not_disabled"
  id: "48471392-d4d0-47c0-b135-cdec95eb3eef"
  display_name: "Service account token auto-mount not disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---
## Metadata

**Id:** `48471392-d4d0-47c0-b135-cdec95eb3eef`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Defaults

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server)

### Description

 Service account tokens are automatically mounted even if not necessary.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  automountServiceAccountToken: false
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
  volumes:
    - name: sec-ctx-vol
      emptyDir: { }
  containers:
    - name: sec-ctx-demo
      image: busybox
      command: [ "sh", "-c", "sleep 1h" ]
      volumeMounts:
        - name: sec-ctx-vol
          mountPath: /data/demo
      securityContext:
        allowPrivilegeEscalation: false
```

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: redistest-sa
automountServiceAccountToken: false
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demoenv
  labels:
    app: redis
spec:
  selector:
    matchLabels:
      app: redis
  template:
    metadata:      
      labels:
        app: redis
    spec:
      serviceAccountName: redistest-sa
      containers:
      - name: redis
        image: redis:latest

```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
  volumes:
    - name: sec-ctx-vol
      emptyDir: { }
  containers:
    - name: sec-ctx-demo
      image: busybox
      command: [ "sh", "-c", "sleep 1h" ]
      volumeMounts:
        - name: sec-ctx-vol
          mountPath: /data/demo
      securityContext:
        allowPrivilegeEscalation: false
---
apiVersion: v1
kind: Pod
metadata:
  name: security.context.demo
spec:
  automountServiceAccountToken: true
  securityContext:
    runAsUser: 1000
    runAsGroup: 3000
    fsGroup: 2000
  volumes:
    - name: sec-ctx-vol
      emptyDir: { }
  containers:
    - name: sec-ctx-demo
      image: busybox
      command: [ "sh", "-c", "sleep 1h" ]
      volumeMounts:
        - name: sec-ctx-vol
          mountPath: /data/demo
      securityContext:
        allowPrivilegeEscalation: false
---
apiVersion: serving.knative.dev/v1
kind: Configuration
metadata:
  name: dummy-config
  namespace: knative-sequence
spec:
  template:
    spec:
      automountServiceAccountToken: true
      securityContext:
        runAsUser: 1000
        runAsGroup: 3000
        fsGroup: 2000
      volumes:
        - name: sec-ctx-vol
          emptyDir: { }
      containers:
        - name: sec-ctx-demo
          image: busybox
          command: [ "sh", "-c", "sleep 1h" ]
          volumeMounts:
            - name: sec-ctx-vol
              mountPath: /data/demo
          securityContext:
            allowPrivilegeEscalation: false        

```

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: redistest-sa
automountServiceAccountToken: true
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demoenv
  labels:
    app: redis
spec:
  selector:
    matchLabels:
      app: redis
  template:
    metadata:      
      labels:
        app: redis
    spec:
      serviceAccountName: redistest-sa
      containers:
      - name: redis
        image: redis:latest
```
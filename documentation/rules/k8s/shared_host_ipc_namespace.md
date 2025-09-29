---
title: "Shared host IPC namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/shared_host_ipc_namespace"
  id: "cd290efd-6c82-4e9d-a698-be12ae31d536"
  display_name: "Shared host IPC namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata

**Id:** `cd290efd-6c82-4e9d-a698-be12ae31d536`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Resource Management

### Description

 Containers should not share the host IPC namespace.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  hostIPC: false
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
apiVersion: serving.knative.dev/v1
kind: Configuration
metadata:
  name: dummy-config
  namespace: knative-sequence
spec:
  template:
    spec:
      hostIPC: false      
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
## Non-Compliant Code Examples
```yaml
apiVersion: serving.knative.dev/v1
kind: Configuration
metadata:
  name: dummy-config
  namespace: knative-sequence
spec:
  template:
    spec:
      hostIPC: true
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
kind: Pod
metadata:
  name: security-context-demo
spec:
  hostIPC: true
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
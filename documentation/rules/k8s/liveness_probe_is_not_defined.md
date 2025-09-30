---
title: "Liveness probe is not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/liveness_probe_is_not_defined"
  id: "ade74944-a674-4e00-859e-c6eab5bde441"
  display_name: "Liveness probe is not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Availability"
---
## Metadata

**Id:** `ade74944-a674-4e00-859e-c6eab5bde441`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Availability

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#when-should-you-use-a-liveness-probe)

### Description

 A liveness probe can improve availability by restarting unresponsive containers; however, it can cause cascading failures. Define one only when necessary.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-exec
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5

---

apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox
            imagePullPolicy: IfNotPresent
            args:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-exec
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600

```
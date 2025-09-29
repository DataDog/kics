---
title: "Kubelet client periodic certificate switch disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_client_periodic_certificate_switch_disabled"
  id: "52d70f2e-3257-474c-b3dc-8ad9ba6a061a"
  display_name: "Kubelet client periodic certificate switch disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `52d70f2e-3257-474c-b3dc-8ad9ba6a061a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 The kubelet argument `--rotate-certificates` should be set to `true`.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: command-demo
  labels:
    purpose: demonstrate-command
spec:
  containers:
    - name: command-demo-container
      image: foo/bar
      command: ["kubelet"]
      args: ["--rotate-certificates"]
  restartPolicy: OnFailure

```

```yaml
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
address: "192.168.0.8"
port: 20250
serializeImagePulls: false
evictionHard:
    memory.available:  "200Mi"
rotateCertificates: true

```

```json
{
    "port": 20250,
    "evictionHard": {
        "memory.available": "200Mi"
    },
    "kind": "KubeletConfiguration",
    "rotateCertificates": true,
    "serializeImagePulls": false,
    "address": "192.168.0.8",
    "apiVersion": "kubelet.config.k8s.io/v1beta1"
}

```
## Non-Compliant Code Examples
```yaml
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
address: "192.168.0.8"
port: 20250
serializeImagePulls: false
evictionHard:
    memory.available:  "200Mi"
rotateCertificates: false

```

```json
{
    "port": 20250,
    "evictionHard": {
        "memory.available": "200Mi"
    },
    "kind": "KubeletConfiguration",
    "makeIPTablesUtilChains": true,
    "address": "192.168.0.8",
    "apiVersion": "kubelet.config.k8s.io/v1beta1"
}

```

```yaml
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
address: "192.168.0.8"
port: 20250
serializeImagePulls: false
evictionHard:
    memory.available:  "200Mi"

```
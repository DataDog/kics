---
title: "TLS connection certificate not set up"
group_id: "rules/k8s"
meta:
  name: "k8s/tls_connection_certificate_not_setup"
  id: "fa750c81-93c2-4fab-9c6d-d3fd3ce3b89f"
  display_name: "TLS connection certificate not set up"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `fa750c81-93c2-4fab-9c6d-d3fd3ce3b89f`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 TLS connection certificate files should be set up.


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
      image: gcr.io/google_containers/kube-apiserver-amd64:v1.6.0
      command: ["kube-apiserver"]
      args: ["--tls-cert-file=someFile.txt","--tls-private-key-file=someFile.txt"]
  restartPolicy: OnFailure

```

```json
{
  "kind": "KubeletConfiguration",
  "apiVersion": "kubelet.config.k8s.io/v1beta1",
  "port": 10250,
  "readOnlyPort": 10255,
  "cgroupDriver": "cgroupfs",
  "hairpinMode": "promiscuous-bridge",
  "serializeImagePulls": false,
  "tlsCertFile": "someFile.txt",
  "tlsPrivateKeyFile": "someFile.txt",
  "featureGates": {
    "RotateKubeletClientCertificate": true,
    "RotateKubeletServerCertificate": true
  }
}

```

```yaml
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
address: "192.168.0.8"
port: 20250
serializeImagePulls: false
tlsCertFile: "someFile.txt"
tlsPrivateKeyFile: "someFile.txt"
evictionHard:
    memory.available:  "200Mi"



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



```

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
      image: gcr.io/google_containers/kube-apiserver-amd64:v1.6.0
      command: ["kube-apiserver"]
      args: []
  restartPolicy: OnFailure




```
---
title: "Kubelet streaming connection timeout disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_streaming_connection_timeout_disabled"
  id: "ed89b97d-04e9-4fd4-919f-ee5b27e555e9"
  display_name: "Kubelet streaming connection timeout disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `ed89b97d-04e9-4fd4-919f-ee5b27e555e9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 The `--streaming-connection-idle-timeout` flag should not be set to `0`.

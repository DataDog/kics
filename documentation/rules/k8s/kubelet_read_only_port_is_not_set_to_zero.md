---
title: "Kubelet read-only port is not set to zero"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_read_only_port_is_not_set_to_zero"
  id: "2940d48a-dc5e-4178-a3f8-bfbd80720b41"
  display_name: "Kubelet read-only port is not set to zero"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `2940d48a-dc5e-4178-a3f8-bfbd80720b41`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `kubelet`, the read-only port should be set to `0` (`--read-only-port=0`).

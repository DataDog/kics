---
title: "Secure port set to zero"
group_id: "rules/k8s"
meta:
  name: "k8s/secure_port_set_to_zero"
  id: "3d24b204-b73d-42cb-b0bf-1a5438c5f71e"
  display_name: "Secure port set to zero"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `3d24b204-b73d-42cb-b0bf-1a5438c5f71e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Networking and Firewall

### Description

 When using `kube-apiserver`, the `--secure-port` flag should not be set to `0`.

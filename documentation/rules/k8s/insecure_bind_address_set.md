---
title: "Insecure bind address set"
group_id: "rules/k8s"
meta:
  name: "k8s/insecure_bind_address_set"
  id: "b9380fd3-5ffe-4d10-9290-13e18e71eee1"
  display_name: "Insecure bind address set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `b9380fd3-5ffe-4d10-9290-13e18e71eee1`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Networking and Firewall

### Description

 When using `kube-apiserver`, the `--insecure-bind-address` flag should not be set.

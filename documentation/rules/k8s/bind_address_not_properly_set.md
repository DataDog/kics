---
title: "Bind address not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/bind_address_not_properly_set"
  id: "46a2e9ec-6a5f-4faa-9d39-4ea44d5d87a2"
  display_name: "Bind address not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `46a2e9ec-6a5f-4faa-9d39-4ea44d5d87a2`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Networking and Firewall

### Description

 When using `kube-controller-manager` or `kube-scheduler`, the `--bind-address` should be set to `127.0.0.1`.

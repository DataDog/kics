---
title: "CNI plugin does not support network policies"
group_id: "rules/k8s"
meta:
  name: "k8s/cni_plugin_does_not_support_network_policies"
  id: "03aabc8c-35d6-481e-9c85-20139cf72d23"
  display_name: "CNI plugin does not support network policies"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `03aabc8c-35d6-481e-9c85-20139cf72d23`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 Ensure use of a CNI plugin that supports NetworkPolicies. Without support, it may not be possible to effectively restrict cluster traffic.

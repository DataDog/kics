---
title: "PodSecurityPolicy allows host network sharing"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_containers_share_host_network_namespace"
  id: "a33e9173-b674-4dfb-9d82-cf3754816e4b"
  display_name: "PodSecurityPolicy allows host network sharing"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `a33e9173-b674-4dfb-9d82-cf3754816e4b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 PodSecurityPolicies should not allow containers to share the host network namespace.

---
title: "Network policy without Pod target"
group_id: "rules/k8s"
meta:
  name: "k8s/network_policy_is_not_targeting_any_pod"
  id: "85ab1c5b-014e-4352-b5f8-d7dea3bb4fd3"
  display_name: "Network policy without Pod target"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `85ab1c5b-014e-4352-b5f8-d7dea3bb4fd3`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Networking and Firewall

### Description

 Each network policy should target at least one Pod.

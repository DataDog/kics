---
title: "PSP set to privileged"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_set_to_privileged"
  id: "c48e57d3-d642-4e0b-90db-37f807b41b91"
  display_name: "PSP set to privileged"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `c48e57d3-d642-4e0b-90db-37f807b41b91`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Pods should not be allowed to request execution as privileged.

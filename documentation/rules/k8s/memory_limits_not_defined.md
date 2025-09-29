---
title: "Memory limits not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/memory_limits_not_defined"
  id: "b14d1bc4-a208-45db-92f0-e21f8e2588e9"
  display_name: "Memory limits not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata

**Id:** `b14d1bc4-a208-45db-92f0-e21f8e2588e9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Resource Management

### Description

 Memory limits should be defined for each container to prevent resource exhaustion by ensuring that containers do not consume more than the designated amount of memory.

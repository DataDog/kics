---
title: "Pod or container without LimitRange"
group_id: "rules/k8s"
meta:
  name: "k8s/pod_or_container_without_limit_range"
  id: "4a20ebac-1060-4c81-95d1-1f7f620e983b"
  display_name: "Pod or container without LimitRange"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `4a20ebac-1060-4c81-95d1-1f7f620e983b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 Each namespace should have an associated LimitRange policy to ensure that resource allocations of Pods, containers, and PersistentVolumeClaims do not exceed defined boundaries.

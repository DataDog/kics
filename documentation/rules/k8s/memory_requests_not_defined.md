---
title: "Memory requests not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/memory_requests_not_defined"
  id: "229588ef-8fde-40c8-8756-f4f2b5825ded"
  display_name: "Memory requests not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata

**Id:** `229588ef-8fde-40c8-8756-f4f2b5825ded`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Resource Management

### Description

 Memory requests should be defined for each container to allow the kubelet to reserve the requested system resources and prevent over-provisioning on individual nodes.

---
title: "CPU requests not set"
group_id: "rules/k8s"
meta:
  name: "k8s/cpu_requests_not_set"
  id: "ca469dd4-c736-448f-8ac1-30a642705e0a"
  display_name: "CPU requests not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Id:** `ca469dd4-c736-448f-8ac1-30a642705e0a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Resource Management

### Description

 CPU requests should be set to ensure the sum of resource requests of scheduled containers is less than the node's capacity.

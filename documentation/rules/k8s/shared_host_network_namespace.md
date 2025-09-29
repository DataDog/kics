---
title: "Shared host network namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/shared_host_network_namespace"
  id: "6b6bdfb3-c3ae-44cb-88e4-7405c1ba2c8a"
  display_name: "Shared host network namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata

**Id:** `6b6bdfb3-c3ae-44cb-88e4-7405c1ba2c8a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Resource Management

### Description

 Containers should not share the host network namespace.

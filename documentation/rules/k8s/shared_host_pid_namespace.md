---
title: "Shared host PID namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/shared_host_pid_namespace"
  id: "302736f4-b16c-41b8-befe-c0baffa0bd9d"
  display_name: "Shared host PID namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `302736f4-b16c-41b8-befe-c0baffa0bd9d`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Containers should not share the host process ID namespace.

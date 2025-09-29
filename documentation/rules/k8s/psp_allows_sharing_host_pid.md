---
title: "PSP allows sharing host PID"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_allows_sharing_host_pid"
  id: "91dacd0e-d189-4a9c-8272-5999a3cc32d9"
  display_name: "PSP allows sharing host PID"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `91dacd0e-d189-4a9c-8272-5999a3cc32d9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 PodSecurityPolicy allows containers to share the host process ID namespace.

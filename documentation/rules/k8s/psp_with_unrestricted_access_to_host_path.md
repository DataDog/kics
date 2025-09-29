---
title: "PSP with unrestricted access to host path"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_with_unrestricted_access_to_host_path"
  id: "de4421f1-4e35-43b4-9783-737dd4e4a47e"
  display_name: "PSP with unrestricted access to host path"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Resource Management"
---
## Metadata

**Id:** `de4421f1-4e35-43b4-9783-737dd4e4a47e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Resource Management

### Description

 PodSecurityPolicy should set `readOnly` to `true` for every allowed host path.

---
title: "Containers with sys admin capabilities"
group_id: "rules/k8s"
meta:
  name: "k8s/containers_with_sys_admin_capabilities"
  id: "235236ee-ad78-4065-bd29-61b061f28ce0"
  display_name: "Containers with sys admin capabilities"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `235236ee-ad78-4065-bd29-61b061f28ce0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Containers should not have the `CAP_SYS_ADMIN` Linux capability.

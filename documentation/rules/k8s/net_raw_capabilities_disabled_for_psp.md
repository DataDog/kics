---
title: "NET_RAW capabilities disabled for PSP"
group_id: "rules/k8s"
meta:
  name: "k8s/net_raw_capabilities_disabled_for_psp"
  id: "2270987f-bb51-479f-b8be-3ca73e5ad648"
  display_name: "NET_RAW capabilities disabled for PSP"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `2270987f-bb51-479f-b8be-3ca73e5ad648`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Containers should drop `NET_RAW` or `ALL` capabilities.

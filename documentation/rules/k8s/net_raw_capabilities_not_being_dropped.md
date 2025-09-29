---
title: "NET_RAW capabilities not dropped"
group_id: "rules/k8s"
meta:
  name: "k8s/net_raw_capabilities_not_being_dropped"
  id: "dbbc6705-d541-43b0-b166-dd4be8208b54"
  display_name: "NET_RAW capabilities not dropped"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `dbbc6705-d541-43b0-b166-dd4be8208b54`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Containers should drop `ALL` or at least `NET_RAW` capabilities.

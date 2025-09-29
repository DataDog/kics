---
title: "Container with unmasked /proc access"
group_id: "rules/k8s"
meta:
  name: "k8s/container_runs_unmasked"
  id: "f922827f-aab6-447c-832a-e1ff63312bd3"
  display_name: "Container with unmasked /proc access"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `f922827f-aab6-447c-832a-e1ff63312bd3`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Containers should not have full (unmasked) access to the host's `/proc` filesystem, as this can expose sensitive information and allow kernel parameter changes.

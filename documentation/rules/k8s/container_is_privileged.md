---
title: "Container is privileged"
group_id: "rules/k8s"
meta:
  name: "k8s/container_is_privileged"
  id: "dd29336b-fe57-445b-a26e-e6aa867ae609"
  display_name: "Container is privileged"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `dd29336b-fe57-445b-a26e-e6aa867ae609`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Privileged containers lack essential security restrictions and should be avoided by removing the `privileged` flag or setting it to `false`.

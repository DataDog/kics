---
title: "Missing AppArmor profile"
group_id: "rules/k8s"
meta:
  name: "k8s/missing_app_armor_config"
  id: "8b36775e-183d-4d46-b0f7-96a6f34a723f"
  display_name: "Missing AppArmor profile"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Access Control"
---
## Metadata

**Id:** `8b36775e-183d-4d46-b0f7-96a6f34a723f`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Access Control

### Description

 Containers should be configured with an AppArmor profile to enforce fine-grained access control over low-level system resources.

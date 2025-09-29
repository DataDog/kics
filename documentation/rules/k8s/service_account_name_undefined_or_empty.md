---
title: "Service account name undefined or empty"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_name_undefined_or_empty"
  id: "591ade62-d6b0-4580-b1ae-209f80ba1cd9"
  display_name: "Service account name undefined or empty"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---
## Metadata

**Id:** `591ade62-d6b0-4580-b1ae-209f80ba1cd9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Defaults

### Description

 A Kubernetes Pod should have a service account defined to restrict Kubernetes API access; the `serviceAccountName` attribute should be set and not empty.

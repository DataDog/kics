---
title: "Service account token auto-mount not disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_token_automount_not_disabled"
  id: "48471392-d4d0-47c0-b135-cdec95eb3eef"
  display_name: "Service account token auto-mount not disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---
## Metadata

**Id:** `48471392-d4d0-47c0-b135-cdec95eb3eef`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Defaults

### Description

 Service account tokens are automatically mounted even if not necessary.

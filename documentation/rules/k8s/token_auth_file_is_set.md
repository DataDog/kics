---
title: "Token auth file is set"
group_id: "rules/k8s"
meta:
  name: "k8s/token_auth_file_is_set"
  id: "32ecd76e-7bbf-402e-bf48-8b9485749558"
  display_name: "Token auth file is set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `32ecd76e-7bbf-402e-bf48-8b9485749558`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--token-auth-file` flag should not be set.

---
title: "Basic auth file is set"
group_id: "rules/k8s"
meta:
  name: "k8s/basic_auth_file_is_set"
  id: "5da47109-f8d6-4585-9e2b-96a8958a12f5"
  display_name: "Basic auth file is set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `5da47109-f8d6-4585-9e2b-96a8958a12f5`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--basic-auth-file` flag should not be set.

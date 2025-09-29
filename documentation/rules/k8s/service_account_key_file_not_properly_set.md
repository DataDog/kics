---
title: "Service account key file not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_key_file_not_properly_set"
  id: "dab4ec72-ce2e-4732-b7c3-1757dcce01a1"
  display_name: "Service account key file not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `dab4ec72-ce2e-4732-b7c3-1757dcce01a1`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `kube-apiserver`, the `--service-account-key-file` flag should be set.

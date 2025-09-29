---
title: "Use service account credentials not set to true"
group_id: "rules/k8s"
meta:
  name: "k8s/use_service_account_credentials_not_set_to_true"
  id: "1acd93f1-5a37-45c0-aaac-82ece818be7d"
  display_name: "Use service account credentials not set to true"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `1acd93f1-5a37-45c0-aaac-82ece818be7d`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 When using `kube-controller-manager`, the `--use-service-account-credentials` flag should be set to `true`.

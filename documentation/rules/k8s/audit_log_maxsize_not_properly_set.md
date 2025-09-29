---
title: "Audit log maxsize not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_log_maxsize_not_properly_set"
  id: "35c0a471-f7c8-4993-aa2c-503a3c712a66"
  display_name: "Audit log maxsize not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `35c0a471-f7c8-4993-aa2c-503a3c712a66`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

### Description

 When using `kube-apiserver`, the `--audit-log-maxsize` flag should be set to 100 megabytes or more.

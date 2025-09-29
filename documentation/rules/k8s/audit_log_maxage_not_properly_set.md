---
title: "Audit log maxage not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_log_maxage_not_properly_set"
  id: "da9f3aa8-fbfb-472f-b5a1-576127944218"
  display_name: "Audit log maxage not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `da9f3aa8-fbfb-472f-b5a1-576127944218`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

### Description

 When using `kube-apiserver`, the `--audit-log-maxage` flag should be set to 30 days or more.

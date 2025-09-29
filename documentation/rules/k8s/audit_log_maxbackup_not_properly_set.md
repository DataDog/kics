---
title: "Audit log maxbackup not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_log_maxbackup_not_properly_set"
  id: "768aab52-2504-4a2f-a3e3-329d5a679848"
  display_name: "Audit log maxbackup not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `768aab52-2504-4a2f-a3e3-329d5a679848`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

### Description

 When using `kube-apiserver`, the `--audit-log-maxbackup` flag should be set to 10 files or more.

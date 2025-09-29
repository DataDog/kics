---
title: "Audit log path not set"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_log_path_not_set"
  id: "73e251f0-363d-4e53-86e2-0a93592437eb"
  display_name: "Audit log path not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Id:** `73e251f0-363d-4e53-86e2-0a93592437eb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Observability

### Description

 When using `kube-apiserver`, the `--audit-log-path` flag should be set.

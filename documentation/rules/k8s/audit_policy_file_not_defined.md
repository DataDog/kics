---
title: "Audit policy file not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_policy_file_not_defined"
  id: "13a49a2e-488e-4309-a7c0-d6b05577a5fb"
  display_name: "Audit policy file not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Id:** `13a49a2e-488e-4309-a7c0-d6b05577a5fb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Observability

### Description

 When using `kube-apiserver`, the `--audit-policy-file` flag should be set.

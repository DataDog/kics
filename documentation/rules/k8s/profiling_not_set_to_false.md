---
title: "Profiling not set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/profiling_not_set_to_false"
  id: "2f491173-6375-4a84-b28e-a4e2b9a58a69"
  display_name: "Profiling not set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `2f491173-6375-4a84-b28e-a4e2b9a58a69`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

### Description

 When using `kube-apiserver`, `kube-controller-manager`, or `kube-scheduler`, the `--profiling` flag should be set to `false`.

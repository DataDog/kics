---
title: "Always admit admission control plugin set"
group_id: "rules/k8s"
meta:
  name: "k8s/always_admit_admission_control_plugin_set"
  id: "ce30e584-b33f-4c7d-b418-a3d7027f8f60"
  display_name: "Always admit admission control plugin set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `ce30e584-b33f-4c7d-b418-a3d7027f8f60`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should not include the `AlwaysAdmit` plugin.

---
title: "Namespace lifecycle admission control plugin disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/namespace_lifecycle_admission_control_plugin_disabled"
  id: "1ffe7bf7-563b-4b3d-a71d-ba6bd8d49b37"
  display_name: "Namespace lifecycle admission control plugin disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Build Process"
---
## Metadata

**Id:** `1ffe7bf7-563b-4b3d-a71d-ba6bd8d49b37`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Build Process

### Description

 When using `kube-apiserver`, the `--disable-admission-plugins` flag should not include the `NamespaceLifecycle` plugin.

---
title: "Kubelet protect-kernel-defaults set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_protect_kernel_defaults_set_to_false"
  id: "6cf42c97-facd-4fda-b8af-ea4529123355"
  display_name: "Kubelet protect-kernel-defaults set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `6cf42c97-facd-4fda-b8af-ea4529123355`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 The `--protect-kernel-defaults` flag should be set to `true`.

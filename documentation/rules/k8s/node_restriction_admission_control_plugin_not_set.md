---
title: "Node restriction admission control plugin not set"
group_id: "rules/k8s"
meta:
  name: "k8s/node_restriction_admission_control_plugin_not_set"
  id: "33fc6923-6553-4fe6-9d3a-4efa51eb874b"
  display_name: "Node restriction admission control plugin not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `33fc6923-6553-4fe6-9d3a-4efa51eb874b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--enable-admission-plugins` flag should include `NodeRestriction`, and it should be correctly configured in the admission control config file.

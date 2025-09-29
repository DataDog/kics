---
title: "Service account admission control plugin disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_admission_control_plugin_disabled"
  id: "9587c890-0524-40c2-9ce2-663af7c2f063"
  display_name: "Service account admission control plugin disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `9587c890-0524-40c2-9ce2-663af7c2f063`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 When using `kube-apiserver`, the `--disable-admission-plugins` flag should not include the `ServiceAccount` plugin.

---
title: "Anonymous auth is not set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/anonymous_auth_is_not_set_to_false"
  id: "1de5cc51-f376-4638-a940-20f2e85ae238"
  display_name: "Anonymous auth is not set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `1de5cc51-f376-4638-a940-20f2e85ae238`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 When using `kubelet` or `kube-apiserver`, the `--anonymous-auth` flag should be set to `false`.

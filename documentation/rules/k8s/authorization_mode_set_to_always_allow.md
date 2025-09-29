---
title: "Authorization mode set to always allow"
group_id: "rules/k8s"
meta:
  name: "k8s/authorization_mode_set_to_always_allow"
  id: "f1f4d8da-1ac4-47d0-b1aa-91e69d33f7d5"
  display_name: "Authorization mode set to always allow"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `f1f4d8da-1ac4-47d0-b1aa-91e69d33f7d5`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 When using `kubelet`, the `--authorization-mode` flag should not be set to `AlwaysAllow`.

---
title: "Authorization mode node not set"
group_id: "rules/k8s"
meta:
  name: "k8s/authorization_mode_node_not_set"
  id: "4d7ee40f-fc5d-427d-8cac-dffbe22d42d1"
  display_name: "Authorization mode node not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `4d7ee40f-fc5d-427d-8cac-dffbe22d42d1`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 When using `kube-apiserver`, the `--authorization-mode` flag should include `Node`.

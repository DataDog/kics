---
title: "Rotate Kubelet server certificate not active"
group_id: "rules/k8s"
meta:
  name: "k8s/rotate_kubelet_server_certificate_not_active"
  id: "1c621b8e-2c6a-44f5-bd6a-fb0fb7ba33e2"
  display_name: "Rotate Kubelet server certificate not active"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `1c621b8e-2c6a-44f5-bd6a-fb0fb7ba33e2`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 The `RotateKubeletServerCertificate` argument should be set to `true`.

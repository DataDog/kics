---
title: "Kubelet certificate authority not set"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_certificate_authority_not_set"
  id: "ec18a0d3-0069-4a58-a7fb-fbfe0b4bbbe0"
  display_name: "Kubelet certificate authority not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `ec18a0d3-0069-4a58-a7fb-fbfe0b4bbbe0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `kube-apiserver`, the `--kubelet-certificate-authority` flag should be set.

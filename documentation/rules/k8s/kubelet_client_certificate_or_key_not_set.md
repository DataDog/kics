---
title: "Kubelet client certificate or key not set"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_client_certificate_or_key_not_set"
  id: "36a27826-1bf5-49da-aeb0-a60a30c0e834"
  display_name: "Kubelet client certificate or key not set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `36a27826-1bf5-49da-aeb0-a60a30c0e834`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `kube-apiserver`, the `--kubelet-client-key` and `--kubelet-client-certificate` flags should be set.

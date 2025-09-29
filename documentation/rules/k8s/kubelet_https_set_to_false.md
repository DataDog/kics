---
title: "Kubelet HTTPS set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_https_set_to_false"
  id: "cdc8b54e-6b16-4538-a1b0-35849dbe29cf"
  display_name: "Kubelet HTTPS set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `cdc8b54e-6b16-4538-a1b0-35849dbe29cf`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `kube-apiserver`, the `--kubelet-https` flag should not be set to `false`.

---
title: "Non kube-system pod with host mount"
group_id: "rules/k8s"
meta:
  name: "k8s/non_kube_system_pod_with_host_mount"
  id: "aa8f7a35-9923-4cad-bd61-a19b7f6aac91"
  display_name: "Non kube-system pod with host mount"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `aa8f7a35-9923-4cad-bd61-a19b7f6aac91`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 Non `kube-system` workloads should not have a `hostPath` volume mounted.

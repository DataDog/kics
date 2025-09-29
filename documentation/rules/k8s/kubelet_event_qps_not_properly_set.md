---
title: "Kubelet event QPS not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_event_qps_not_properly_set"
  id: "1a07a446-8e61-4e4d-bc16-b0781fcb8211"
  display_name: "Kubelet event QPS not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `1a07a446-8e61-4e4d-bc16-b0781fcb8211`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

### Description

 When using `kubelet`, the `--event-qps` value should be set to `0`.

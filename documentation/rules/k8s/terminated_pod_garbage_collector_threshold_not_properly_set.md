---
title: "Terminated pod garbage collector threshold not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/terminated_pod_garbage_collector_threshold_not_properly_set"
  id: "49113af4-29ca-458e-b8d4-724c01a4a24f"
  display_name: "Terminated pod garbage collector threshold not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata

**Id:** `49113af4-29ca-458e-b8d4-724c01a4a24f`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Availability

### Description

 When using `kube-controller-manager`, the `--terminated-pod-gc-threshold` flag should be set between `0` and `12501`.

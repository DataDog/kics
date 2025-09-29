---
title: "Liveness probe is not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/liveness_probe_is_not_defined"
  id: "ade74944-a674-4e00-859e-c6eab5bde441"
  display_name: "Liveness probe is not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Availability"
---
## Metadata

**Id:** `ade74944-a674-4e00-859e-c6eab5bde441`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Availability

### Description

 A liveness probe can improve availability by restarting unresponsive containers; however, it can cause cascading failures. Define one only when necessary.

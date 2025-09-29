---
title: "StatefulSet without PodDisruptionBudget"
group_id: "rules/k8s"
meta:
  name: "k8s/statefulset_without_pod_disruption_budget"
  id: "1db3a5a5-bf75-44e5-9e44-c56cfc8b1ac5"
  display_name: "StatefulSet without PodDisruptionBudget"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Availability"
---
## Metadata

**Id:** `1db3a5a5-bf75-44e5-9e44-c56cfc8b1ac5`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Availability

### Description

 StatefulSets should be assigned a PodDisruptionBudget to ensure high availability.

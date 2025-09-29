---
title: "StatefulSet without podAntiAffinity"
group_id: "rules/k8s"
meta:
  name: "k8s/statefulset_has_no_pod_anti_affinity"
  id: "d740d048-8ed3-49d3-b77b-6f072f3b669e"
  display_name: "StatefulSet without podAntiAffinity"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Id:** `d740d048-8ed3-49d3-b77b-6f072f3b669e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Resource Management

### Description

 StatefulSets should define a `podAntiAffinity` policy to avoid scheduling multiple Pods on the same node.

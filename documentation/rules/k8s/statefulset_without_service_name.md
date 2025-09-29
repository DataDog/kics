---
title: "StatefulSet without service name"
group_id: "rules/k8s"
meta:
  name: "k8s/statefulset_without_service_name"
  id: "bb241e61-77c3-4b97-9575-c0f8a1e008d0"
  display_name: "StatefulSet without service name"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Availability"
---
## Metadata

**Id:** `bb241e61-77c3-4b97-9575-c0f8a1e008d0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Availability

### Description

 StatefulSets should have an existing headless `serviceName`. The headless service labels should also be applied to StatefulSet labels.

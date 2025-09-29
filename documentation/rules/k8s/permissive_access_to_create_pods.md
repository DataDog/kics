---
title: "Permissive access to create Pods"
group_id: "rules/k8s"
meta:
  name: "k8s/permissive_access_to_create_pods"
  id: "592ad21d-ad9b-46c6-8d2d-fad09d62a942"
  display_name: "Permissive access to create Pods"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `592ad21d-ad9b-46c6-8d2d-fad09d62a942`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 The permission to create Pods in a cluster should be restricted because it allows privilege escalation.

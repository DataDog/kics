---
title: "Deployment without PodDisruptionBudget"
group_id: "rules/k8s"
meta:
  name: "k8s/deployment_without_pod_disruption_budget"
  id: "b23e9b98-0cb6-4fc9-b257-1f3270442678"
  display_name: "Deployment without PodDisruptionBudget"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Availability"
---
## Metadata

**Id:** `b23e9b98-0cb6-4fc9-b257-1f3270442678`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Availability

### Description

 Deployments should be assigned a PodDisruptionBudget to ensure high availability.

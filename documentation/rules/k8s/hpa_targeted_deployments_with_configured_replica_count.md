---
title: "HPA targeted deployments with configured replica count"
group_id: "rules/k8s"
meta:
  name: "k8s/hpa_targeted_deployments_with_configured_replica_count"
  id: "5744cbb8-5946-4b75-a196-ade44449525b"
  display_name: "HPA targeted deployments with configured replica count"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Availability"
---
## Metadata

**Id:** `5744cbb8-5946-4b75-a196-ade44449525b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Availability

### Description

 Deployments targeted by HorizontalPodAutoscaler should not have a statically configured replica count.

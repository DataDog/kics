---
title: "Deployment without podAntiAffinity"
group_id: "rules/k8s"
meta:
  name: "k8s/deployment_has_no_pod_anti_affinity"
  id: "a31b7b82-d994-48c4-bd21-3bab6c31827a"
  display_name: "Deployment without podAntiAffinity"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Id:** `a31b7b82-d994-48c4-bd21-3bab6c31827a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Resource Management

### Description

 Deployments should define a `podAntiAffinity` policy to prevent multiple Pods from being scheduled on the same node.

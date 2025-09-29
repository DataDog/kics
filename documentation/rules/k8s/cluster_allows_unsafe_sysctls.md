---
title: "Cluster allows unsafe sysctls"
group_id: "rules/k8s"
meta:
  name: "k8s/cluster_allows_unsafe_sysctls"
  id: "9127f0d9-2310-42e7-866f-5fd9d20dcbad"
  display_name: "Cluster allows unsafe sysctls"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `9127f0d9-2310-42e7-866f-5fd9d20dcbad`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 A Kubernetes cluster must not allow unsafe sysctls to prevent a Pod from influencing other Pods, harming node health, or gaining CPU or memory outside resource limits. `spec.securityContext.sysctls` must not specify unsafe sysctls, and `allowedUnsafeSysctls` must be undefined.

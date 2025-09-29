---
title: "Using Kubernetes native secret management"
group_id: "rules/k8s"
meta:
  name: "k8s/using_kubernetes_native_secret_management"
  id: "b9c83569-459b-4110-8f79-6305aa33cb37"
  display_name: "Using Kubernetes native secret management"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "INFO"
  category: "Secret Management"
---
## Metadata

**Id:** `b9c83569-459b-4110-8f79-6305aa33cb37`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Info

**Category:** Secret Management

### Description

 Use of an external secret storage and management system should be considered for complex secret management needs, rather than using Kubernetes Secrets directly. Additionally, ensure that access to secrets is carefully limited.

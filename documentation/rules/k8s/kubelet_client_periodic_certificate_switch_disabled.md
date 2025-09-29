---
title: "Kubelet client periodic certificate switch disabled"
group_id: "rules/k8s"
meta:
  name: "k8s/kubelet_client_periodic_certificate_switch_disabled"
  id: "52d70f2e-3257-474c-b3dc-8ad9ba6a061a"
  display_name: "Kubelet client periodic certificate switch disabled"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `52d70f2e-3257-474c-b3dc-8ad9ba6a061a`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 The kubelet argument `--rotate-certificates` should be set to `true`.

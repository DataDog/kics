---
title: "Using unrecommended namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/using_unrecommended_namespace"
  id: "611ab018-c4aa-4ba2-b0f6-a448337509a6"
  display_name: "Using unrecommended namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `611ab018-c4aa-4ba2-b0f6-a448337509a6`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Namespaces such as `default`, `kube-system`, or `kube-public` should not be used.

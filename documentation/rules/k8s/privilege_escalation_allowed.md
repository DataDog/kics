---
title: "Privilege escalation allowed"
group_id: "rules/k8s"
meta:
  name: "k8s/privilege_escalation_allowed"
  id: "5572cc5e-1e4c-4113-92a6-7a8a3bd25e6d"
  display_name: "Privilege escalation allowed"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `5572cc5e-1e4c-4113-92a6-7a8a3bd25e6d`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 Containers should not run with `allowPrivilegeEscalation` to prevent them from gaining more privileges than their parent process.

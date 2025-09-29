---
title: "RBAC roles with exec permission"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_with_exec_permission"
  id: "c589f42c-7924-4871-aee2-1cede9bc7cbc"
  display_name: "RBAC roles with exec permission"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `c589f42c-7924-4871-aee2-1cede9bc7cbc`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 Roles or ClusterRoles with permissions to run commands in containers via `kubectl exec` could be abused by attackers to execute malicious code in case of compromise. To prevent this, the `pods/exec` verb should not be used in production environments.

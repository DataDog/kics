---
title: "RBAC roles with impersonate permission"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_with_impersonate_permission"
  id: "9f85c3f6-26fd-4007-938a-2e0cb0100980"
  display_name: "RBAC roles with impersonate permission"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `9f85c3f6-26fd-4007-938a-2e0cb0100980`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 Roles or ClusterRoles with the `impersonate` permission allow subjects to assume the rights of other users, groups, or service accounts. In case of compromise, attackers may abuse this sudo-like functionality to achieve privilege escalation.

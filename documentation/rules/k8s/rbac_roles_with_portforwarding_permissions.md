---
title: "RBAC roles with port-forwarding permission"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_with_portforwarding_permissions"
  id: "38fa11ef-dbcc-4da8-9680-7e1fd855b6fb"
  display_name: "RBAC roles with port-forwarding permission"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `38fa11ef-dbcc-4da8-9680-7e1fd855b6fb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 Roles or ClusterRoles with permissions to port-forward into pods can open socket-level communication channels to containers. In case of compromise, attackers may abuse this for direct communication that bypasses network security controls.

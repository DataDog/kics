---
title: "RBAC wildcard in rule"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_wildcard_in_rule"
  id: "6b896afb-ca07-467a-b256-1a0077a1c08e"
  display_name: "RBAC wildcard in rule"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata

**Id:** `6b896afb-ca07-467a-b256-1a0077a1c08e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Access Control

### Description

 Roles and ClusterRoles with wildcard RBAC permissions grant excessive rights to the Kubernetes API and should be avoided. The principle of least privilege recommends specifying only the needed objects and actions.

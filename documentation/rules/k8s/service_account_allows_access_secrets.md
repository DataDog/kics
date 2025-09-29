---
title: "ServiceAccount allows access to secrets"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_allows_access_secrets"
  id: "056ac60e-fe07-4acc-9b34-8e1d51716ab9"
  display_name: "ServiceAccount allows access to secrets"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `056ac60e-fe07-4acc-9b34-8e1d51716ab9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 Roles and ClusterRoles, when bound, should not use `get`, `list`, or `watch` verbs.

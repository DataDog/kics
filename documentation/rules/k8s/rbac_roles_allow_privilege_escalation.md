---
title: "RBAC roles allow privilege escalation"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_allow_privilege_escalation"
  id: "8320826e-7a9c-4b0b-9535-578333193432"
  display_name: "RBAC roles allow privilege escalation"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `8320826e-7a9c-4b0b-9535-578333193432`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 Roles or ClusterRoles with `bind` or `escalate` permissions allow subjects to create new bindings with other roles. This is dangerous, as users with these privileges can bind to roles that may exceed their own privileges.


## Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: not-rbac-binder
rules:
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterrolebindings"]
  verbs: ["create"]

```
## Non-Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: rbac-binder
rules:
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterroles"]
  verbs: ["bind"]
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterrolebindings"]
  verbs: ["create"]

```
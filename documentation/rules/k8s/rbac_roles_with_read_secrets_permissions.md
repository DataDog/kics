---
title: "RBAC roles with read secrets permissions"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_with_read_secrets_permissions"
  id: "b7bca5c4-1dab-4c2c-8cbe-3050b9d59b14"
  display_name: "RBAC roles with read secrets permissions"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `b7bca5c4-1dab-4c2c-8cbe-3050b9d59b14`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

### Description

 Roles and ClusterRoles with `get`/`watch`/`list` permissions on Kubernetes Secrets are dangerous and should be avoided. In case of compromise, attackers could abuse these roles to access sensitive data, such as passwords, tokens, and keys.


## Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: role-pod-and-logs-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/logs"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cluster-role-pod-and-pod-logs-reader
rules:
- apiGroups: [""]
  resources: ["pods", "pods/log"]
  verbs: ["get", "list"]

```
## Non-Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: role-secret-reader
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cluster-role-secret-reader
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "list"]

```
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


## Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: my-namespace
  name: allow-exec-neg
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: allow-exec-neg
  namespace: my-namespace
subjects:
- kind: User
  name: bob
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: allow-exec-neg
  apiGroup: ""
```
## Non-Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: my-namespace
  name: allow-exec
rules:
- apiGroups: [""]
  resources: ["pods", "pods/exec"]
  verbs: ["get", "list", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: allow-exec
  namespace: my-namespace
subjects:
- kind: User
  name: bob
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: allow-exec
  apiGroup: ""
```
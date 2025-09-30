---
title: "RBAC roles with attach permission"
group_id: "rules/k8s"
meta:
  name: "k8s/rbac_roles_with_attach_permission"
  id: "d45330fd-f58d-45fb-a682-6481477a0f84"
  display_name: "RBAC roles with attach permission"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `d45330fd-f58d-45fb-a682-6481477a0f84`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

### Description

 Roles or ClusterRoles with permissions to attach to containers via `kubectl attach` could be abused by attackers to read log output (stdout, stderr) and send input data (stdin) to running processes. Additionally, it could allow a malicious user to attach to a privileged container, resulting in privilege escalation. To prevent this, the `pods/attach` verb should not be used in production environments.


## Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: my-namespace
  name: allow-attach-neg
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: allow-attach-neg
  namespace: my-namespace
subjects:
- kind: User
  name: bob
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: allow-attach-neg
  apiGroup: ""

```
## Non-Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: my-namespace
  name: allow-attach
rules:
- apiGroups: [""]
  resources: ["pods", "pods/attach"]
  verbs: ["get", "list", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: allow-attach
  namespace: my-namespace
subjects:
- kind: User
  name: bob
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: allow-attach
  apiGroup: ""

```
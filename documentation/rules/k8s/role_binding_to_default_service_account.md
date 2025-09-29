---
title: "Role binding to default service account"
group_id: "rules/k8s"
meta:
  name: "k8s/role_binding_to_default_service_account"
  id: "1e749bc9-fde8-471c-af0c-8254efd2dee5"
  display_name: "Role binding to default service account"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---
## Metadata

**Id:** `1e749bc9-fde8-471c-af0c-8254efd2dee5`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Defaults

### Description

 No Role or ClusterRole should bind to a default service account.


## Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
- kind: User
  name: jane
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```
## Non-Compliant Code Examples
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods
  namespace: default
subjects:
- kind: User
  name: jane
  apiGroup: rbac.authorization.k8s.io
- kind: ServiceAccount
  name: default
  namespace: kube-system
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```
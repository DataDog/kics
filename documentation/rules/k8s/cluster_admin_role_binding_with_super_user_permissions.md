---
title: "Cluster admin rolebinding with superuser permissions"
group_id: "rules/k8s"
meta:
  name: "k8s/cluster_admin_role_binding_with_super_user_permissions"
  id: "249328b8-5f0f-409f-b1dd-029f07882e11"
  display_name: "Cluster admin rolebinding with superuser permissions"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Access Control"
---
## Metadata

**Id:** `249328b8-5f0f-409f-b1dd-029f07882e11`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles)

### Description

 Ensure that the `cluster-admin` role is used only where required (RBAC).


## Compliant Code Examples
```yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: tiller-clusterrolebinding
subjects:
- kind: ServiceAccount
  name: tiller
  namespace: kube-system
roleRef:
  kind: ClusterRole
  name: view
  apiGroup: ""
# trigger validation

```
## Non-Compliant Code Examples
```yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: tiller-clusterrolebinding
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: ""

```
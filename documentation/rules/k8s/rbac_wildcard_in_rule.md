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

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

### Description

 Roles and ClusterRoles with wildcard RBAC permissions grant excessive rights to the Kubernetes API and should be avoided. The principle of least privilege recommends specifying only the needed objects and actions.


## Compliant Code Examples
```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: opa
  name: configmap-modifier
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["update", "patch"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: opa
  name: configmap-modifier
rules:
- apiGroups: [""]
  resources: ["searchmaps"]
  verbs: ["create", "patch"]
```
## Non-Compliant Code Examples
```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: rbac1
  name: configmap-modifier
rules:
- apiGroups: ["*"]
  resources: ["configmaps"]
  verbs: ["*"]
---
# Define role for OPA/kube-mgmt to update configmaps with policy status.
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: rbac2
  name: configmap-modifier1
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: rbac3
  name: configmap-modifier2
rules:
- operations: ["CREATE", "UPDATE"]
  apiGroups: ["*"]
  apiVersions: ["*"]
  resources: ["*"]
  verbs: ["POST"]

```
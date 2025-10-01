---
title: "Permissive access to create Pods"
group_id: "rules/k8s"
meta:
  name: "k8s/permissive_access_to_create_pods"
  id: "592ad21d-ad9b-46c6-8d2d-fad09d62a942"
  display_name: "Permissive access to create Pods"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `592ad21d-ad9b-46c6-8d2d-fad09d62a942`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#privilege-escalation-prevention-and-bootstrapping)

### Description

 The permission to create Pods in a cluster should be restricted because it allows privilege escalation.


## Compliant Code Examples
```yaml
#this code is a correct code for which the query should not find any result
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader
rules:
- apiGroups: [""]

  resources: ["pods"]
  verbs: ["get", "watch", "list"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader2
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "watch", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader4
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs:
    - "get"
    - "watch"

```

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader
rules:
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - "*"
    verbs:
      - create
      - delete

```
## Non-Compliant Code Examples
```yaml
#this is a problematic code where the query should report a result(s)
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs:
    - "get"
    - "watch"
    - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: secret-reader2
rules:
- apiGroups: [""]
  resources: ["*"]
  verbs: ["get", "watch", "create"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader3
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "watch", "*"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: secret-reader4
rules:
- apiGroups: [""]
  resources: ["*"]
  verbs: ["get", "watch", "*"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader5
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs:
    - "get"
    - "watch"
    - "c*e"
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader6
rules:
- apiGroups: [""]
  resources: ["p*ds"]
  verbs: ["get", "watch", "create"]
```

```yaml
#this is a problematic code where the query should report a result(s)
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: secret-reader
rules:
  - apiGroups:
      - "*"
    resources:
      - "*"
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - custom
    verbs:
      - create
      - delete
  - apiGroups:
      - "*"
    resources:
      - "*"
    verbs:
      - create
      - delete
      - get
      - list
      - patch
      - update
      - watch

```
---
title: "Audit policy does not cover key security concerns"
group_id: "rules/k8s"
meta:
  name: "k8s/audit_policy_not_cover_key_security_concerns"
  id: "1828a670-5957-4bc5-9974-47da228f75e2"
  display_name: "Audit policy does not cover key security concerns"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `1828a670-5957-4bc5-9974-47da228f75e2`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Observability

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/)

### Description

 The audit policy should cover key security concerns about sensitive data logged in Kubernetes audit policies.


## Compliant Code Examples
```yaml
apiVersion: audit.k8s.io/v1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
omitStages:
  - "RequestReceived"
rules:
  - level: Metadata
    resources:
    - group: ""
      resources: ["secrets","configmaps","tokenreviews"]
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods","deployments"]
  - level: RequestResponse
    resources:
    - group: ""
      resources: ["pods/exec", "pods/portforward", "pods/proxy", "services/proxy"]

```
## Non-Compliant Code Examples
```yaml
apiVersion: audit.k8s.io/v1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
rules:
  - level: RequestResponse
    resources:
    - group: ""
      resources: ["secrets","configmaps","tokenreviews"]
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods","deployments"]
  - level: None
    resources:
    - group: ""
      resources: ["pods/exec", "pods/portforward", "pods/proxy", "services/proxy"]

```

```yaml
apiVersion: audit.k8s.io/v1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
omitStages:
  - "RequestReceived"
rules:
  - level: Metadata
    resources:
    - group: ""
      resources: ["secrets","configmaps","tokenreviews"]
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods"]
  - level: RequestResponse
    resources:
    - group: ""
      resources: ["pods/exec", "pods/portforward", "pods/proxy", "services/proxy"]

```

```yaml
apiVersion: audit.k8s.io/v1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
omitStages:
  - "RequestReceived"
rules:

```
---
title: "Secrets used as environment variables"
group_id: "rules/k8s"
meta:
  name: "k8s/secrets_as_environment_variables"
  id: "3d658f8b-d988-41a0-a841-40043121de1e"
  display_name: "Secrets used as environment variables"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Secret Management"
---
## Metadata

**Id:** `3d658f8b-d988-41a0-a841-40043121de1e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Secret Management

#### Learn More

 - [Provider Reference](https://kubernetes.io/docs/concepts/configuration/secret/#using-secrets-as-environment-variables)

### Description

 Containers should not use secrets as environment variables.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secret-env-pod
spec:
  containers:
  - name: mycontainer
    image: redis
  restartPolicy: Never
```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secret-env-pod
spec:
  containers:
  - name: mycontainer
    image: redis
    env:
      - name: SECRET_USERNAME
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: username
      - name: SECRET_PASSWORD
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: password
  restartPolicy: Never
---
apiVersion: v1
kind: Pod
metadata:
  name: envfrom-secret
spec:
  containers:
  - name: envars-test-container
    image: nginx
    envFrom:
    - secretRef:
        name: test-secret
```
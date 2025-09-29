---
title: "Image without digest"
group_id: "rules/k8s"
meta:
  name: "k8s/image_without_digest"
  id: "7c81d34c-8e5a-402b-9798-9f442630e678"
  display_name: "Image without digest"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `7c81d34c-8e5a-402b-9798-9f442630e678`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 Images should be specified with their digests to ensure integrity.


## Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: private-image-test-1
spec:
  containers:
    - name: uses-private-image
      image: image@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb
      imagePullPolicy: Always
      command: [ "echo", "SUCCESS" ]
```
## Non-Compliant Code Examples
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: private-image-test-1
spec:
  containers:
    - name: uses-private-image
      image: $PRIVATE_IMAGE_NAME
      imagePullPolicy: Always
      command: [ "echo", "SUCCESS" ]

```
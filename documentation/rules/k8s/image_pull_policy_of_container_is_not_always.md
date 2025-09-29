---
title: "Image pull policy of the container is not set to always"
group_id: "rules/k8s"
meta:
  name: "k8s/image_pull_policy_of_container_is_not_always"
  id: "caa3479d-885d-4882-9aac-95e5e78ef5c2"
  display_name: "Image pull policy of the container is not set to always"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `caa3479d-885d-4882-9aac-95e5e78ef5c2`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Insecure Configurations

### Description

 The container `imagePullPolicy` must be defined and set to `Always`.

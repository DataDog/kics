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

### Description

 Containers should not use secrets as environment variables.

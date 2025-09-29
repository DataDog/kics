---
title: "Object is using a deprecated API version"
group_id: "rules/k8s"
meta:
  name: "k8s/object_is_using_a_deprecated_api_version"
  id: "94b76ea5-e074-4ca2-8a03-c5a606e30645"
  display_name: "Object is using a deprecated API version"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata

**Id:** `94b76ea5-e074-4ca2-8a03-c5a606e30645`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Best Practices

### Description

 Kubernetes APIs evolve over time and may be removed in newer releases. To prevent incompatibilities when upgrading Kubernetes, deprecated APIs should be replaced with newer, stable API versions.

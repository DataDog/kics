---
title: "Containers missing drop capabilities"
group_id: "rules/k8s"
meta:
  name: "k8s/no_drop_capabilities_for_containers"
  id: "268ca686-7fb7-4ae9-b129-955a2a89064e"
  display_name: "Containers missing drop capabilities"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata

**Id:** `268ca686-7fb7-4ae9-b129-955a2a89064e`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Low

**Category:** Best Practices

### Description

 Containers should configure drop capabilities in the security context.

---
title: "Seccomp profile is not configured"
group_id: "rules/k8s"
meta:
  name: "k8s/seccomp_profile_is_not_configured"
  id: "f377b83e-bd07-4f48-a591-60c82b14a78b"
  display_name: "Seccomp profile is not configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `f377b83e-bd07-4f48-a591-60c82b14a78b`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Insecure Configurations

### Description

 Containers should be configured with a secure seccomp profile to restrict potentially dangerous syscalls.

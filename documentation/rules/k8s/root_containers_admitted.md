---
title: "Root containers admitted"
group_id: "rules/k8s"
meta:
  name: "k8s/root_containers_admitted"
  id: "e3aa0612-4351-4a0d-983f-aefea25cf203"
  display_name: "Root containers admitted"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata

**Id:** `e3aa0612-4351-4a0d-983f-aefea25cf203`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Best Practices

### Description

 Containers must not be allowed to run with root privileges; `privileged`, `allowPrivilegeEscalation`, and `readOnlyRootFilesystem` must be set to `false`; `runAsUser.rule` must be set to `MustRunAsNonRoot`; and adding the root group must be forbidden.

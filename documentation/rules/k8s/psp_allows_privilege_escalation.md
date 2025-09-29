---
title: "PSP allows privilege escalation"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_allows_privilege_escalation"
  id: "87554eef-154d-411d-bdce-9dbd91e56851"
  display_name: "PSP allows privilege escalation"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `87554eef-154d-411d-bdce-9dbd91e56851`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 PodSecurityPolicy should not allow privilege escalation.

---
title: "Incorrect volume claim access mode ReadWriteOnce"
group_id: "rules/k8s"
meta:
  name: "k8s/incorrect_volume_claim_access_mode_read_write_once"
  id: "3878dc92-8e5d-47cf-9cdd-7590f71d21b9"
  display_name: "Incorrect volume claim access mode ReadWriteOnce"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Build Process"
---
## Metadata

**Id:** `3878dc92-8e5d-47cf-9cdd-7590f71d21b9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Build Process

### Description

 Kubernetes StatefulSets must have one volume claim template with the access mode `ReadWriteOnce`.

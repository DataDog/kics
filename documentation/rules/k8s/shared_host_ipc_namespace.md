---
title: "Shared host IPC namespace"
group_id: "rules/k8s"
meta:
  name: "k8s/shared_host_ipc_namespace"
  id: "cd290efd-6c82-4e9d-a698-be12ae31d536"
  display_name: "Shared host IPC namespace"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata

**Id:** `cd290efd-6c82-4e9d-a698-be12ae31d536`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Resource Management

### Description

 Containers should not share the host IPC namespace.

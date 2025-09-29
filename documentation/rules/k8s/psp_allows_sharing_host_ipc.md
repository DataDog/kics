---
title: "PSP allows sharing host IPC"
group_id: "rules/k8s"
meta:
  name: "k8s/psp_allows_sharing_host_ipc"
  id: "80f93444-b240-4ebb-a4c6-5c40b76c04ea"
  display_name: "PSP allows sharing host IPC"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `80f93444-b240-4ebb-a4c6-5c40b76c04ea`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Insecure Configurations

### Description

 PodSecurityPolicy allows containers to share the host IPC namespace.

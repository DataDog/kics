---
title: "Volume mount with OS directory write permissions"
group_id: "rules/k8s"
meta:
  name: "k8s/volume_mount_with_os_directory_write_permissions"
  id: "b7652612-de4e-4466-a0bf-1cd81f0c6063"
  display_name: "Volume mount with OS directory write permissions"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Resource Management"
---
## Metadata

**Id:** `b7652612-de4e-4466-a0bf-1cd81f0c6063`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Resource Management

### Description

 Containers can mount sensitive directories from the host, granting potentially dangerous access to critical host configurations and binaries.

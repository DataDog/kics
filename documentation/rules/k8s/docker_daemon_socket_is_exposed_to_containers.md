---
title: "Docker daemon socket is exposed to containers"
group_id: "rules/k8s"
meta:
  name: "k8s/docker_daemon_socket_is_exposed_to_containers"
  id: "a6f34658-fdfb-4154-9536-56d516f65828"
  display_name: "Docker daemon socket is exposed to containers"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `a6f34658-fdfb-4154-9536-56d516f65828`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Access Control

### Description

 The Docker daemon socket should not be exposed to containers.

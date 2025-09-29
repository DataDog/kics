---
title: "Container running as root"
group_id: "rules/k8s"
meta:
  name: "k8s/containers_running_as_root"
  id: "cf34805e-3872-4c08-bf92-6ff7bb0cfadb"
  display_name: "Container running as root"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata

**Id:** `cf34805e-3872-4c08-bf92-6ff7bb0cfadb`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Best Practices

### Description

 Containers should run only as a non-root user. This limits the exploitability of security misconfigurations and restricts an attacker's options in case of compromise.

---
title: "etcd TLS certificate files not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_tls_certificate_files_not_properly_set"
  id: "075ca296-6768-4322-aea2-ba5063b969a9"
  display_name: "etcd TLS certificate files not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `075ca296-6768-4322-aea2-ba5063b969a9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `etcd`, the `--cert-file` and `--key-file` flags should be set.

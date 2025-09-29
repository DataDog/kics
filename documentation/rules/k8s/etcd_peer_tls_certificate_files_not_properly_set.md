---
title: "etcd peer TLS certificate files not properly set"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_peer_tls_certificate_files_not_properly_set"
  id: "09bb9e96-8da3-4736-b89a-b36814acca60"
  display_name: "etcd peer TLS certificate files not properly set"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `09bb9e96-8da3-4736-b89a-b36814acca60`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** High

**Category:** Networking and Firewall

### Description

 When using `etcd`, the `--peer-cert-file` and `--peer-key-file` flags should be set.

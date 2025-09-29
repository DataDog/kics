---
title: "Peer auto TLS set to true"
group_id: "rules/k8s"
meta:
  name: "k8s/peer_auto_tls_set_to_true"
  id: "ae8827e2-4af9-4baa-9998-87539ae0d6f0"
  display_name: "Peer auto TLS set to true"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `ae8827e2-4af9-4baa-9998-87539ae0d6f0`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `etcd`, the `--peer-auto-tls` flag should be set to `false`.

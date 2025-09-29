---
title: "etcd TLS certificate not properly configured"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_tls_certificate_not_properly_configured"
  id: "895a5a95-3756-4b04-9924-2f3bc93181bd"
  display_name: "etcd TLS certificate not properly configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `895a5a95-3756-4b04-9924-2f3bc93181bd`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Networking and Firewall

### Description

 When using `kube-apiserver`, the `--etcd-certfile` and `--etcd-keyfile` flags should be set.

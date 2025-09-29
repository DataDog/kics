---
title: "etcd peer client certificate authentication set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_peer_client_certificate_authentication_set_to_false"
  id: "b7d0181d-0a9b-4611-9d1c-1ad4f0b620ff"
  display_name: "etcd peer client certificate authentication set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `b7d0181d-0a9b-4611-9d1c-1ad4f0b620ff`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `etcd`, the `--peer-client-cert-auth` flag should be set to `true`.

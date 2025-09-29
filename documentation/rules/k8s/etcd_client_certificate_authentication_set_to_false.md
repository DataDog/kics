---
title: "etcd client certificate authentication set to false"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_client_certificate_authentication_set_to_false"
  id: "9391103a-d8d7-4671-ac5d-606ba7ccb0ac"
  display_name: "etcd client certificate authentication set to false"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `9391103a-d8d7-4671-ac5d-606ba7ccb0ac`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `etcd`, the `--client-cert-auth` flag should be set.

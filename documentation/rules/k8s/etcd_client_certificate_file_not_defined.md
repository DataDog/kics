---
title: "etcd client certificate file not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/etcd_client_certificate_file_not_defined"
  id: "3f5ff8a7-5ad6-4d02-86f5-666307da1b20"
  display_name: "etcd client certificate file not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `3f5ff8a7-5ad6-4d02-86f5-666307da1b20`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 When using `kube-apiserver`, the `--etcd-cafile` flag should be set.

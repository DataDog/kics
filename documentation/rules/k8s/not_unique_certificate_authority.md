---
title: "Certificate authority is not unique"
group_id: "rules/k8s"
meta:
  name: "k8s/not_unique_certificate_authority"
  id: "cb7e695d-6a85-495c-b15f-23aed2519303"
  display_name: "Certificate authority is not unique"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata

**Id:** `cb7e695d-6a85-495c-b15f-23aed2519303`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Secret Management

### Description

 The certificate authority should be unique for etcd.

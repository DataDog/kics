---
title: "Root CA file not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/root_ca_file_not_defined"
  id: "05fb986f-ac73-4ebb-a5b2-7faafa93d882"
  display_name: "Root CA file not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `05fb986f-ac73-4ebb-a5b2-7faafa93d882`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Encryption

### Description

 When using `kube-controller-manager`, the `--root-ca-file` flag should be set.

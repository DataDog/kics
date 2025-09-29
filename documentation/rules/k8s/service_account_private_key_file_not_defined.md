---
title: "Service account private key file not defined"
group_id: "rules/k8s"
meta:
  name: "k8s/service_account_private_key_file_not_defined"
  id: "ccc98ff7-68a7-436e-9218-185cb0b0b780"
  display_name: "Service account private key file not defined"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `ccc98ff7-68a7-436e-9218-185cb0b0b780`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Encryption

### Description

 When using `kube-controller-manager`, the `--service-account-private-key-file` flag should be set.

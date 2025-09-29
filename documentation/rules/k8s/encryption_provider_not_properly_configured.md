---
title: "Encryption provider not properly configured"
group_id: "rules/k8s"
meta:
  name: "k8s/encryption_provider_not_properly_configured"
  id: "10efce34-5af6-4d83-b414-9e096d5a06a9"
  display_name: "Encryption provider not properly configured"
  cloud_provider: "k8s"
  framework: "Kubernetes"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `10efce34-5af6-4d83-b414-9e096d5a06a9`

**Cloud Provider:** k8s

**Framework:** Kubernetes

**Severity:** Medium

**Category:** Encryption

### Description

 The encryption configuration should include at least one provider: `aescbc`, `kms`, or `secretbox`.

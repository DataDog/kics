---
title: "Google Compute SSL Policy Weak Cipher In Use"
meta:
  name: "gcp/google_compute_ssl_policy_weak_cipher_in_use"
  id: "14a457f0-473d-4d1d-9e37-6d99b355b336"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `gcp/google_compute_ssl_policy_weak_cipher_in_use`

**Id:** `14a457f0-473d-4d1d-9e37-6d99b355b336`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Encryption

## Description
This query confirms if Google Compute SSL Policy Weak Chyper Suits is Enabled, to do so we need to check if TLS is TLS_1_2, because other version have Weak Chypers

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_ssl_policy)

## Non-Compliant Code Examples
```terraform
resource "google_compute_ssl_policy" "positive1" {
  name            = "custom-ssl-policy"
  min_tls_version = "TLS_1_1"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}

resource "google_compute_ssl_policy" "positive2" {
  name            = "custom-ssl-policy"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}
```

## Compliant Code Examples
```terraform
resource "google_compute_ssl_policy" "negative1" {
  name            = "custom-ssl-policy"
  min_tls_version = "TLS_1_2"
  profile         = "CUSTOM"
  custom_features = ["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"]
}
```
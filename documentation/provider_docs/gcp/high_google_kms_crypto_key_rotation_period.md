---
title: "High Google KMS Crypto Key Rotation Period"
meta:
  name: "gcp/high_google_kms_crypto_key_rotation_period"
  id: "d8c57c4e-bf6f-4e32-a2bf-8643532de77b"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Secret Management"
---

## Metadata
**Name:** `gcp/high_google_kms_crypto_key_rotation_period`

**Id:** `d8c57c4e-bf6f-4e32-a2bf-8643532de77b`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Secret Management

## Description
KMS encryption keys should be rotated every 90 days or less. A short lifetime of encryption keys reduces the potential blast radius in case of compromise.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/kms_crypto_key)

## Non-Compliant Code Examples
```terraform
resource "google_kms_crypto_key" "positive1" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.id
  rotation_period = "77760009s"
  lifecycle {
    prevent_destroy = true
  }
}

resource "google_kms_crypto_key" "positive2" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.id
  lifecycle {
    prevent_destroy = true
  }
}

```

## Compliant Code Examples
```terraform
resource "google_kms_crypto_key" "negative1" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.id
  rotation_period = "100000s"
  lifecycle {
    prevent_destroy = true
  }
}

```
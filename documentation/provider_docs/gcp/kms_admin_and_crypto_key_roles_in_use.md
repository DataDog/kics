---
title: "KMS Admin and CryptoKey Roles In Use"
meta:
  name: "gcp/kms_admin_and_crypto_key_roles_in_use"
  id: "92e4464a-4139-4d57-8742-b5acc0347680"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `gcp/kms_admin_and_crypto_key_roles_in_use`

**Id:** `92e4464a-4139-4d57-8742-b5acc0347680`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Access Control

## Description
Google Project IAM Policy should not assign a KMS admin role and CryptoKey role to the same member

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam#policy_data)

## Non-Compliant Code Examples
```terraform
resource "google_project_iam_policy" "positive1" {
  project     = "your-project-id"
  policy_data = data.google_iam_policy.positive1.policy_data
}

data "google_iam_policy" "positive1" {
  binding {
    role = "roles/cloudkms.admin"

    members = [
      "user:jane@example.com",
    ]
  }

  binding {
    role = "roles/cloudkms.cryptoKeyDecrypter"

    members = [
      "user:jane@example.com",
    ]
  }
}

```

## Compliant Code Examples
```terraform
resource "google_project_iam_policy" "negative1" {
  project     = "your-project-id"
  policy_data = data.google_iam_policy.negative1.policy_data
}

data "google_iam_policy" "negative1" {
  binding {
    role = "roles/cloudkms.admin"

    members = [
      "user:jane@example.com",
    ]
  }

  binding {
    role = "roles/cloudkms.cryptoKeyDecrypter"

    members = [
      "user:jane2@example.com",
    ]
  }
}

```
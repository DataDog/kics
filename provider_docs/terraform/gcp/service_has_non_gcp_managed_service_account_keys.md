---
title: "There are non GCP-managed service account keys for a service account"
meta:
  name: "terraform/service_has_non_gcp_managed_service_account_keys"
  id: "3jh54js8-e5f6-7890-abcd-ef1234567890"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/service_has_non_gcp_managed_service_account_keys`
**Id:** `3jh54js8-e5f6-7890-abcd-ef1234567890`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
Service account keys provide access to GCP resources, and manually created keys introduce significant security risks as they can be leaked, shared inappropriately, or remain active indefinitely without rotation. GCP-managed service account keys follow security best practices by default, with automatic key rotation and secure storage handled by Google. To ensure security, avoid manually specifying key data in your Terraform configuration like `public_key_data = "dummy-key"` and instead rely on GCP's managed keys by omitting this attribute.

#### Learn More

 - [Provider Reference](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys)

## Non-Compliant Code Examples
```gcp
resource "google_service_account_key" "bad_key" {
  service_account_id = "projects/my-project/serviceAccounts/my-service-account"
  public_key_data    = "dummy-key"
}

```

## Compliant Code Examples
```gcp
resource "google_service_account_key" "bad_key" {
  service_account_id = "projects/my-project/serviceAccounts/my-service-account"
}

```
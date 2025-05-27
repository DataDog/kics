---
title: "Cloud Storage Bucket Versioning Disabled"
meta:
  name: "gcp/cloud_storage_bucket_versioning_disabled"
  id: "e7e961ac-d17e-4413-84bc-8a1fbe242944"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `gcp/cloud_storage_bucket_versioning_disabled`

**Id:** `e7e961ac-d17e-4413-84bc-8a1fbe242944`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Observability

## Description
Cloud Storage Bucket should have versioning enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket#enabled)

## Non-Compliant Code Examples
```terraform
resource "google_storage_bucket" "positive1" {
  name     = "foo"
  location = "EU"

  versioning = {
    enabled = false
  }
}

resource "google_storage_bucket" "positive2" {
  name     = "foo"
  location = "EU"
}
```

## Compliant Code Examples
```terraform
resource "google_storage_bucket" "negative1" {
  name     = "foo"
  location = "EU"

  versioning = {
    enabled = true
  }
}
```
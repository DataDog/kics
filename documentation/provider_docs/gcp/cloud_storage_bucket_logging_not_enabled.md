---
title: "Cloud Storage Bucket Logging Not Enabled"
meta:
  name: "gcp/cloud_storage_bucket_logging_not_enabled"
  id: "d6cabc3a-d57e-48c2-b341-bf3dd4f4a120"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `gcp/cloud_storage_bucket_logging_not_enabled`

**Id:** `d6cabc3a-d57e-48c2-b341-bf3dd4f4a120`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Observability

## Description
Cloud storage bucket should have logging enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket#log_bucket)

## Non-Compliant Code Examples
```terraform
resource "google_storage_bucket" "positive1" {
  name          = "auto-expiring-bucket"
  location      = "US"
  force_destroy = true

  lifecycle_rule {
    condition {
      age = 3
    }
    action {
      type = "Delete"
    }
  }
}

```

## Compliant Code Examples
```terraform
resource "google_storage_bucket" "negative1" {
  name          = "auto-expiring-bucket"
  location      = "US"
  force_destroy = true

  logging {
	logBucket = "example-logs-bucket"
  }

  lifecycle_rule {
    condition {
      age = 3
    }
    action {
      type = "Delete"
    }
  }
}
```
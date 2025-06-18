---
title: "Cloud Storage Bucket Logging Not Enabled"
meta:
  name: "terraform/cloud_storage_bucket_logging_not_enabled"
  id: "d6cabc3a-d57e-48c2-b341-bf3dd4f4a120"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/cloud_storage_bucket_logging_not_enabled`
**Id:** `d6cabc3a-d57e-48c2-b341-bf3dd4f4a120`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
Cloud storage buckets should have logging enabled to capture access and usage data for auditing and monitoring purposes. If the `logging` block is not configured in the Terraform resource, such as  

```
resource "google_storage_bucket" "auto_expiring_bucket" {
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

access to and modifications of storage objects may go unnoticed, making it difficult to detect unauthorized activities or investigate security incidents. Enabling logging helps improve visibility into data access and can aid in compliance efforts.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket#log_bucket)

## Non-Compliant Code Examples
```gcp
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
```gcp
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
---
title: "IAM Audit Not Properly Configured"
meta:
  name: "gcp/iam_audit_not_properly_configured"
  id: "89fe890f-b480-460c-8b6b-7d8b1468adb4"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `gcp/iam_audit_not_properly_configured`

**Id:** `89fe890f-b480-460c-8b6b-7d8b1468adb4`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Observability

## Description
Audit Logging Configuration is defective

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam#google_project_iam_audit_config)

## Non-Compliant Code Examples
```terraform
resource "google_project_iam_audit_config" "positive1" {
  project = "your-project-id"
  service = "some_specific_service"
  audit_log_config {
    log_type = "ADMIN_READ"
  }
  audit_log_config {
    log_type = "DATA_READ"
    exempted_members = [
      "user:joebloggs@hashicorp.com"
    ]
  }
}

resource "google_project_iam_audit_config" "positive2" {
  project = "your-project-id"
  service = "allServices"
  audit_log_config {
    log_type = "INVALID_TYPE"
  }
  audit_log_config {
    log_type = "DATA_READ"
    exempted_members = [
        "user:joebloggs@hashicorp.com"
    ]
  }
}
```

## Compliant Code Examples
```terraform
resource "google_project_iam_audit_config" "negative1" {
  project = "your-project-id"
  service = "allServices"
  audit_log_config {
    log_type = "ADMIN_READ"
  }
  audit_log_config {
    log_type = "DATA_READ"
  }
}
```
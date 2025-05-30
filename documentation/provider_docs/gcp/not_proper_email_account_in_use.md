---
title: "Not Proper Email Account In Use"
meta:
  name: "gcp/not_proper_email_account_in_use"
  id: "9356962e-4a4f-4d06-ac59-dc8008775eaa"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/not_proper_email_account_in_use`

**Id:** `9356962e-4a4f-4d06-ac59-dc8008775eaa`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Insecure Configurations

## Description
Gmail accounts are being used instead of corporate credentials

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam#google_project_iam_binding)

## Non-Compliant Code Examples
```terraform
resource "google_project_iam_binding" "positive1" {
  project = "your-project-id"
  role    = "roles/editor"

  members = [
    "user:jane@gmail.com",
  ]
}
```

## Compliant Code Examples
```terraform
resource "google_project_iam_binding" "negative1" {
  project = "your-project-id"
  role    = "roles/editor"

  members = [
    "user:jane@example.com",
  ]
}
```
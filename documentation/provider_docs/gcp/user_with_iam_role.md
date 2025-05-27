---
title: "User with IAM Role"
meta:
  name: "gcp/user_with_iam_role"
  id: "704fcc44-a58f-4af5-82e2-93f2a58ef918"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Access Control"
---

## Metadata
**Name:** `gcp/user_with_iam_role`

**Id:** `704fcc44-a58f-4af5-82e2-93f2a58ef918`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Access Control

## Description
As a best practice, it is better to assign an IAM Role to a group than to a user

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/iam_policy#role)

## Non-Compliant Code Examples
```terraform
resource "google_project_iam_binding" "positive2" {
  project = "your-project-id"
  role    = "roles/container.admin"

  members = [
    "user:jane@example.com",
  ]

  condition {
    title       = "expires_after_2019_12_31"
    description = "Expiring at midnight of 2019-12-31"
    expression  = "request.time < timestamp(\"2020-01-01T00:00:00Z\")"
  }
}

resource "google_project_iam_member" "positive3" {
  project = "your-project-id"
  role    = "roles/editor"
  member  = "user:jane@example.com"
}

```

```terraform
data "google_iam_policy" "positive" {
  binding {
    role = "roles/apigee.runtimeAgent"

    members = [
      "user:jane@example.com",
    ]
  }
}

```

## Compliant Code Examples
```terraform
data "google_iam_policy" "negative" {
  binding {
    role = "roles/apigee.runtimeAgent"

    members = [
      "group:jane@example.com",
    ]
  }
}

```
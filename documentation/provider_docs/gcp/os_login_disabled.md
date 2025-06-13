---
title: "OSLogin Disabled"
meta:
  name: "gcp/os_login_disabled"
  id: "32ecd6eb-0711-421f-9627-1a28d9eff217"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `gcp/os_login_disabled`

**Id:** `32ecd6eb-0711-421f-9627-1a28d9eff217`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Access Control

## Description
Verifies that the OSLogin is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_project_metadata#metadata)

## Non-Compliant Code Examples
```terraform
resource "google_compute_project_metadata" "positive1" {
  metadata = {
    enable-oslogin = false
  }
}

resource "google_compute_project_metadata" "positive2" {
  metadata = {
      foo  = "bar"
  }
}

```

## Compliant Code Examples
```terraform
resource "google_compute_project_metadata" "negative1" {
  metadata = {
    enable-oslogin = true
  }
}

```
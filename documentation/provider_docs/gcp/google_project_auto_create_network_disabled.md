---
title: "Google Project Auto Create Network Disabled"
meta:
  name: "gcp/google_project_auto_create_network_disabled"
  id: "59571246-3f62-4965-a96f-c7d97e269351"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/google_project_auto_create_network_disabled`

**Id:** `59571246-3f62-4965-a96f-c7d97e269351`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Verifies if the Google Project Auto Create Network is Disabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project)

## Non-Compliant Code Examples
```terraform
resource "google_project" "positive1" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
  auto_create_network = true
}

resource "google_project" "positive2" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
}

```

## Compliant Code Examples
```terraform
resource "google_project" "negative1" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
  auto_create_network = false
}

```
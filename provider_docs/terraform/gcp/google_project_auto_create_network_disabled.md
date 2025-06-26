---
title: "Google Project Auto Create Network Disabled"
meta:
  name: "terraform/google_project_auto_create_network_disabled"
  id: "59571246-3f62-4965-a96f-c7d97e269351"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/google_project_auto_create_network_disabled`
**Id:** `59571246-3f62-4965-a96f-c7d97e269351`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
This check ensures that the `auto_create_network` attribute in the `google_project` resource is set to `false`. When `auto_create_network` is set to `true` or left unset (the default), Google Cloud automatically creates a default network with permissive firewall rules, potentially exposing resources to unauthorized access. Secure configuration requires explicitly setting `auto_create_network = false` as shown below:

```
resource "google_project" "example" {
  name                = "My Project"
  project_id          = "your-project-id"
  org_id              = "1234567"
  auto_create_network = false
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project)

## Non-Compliant Code Examples
```gcp
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
```gcp
resource "google_project" "negative1" {
  name       = "My Project"
  project_id = "your-project-id"
  org_id     = "1234567"
  auto_create_network = false
}

```
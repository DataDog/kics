---
title: "Default Service Account In Use"
meta:
  name: "terraform/default_service_account_in_use"
  id: "737a0dd9-0aaa-4145-8118-f01778262b8a"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/default_service_account_in_use`
**Id:** `737a0dd9-0aaa-4145-8118-f01778262b8a`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Insecure Configurations
## Description
Default service accounts should not be actively used

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_account#automount_service_account_token)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_service_account" "example" {
  metadata {
    name = "default"
  }
}

resource "kubernetes_service_account" "example2" {
  metadata {
    name = "default"
  }

  automount_service_account_token = true
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_service_account" "example3" {
  metadata {
    name = "default"
  }

  automount_service_account_token = false
}

```
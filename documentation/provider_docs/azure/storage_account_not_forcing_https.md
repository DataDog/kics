---
title: "Storage Account Not Forcing HTTPS"
meta:
  name: "azure/storage_account_not_forcing_https"
  id: "12944ec4-1fa0-47be-8b17-42a034f937c2"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `azure/storage_account_not_forcing_https`

**Id:** `12944ec4-1fa0-47be-8b17-42a034f937c2`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Encryption

## Description
Storage Accounts should enforce the use of HTTPS

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_account#https_traffic_only_enabled-1)

## Non-Compliant Code Examples
```terraform
resource "azurerm_storage_account" "positive1" {
  name                       = "example1"
  resource_group_name        = data.azurerm_resource_group.example.name
  location                   = data.azurerm_resource_group.example.location
  account_tier               = "Standard"
  account_replication_type   = "GRS"
  https_traffic_only_enabled = false
}

```

## Compliant Code Examples
```terraform
resource "azurerm_storage_account" "positive2" {
  name                     = "example2"
  resource_group_name      = data.azurerm_resource_group.example.name
  location                 = data.azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"

  # will not flag because https_traffic_only_enabled is set to true by default so we do not error
}

```

```terraform
resource "azurerm_storage_account" "negative1" {
  name                       = "example"
  resource_group_name        = data.azurerm_resource_group.example.name
  location                   = data.azurerm_resource_group.example.location
  account_tier               = "Standard"
  account_replication_type   = "GRS"
  https_traffic_only_enabled = true # Set to true as desired
}

```
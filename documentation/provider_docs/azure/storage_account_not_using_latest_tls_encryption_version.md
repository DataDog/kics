---
title: "Storage Account Not Using Latest TLS Encryption Version"
meta:
  name: "azure/storage_account_not_using_latest_tls_encryption_version"
  id: "8263f146-5e03-43e0-9cfe-db960d56d1e7"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `azure/storage_account_not_using_latest_tls_encryption_version`

**Id:** `8263f146-5e03-43e0-9cfe-db960d56d1e7`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Encryption

## Description
Ensure Storage Account is using the latest version of TLS encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_account)

## Non-Compliant Code Examples
```terraform
resource "azurerm_storage_account" "positive2" {
  name                     = "storageaccountname"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  min_tls_version = "TLS1_1"

  tags = {
    environment = "staging"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_storage_account" "negative1" {
  name                     = "storageaccountname"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "GRS"
  min_tls_version = "TLS1_2"

  tags = {
    environment = "staging"
  }
}

```
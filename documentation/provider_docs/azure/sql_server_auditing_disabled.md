---
title: "SQL Server Auditing Disabled"
meta:
  name: "azure/sql_server_auditing_disabled"
  id: "f7e296b0-6660-4bc5-8f87-22ac4a815edf"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/sql_server_auditing_disabled`

**Id:** `f7e296b0-6660-4bc5-8f87-22ac4a815edf`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that for SQL Servers, 'Auditing' is set to 'On'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/sql_server)

## Non-Compliant Code Examples
```terraform
resource "azurerm_sql_server" "positive1" {
    name                         = "mssqlserver"
    resource_group_name          = azurerm_resource_group.example.name
    location                     = azurerm_resource_group.example.location
    version                      = "12.0"
    administrator_login          = "mradministrator"
    administrator_login_password = "thisIsDog11"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_sql_server" "negative1" {
    name                         = "mssqlserver"
    resource_group_name          = azurerm_resource_group.example.name
    location                     = azurerm_resource_group.example.location
    version                      = "12.0"
    administrator_login          = "mradministrator"
    administrator_login_password = "thisIsDog11"

    extended_auditing_policy {
       storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
       storage_account_access_key = azurerm_storage_account.example.primary_access_key
       storage_account_access_key_is_secondary = true
       retention_in_days                       = 90
    }
}
```
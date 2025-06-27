---
title: "SQL Server Auditing Disabled"
meta:
  name: "azure/sql_server_auditing_disabled"
  id: "f7e296b0-6660-4bc5-8f87-22ac4a815edf"
  display_name: "SQL Server Auditing Disabled"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Name:** `azure/sql_server_auditing_disabled`

**Query Name** `SQL Server Auditing Disabled`

**Id:** `f7e296b0-6660-4bc5-8f87-22ac4a815edf`

**Cloud Provider:** azure

**Platform** Terraform

**Severity:** Medium

**Category:** Observability

## Description
SQL Server auditing should be enabled to ensure that all database events and activities are properly logged for security monitoring and compliance purposes. Without the `extended_auditing_policy` block set as shown below, malicious or unauthorized actions may go undetected, increasing the risk of data breaches and making forensic analysis difficult. For example, secure Terraform configuration includes:

```
extended_auditing_policy {
   storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
   storage_account_access_key = azurerm_storage_account.example.primary_access_key
   storage_account_access_key_is_secondary = true
   retention_in_days                       = 90
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/sql_server)


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
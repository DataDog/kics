---
title: "SQL Server Alert Email Disabled"
meta:
  name: "azure/sql_server_alert_email_disabled"
  id: "55975007-f6e7-4134-83c3-298f1fe4b519"
  cloud_provider: "azure"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/sql_server_alert_email_disabled`

**Id:** `55975007-f6e7-4134-83c3-298f1fe4b519`

**Cloud Provider:** azure

**Severity:** Info

**Category:** Best Practices

## Description
SQL Server alert email should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mssql_server_security_alert_policy#email_account_admins)

## Non-Compliant Code Examples
```terraform
resource "azurerm_mssql_server_security_alert_policy" "positive2" {
  resource_group_name        = azurerm_resource_group.example.name
  server_name                = azurerm_sql_server.example.name
  state                      = "Enabled"
  storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
  retention_days = 20
  email_account_admins = false
}

```

```terraform
resource "azurerm_mssql_server_security_alert_policy" "positive1" {
  resource_group_name        = azurerm_resource_group.example.name
  server_name                = azurerm_sql_server.example.name
  state                      = "Enabled"
  storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
  retention_days = 20
}

```

## Compliant Code Examples
```terraform
resource "azurerm_mssql_server_security_alert_policy" "negative" {
  resource_group_name        = azurerm_resource_group.example.name
  server_name                = azurerm_sql_server.example.name
  state                      = "Enabled"
  storage_endpoint           = azurerm_storage_account.example.primary_blob_endpoint
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
  disabled_alerts = [
    "Sql_Injection",
    "Data_Exfiltration"
  ]
  retention_days = 20
  email_account_admins = true
}


```
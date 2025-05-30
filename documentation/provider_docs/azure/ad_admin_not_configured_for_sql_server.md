---
title: "AD Admin Not Configured For SQL Server"
meta:
  name: "azure/ad_admin_not_configured_for_sql_server"
  id: "a3a055d2-9a2e-4cc9-b9fb-12850a1a3a4b"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/ad_admin_not_configured_for_sql_server`

**Id:** `a3a055d2-9a2e-4cc9-b9fb-12850a1a3a4b`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
The Active Directory Administrator is not configured for a SQL server

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/sql_active_directory_administrator)

## Non-Compliant Code Examples
```terraform
resource "azurerm_resource_group" "positive1" {
  name     = "acceptanceTestResourceGroup1"
  location = "West US"
}

resource "azurerm_sql_server" "positive2" {
  name                         = "mysqlserver1"
  resource_group_name          = "acceptanceTestResourceGroup1"
  location                     = "West US"
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

resource "azurerm_sql_active_directory_administrator" "positive3" {
  server_name         = "mysqlserver2"
  resource_group_name = "acceptanceTestResourceGroup1"
  login               = "sqladmin"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  object_id           = data.azurerm_client_config.current.object_id
}

```

## Compliant Code Examples
```terraform
resource "azurerm_resource_group" "negative1" {
  name     = "acceptanceTestResourceGroup1"
  location = "West US"
}

resource "azurerm_sql_server" "negative2" {
  name                         = "mysqlserver"
  resource_group_name          = "acceptanceTestResourceGroup1"
  location                     = "West US"
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

resource "azurerm_sql_active_directory_administrator" "negative3" {
  server_name         = "mysqlserver"
  resource_group_name = "acceptanceTestResourceGroup1"
  login               = "sqladmin"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  object_id           = data.azurerm_client_config.current.object_id
}
```
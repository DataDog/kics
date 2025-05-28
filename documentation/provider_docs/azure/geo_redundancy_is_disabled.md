---
title: "Geo Redundancy Is Disabled"
meta:
  name: "azure/geo_redundancy_is_disabled"
  id: "8b042c30-e441-453f-b162-7696982ebc58"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Backup"
---

## Metadata
**Name:** `azure/geo_redundancy_is_disabled`

**Id:** `8b042c30-e441-453f-b162-7696982ebc58`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Backup

## Description
Make sure that on PostgreSQL Geo Redundant Backups is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_server)

## Non-Compliant Code Examples
```terraform

resource "azurerm_postgresql_server" "positive1" {
    name                = "dbserver"
    location            = "usgovvirginia"
    resource_group_name = azurerm_resource_group.jira_rg.name

    sku_name   = "GP_Gen5_4"
    version    = "9.6"
    storage_mb = 640000

    backup_retention_days        = var.jira_postgre_data_retention
    auto_grow_enabled            = true

    administrator_login          = var.mp_db_username
    administrator_login_password = azurerm_key_vault_secret.db_pswd.value
    ssl_enforcement_enabled      = true

    tags                         = local.postgresqlserver_tags
}

resource "azurerm_postgresql_server" "positive2" {
    name                = "dbserver"
    location            = "usgovvirginia"
    resource_group_name = azurerm_resource_group.jira_rg.name

    sku_name   = "GP_Gen5_4"
    version    = "9.6"
    storage_mb = 640000

    backup_retention_days        = var.jira_postgre_data_retention
    geo_redundant_backup_enabled = false
    auto_grow_enabled            = true

    administrator_login          = var.mp_db_username
    administrator_login_password = azurerm_key_vault_secret.db_pswd.value
    ssl_enforcement_enabled      = false

    tags                         = local.postgresqlserver_tags
}

```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_server" "negative1" {
    name                = "dbserver"
    location            = "usgovvirginia"
    resource_group_name = azurerm_resource_group.jira_rg.name

    sku_name   = "GP_Gen5_4"
    version    = "9.6"
    storage_mb = 640000

    backup_retention_days        = var.jira_postgre_data_retention
    geo_redundant_backup_enabled = true
    auto_grow_enabled            = true

    administrator_login          = var.mp_db_username
    administrator_login_password = azurerm_key_vault_secret.db_pswd.value
    ssl_enforcement_enabled      = false

    tags                         = local.postgresqlserver_tags
}

```
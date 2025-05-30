---
title: "PostgreSQL Server Threat Detection Policy Disabled"
meta:
  name: "azure/postgresql_server_threat_detection_policy_disabled"
  id: "c407c3cf-c409-4b29-b590-db5f4138d332"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Resource Management"
---

## Metadata
**Name:** `azure/postgresql_server_threat_detection_policy_disabled`

**Id:** `c407c3cf-c409-4b29-b590-db5f4138d332`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Resource Management

## Description
PostgreSQL Server Threat Detection Policy should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_server#threat_detection_policy)

## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_server" "positive2" {
  name                = "example-psqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "psqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "GP_Gen5_4"
  version    = "9.6"
  storage_mb = 640000

  backup_retention_days        = 7
  geo_redundant_backup_enabled = true
  auto_grow_enabled            = true

  public_network_access_enabled    = false
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"
}

```

```terraform
resource "azurerm_postgresql_server" "positive1" {
  name                = "example-psqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "psqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "GP_Gen5_4"
  version    = "9.6"
  storage_mb = 640000

  backup_retention_days        = 7
  geo_redundant_backup_enabled = true
  auto_grow_enabled            = true

  public_network_access_enabled    = false
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"

  threat_detection_policy {
    enabled = false
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_server" "negative" {
  name                = "example-psqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "psqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "GP_Gen5_4"
  version    = "9.6"
  storage_mb = 640000

  backup_retention_days        = 7
  geo_redundant_backup_enabled = true
  auto_grow_enabled            = true

  public_network_access_enabled    = false
  ssl_enforcement_enabled          = true
  ssl_minimal_tls_version_enforced = "TLS1_2"

  threat_detection_policy {
    enabled = true
  }
}

```
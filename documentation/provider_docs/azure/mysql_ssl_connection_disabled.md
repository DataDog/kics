---
title: "MySQL SSL Connection Disabled"
meta:
  name: "azure/mysql_ssl_connection_disabled"
  id: "73e42469-3a86-4f39-ad78-098f325b4e9f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `azure/mysql_ssl_connection_disabled`

**Id:** `73e42469-3a86-4f39-ad78-098f325b4e9f`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Encryption

## Description
Make sure that for MySQL Database Server, 'Enforce SSL connection' is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mysql_server)

## Non-Compliant Code Examples
```terraform
resource "azurerm_mysql_server" "positive1" {
  name                = "webflux-mysql-${var.environment}${random_integer.rnd_int.result}"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "webflux-${var.environment}"
  administrator_login_password = random_string.password.result

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  infrastructure_encryption_enabled = true
  public_network_access_enabled     = true
  ssl_enforcement_enabled           = false
}

```

## Compliant Code Examples
```terraform
resource "azurerm_mysql_server" "negative1" {
  name                = "webflux-mysql-${var.environment}${random_integer.rnd_int.result}"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "webflux-${var.environment}"
  administrator_login_password = random_string.password.result

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  infrastructure_encryption_enabled = true
  public_network_access_enabled     = true
  ssl_enforcement_enabled           = true
}

```
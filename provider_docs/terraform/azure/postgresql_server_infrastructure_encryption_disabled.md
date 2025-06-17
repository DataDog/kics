---
title: "PostgreSQL Server Infrastructure Encryption Disabled"
meta:
  name: "terraform/postgresql_server_infrastructure_encryption_disabled"
  id: "6425c98b-ca4e-41fe-896a-c78772c131f8"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/postgresql_server_infrastructure_encryption_disabled`
**Id:** `6425c98b-ca4e-41fe-896a-c78772c131f8`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Encryption
## Description
PostgreSQL Server infrastructure encryption provides an additional layer of protection for data at rest, beyond the default storage encryption. If the `infrastructure_encryption_enabled` attribute is set to `false` in the Terraform `azurerm_postgresql_server` resource, sensitive data may be exposed in the event of unauthorized access to underlying disks or snapshots. To secure the configuration, set `infrastructure_encryption_enabled = true`.

This ensures that the PostgreSQL server uses Azure's strong encryption mechanisms to protect data at the infrastructure level, significantly reducing the risk of data compromise.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_server#infrastructure_encryption_enabled)

## Non-Compliant Code Examples
```azure
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

  infrastructure_encryption_enabled = false
}

```

```azure
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

## Compliant Code Examples
```azure
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

  infrastructure_encryption_enabled = true
}

```
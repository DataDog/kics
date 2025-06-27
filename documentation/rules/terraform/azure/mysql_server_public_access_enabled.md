---
title: "MySQL Server Public Access Enabled"
meta:
  name: "azure/mysql_server_public_access_enabled"
  id: "f118890b-2468-42b1-9ce9-af35146b425b"
  display_name: "MySQL Server Public Access Enabled"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Name:** `azure/mysql_server_public_access_enabled`

**Query Name** `MySQL Server Public Access Enabled`

**Id:** `f118890b-2468-42b1-9ce9-af35146b425b`

**Cloud Provider:** azure

**Platform** Terraform

**Severity:** High

**Category:** Networking and Firewall

## Description
This check verifies if the Azure MySQL Server has public network access enabled, which allows connections from the internet to reach your database. When public network access is enabled, your database is accessible to anyone who has the connection information, creating a significant security risk. Properly secured MySQL servers should have public_network_access_enabled set to false, forcing all connections to be made through private endpoints or service endpoints. Secure configuration example: 
```
resource "azurerm_mysql_server" "example" {
  // ... other configuration ...
  public_network_access_enabled = false
  // ... other configuration ...
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mysql_server#public_network_access_enabled)


## Compliant Code Examples
```terraform
resource "azurerm_mysql_server" "negative" {
  name                = "example-mysqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mysqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = false
  public_network_access_enabled     = false
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}

```
## Non-Compliant Code Examples
```terraform
resource "azurerm_mysql_server" "positive2" {
  name                = "example-mysqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mysqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = false
  public_network_access_enabled     = true
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}

```

```terraform
resource "azurerm_mysql_server" "positive1" {
  name                = "example-mysqlserver"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mysqladminun"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "5.7"

  auto_grow_enabled                 = true
  backup_retention_days             = 7
  geo_redundant_backup_enabled      = false
  infrastructure_encryption_enabled = false
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}

```
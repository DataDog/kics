---
title: "MariaDB Server Public Network Access Enabled"
meta:
  name: "azure/mariadb_public_network_access_enabled"
  id: "7f0a8696-7159-4337-ad0d-8a3ab4a78195"
  cloud_provider: "azure"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `azure/mariadb_public_network_access_enabled`
**Id:** `7f0a8696-7159-4337-ad0d-8a3ab4a78195`
**Cloud Provider:** azure
**Severity:** High
**Category:** Networking and Firewall
## Description
Azure MariaDB Server with public network access enabled allows direct connections from the internet, significantly increasing the risk of unauthorized access and potential data breaches. Attackers can attempt brute force attacks against your database credentials or exploit vulnerabilities if the server is publicly accessible. To properly secure your MariaDB Server, you should disable public network access by setting the 'public_network_access_enabled' attribute to 'false', as shown in the secure configuration example:

```
resource "azurerm_mariadb_server" "example" {
  // other configuration
  public_network_access_enabled = false
}
```

Instead, use private endpoints, service endpoints, or VPN connections to access your database securely.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mariadb_server#public_network_access_enabled)


## Compliant Code Examples
```terraform
resource "azurerm_mariadb_server" "negative" {
  name                = "example-mariadb-server"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mariadbadmin"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "10.2"

  auto_grow_enabled             = true
  backup_retention_days         = 7
  geo_redundant_backup_enabled  = false
  public_network_access_enabled = false
  ssl_enforcement_enabled       = true
}

```
## Non-Compliant Code Examples
```terraform
resource "azurerm_mariadb_server" "positive2" {
  name                = "example-mariadb-server"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mariadbadmin"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "10.2"

  auto_grow_enabled             = true
  backup_retention_days         = 7
  geo_redundant_backup_enabled  = false
  ssl_enforcement_enabled       = true
}

```

```terraform
resource "azurerm_mariadb_server" "positive" {
  name                = "example-mariadb-server"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  administrator_login          = "mariadbadmin"
  administrator_login_password = "H@Sh1CoR3!"

  sku_name   = "B_Gen5_2"
  storage_mb = 5120
  version    = "10.2"

  auto_grow_enabled             = true
  backup_retention_days         = 7
  geo_redundant_backup_enabled  = false
  public_network_access_enabled = true
  ssl_enforcement_enabled       = true
}

```
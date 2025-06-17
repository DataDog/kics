---
title: "'ssl_enforcement_enabled' is not set to 'ENABLED' for MySQL Database Server"
meta:
  name: "terraform/mysql_enforce_ssl_connection_disabled"
  id: "c3f2e1d0-b9a8-7c6d-5e4f-3210fedcba98"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/mysql_enforce_ssl_connection_disabled`
**Id:** `c3f2e1d0-b9a8-7c6d-5e4f-3210fedcba98`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Networking and Firewall
## Description
When SSL enforcement is disabled on Azure MySQL Database Servers, connections to the database are vulnerable to man-in-the-middle attacks and data interception. This security vulnerability allows attackers to potentially capture sensitive data transmitted between client applications and the database server, including credentials, personally identifiable information, and business-critical data. To secure your MySQL server, you must explicitly set 'ssl_enforcement_enabled' to 'ENABLED' as shown below:

```terraform
resource "azurerm_mysql_server" "good_example" {
  name                = "good-mysql-server"
  location            = "East US"
  resource_group_name = "example-rg"

  ssl_enforcement_enabled = ["ENABLED"]
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mysql_server)

## Non-Compliant Code Examples
```azure
resource "azurerm_mysql_server" "bad_example" {
  name                = "bad-mysql-server"
  location            = "East US"
  resource_group_name = "example-rg"

  ssl_enforcement_enabled = ["DISABLED"] # ❌ SSL enforcement is not enabled
}

```

## Compliant Code Examples
```azure
resource "azurerm_mysql_server" "good_example" {
  name                = "good-mysql-server"
  location            = "East US"
  resource_group_name = "example-rg"

  ssl_enforcement_enabled = ["ENABLED"] # ✅ Correct setting
}

```
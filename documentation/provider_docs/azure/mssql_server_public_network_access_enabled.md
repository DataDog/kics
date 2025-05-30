---
title: "MSSQL Server Public Network Access Enabled"
meta:
  name: "azure/mssql_server_public_network_access_enabled"
  id: "ade36cf4-329f-4830-a83d-9db72c800507"
  cloud_provider: "azure"
  severity: "HIGH"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/mssql_server_public_network_access_enabled`

**Id:** `ade36cf4-329f-4830-a83d-9db72c800507`

**Cloud Provider:** azure

**Severity:** High

**Category:** Networking and Firewall

## Description
When MSSQL Server public network access is enabled, it allows connections from the internet to your database server, significantly expanding the attack surface and potentially exposing it to unauthorized access. This vulnerability could lead to data breaches, unauthorized data manipulation, or denial of service attacks if credentials are compromised or if there are exploitable vulnerabilities in the database server. To mitigate this risk, set `public_network_access_enabled = false` in your MSSQL Server configuration, which restricts access to private endpoints or services within your Azure network only.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mssql_server#public_network_access_enabled)

## Non-Compliant Code Examples
```terraform
resource "azurerm_mssql_server" "positive2" {
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

    public_network_access_enabled = true
}

```

```terraform
resource "azurerm_mssql_server" "positive1" {
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

## Compliant Code Examples
```terraform
resource "azurerm_mssql_server" "negative1" {
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

    public_network_access_enabled = false
}

```
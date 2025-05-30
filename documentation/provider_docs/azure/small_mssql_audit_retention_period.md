---
title: "Small MSSQL Audit Retention Period"
meta:
  name: "azure/small_mssql_audit_retention_period"
  id: "9c301481-e6ec-44f7-8a49-8ec63e2969ea"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `azure/small_mssql_audit_retention_period`

**Id:** `9c301481-e6ec-44f7-8a49-8ec63e2969ea`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Observability

## Description
Make sure that for MSSQL Server, the Auditing Retention is greater than 90 days

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/mssql_server)

## Non-Compliant Code Examples
```terraform
resource "azurerm_mssql_database" "positive1" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_mssql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 6
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_database" "positive2" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_mssql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 90
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_database" "positive3" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_mssql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 0
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_server" "positive4" {
    name                         = "mssqlserver"
    resource_group_name          = azurerm_resource_group.example.name
    location                     = azurerm_resource_group.example.location
    version                      = "12.0"
    administrator_login          = "mradministrator"
    administrator_login_password = "thisIsDog11"

    extended_auditing_policy {
      storage_endpoint            = azurerm_storage_account.example.primary_blob_endpoint
      storage_account_access_key  = azurerm_storage_account.example.primary_access_key
      storage_account_access_key_is_secondary = true
      retention_in_days                       = 20
    }
}
```

## Compliant Code Examples
```terraform
resource "azurerm_mssql_database" "negative1" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_sql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 91
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_database" "negative2" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_sql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 214
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_database" "negative3" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_sql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 30000
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_database" "negative4" {
  name                = "myexamplesqldatabase"
  resource_group_name = azurerm_resource_group.example.name
  location            = "West US"
  server_name         = azurerm_sql_server.example.name

  extended_auditing_policy {
    storage_endpoint                        = azurerm_storage_account.example.primary_blob_endpoint
    storage_account_access_key              = azurerm_storage_account.example.primary_access_key
    storage_account_access_key_is_secondary = true
    retention_in_days                       = 900
  }

  tags = {
    environment = "production"
  }
}

resource "azurerm_mssql_server" "negative5" {
    name                         = "mssqlserver"
    resource_group_name          = azurerm_resource_group.example.name
    location                     = azurerm_resource_group.example.location
    version                      = "12.0"
    administrator_login          = "mradministrator"
    administrator_login_password = "thisIsDog11"

    extended_auditing_policy {
      storage_endpoint            = azurerm_storage_account.example.primary_blob_endpoint
      storage_account_access_key  = azurerm_storage_account.example.primary_access_key
      storage_account_access_key_is_secondary = true
      retention_in_days                       = 95
    }
}
```
---
title: "Function App FTPS Enforce Disabled"
meta:
  name: "azure/function_app_ftps_enforce_disabled"
  id: "9dab0179-433d-4dff-af8f-0091025691df"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/function_app_ftps_enforce_disabled`

**Id:** `9dab0179-433d-4dff-af8f-0091025691df`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Azure Function App should only enforce FTPS when 'ftps_state' is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#ftps_state)

## Non-Compliant Code Examples
```terraform
resource "azurerm_function_app" "positive2" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   site_config {
    http2_enabled = true
  }
}

```

```terraform
resource "azurerm_function_app" "positive1" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   site_config {
    http2_enabled = true
    ftps_state = "AllAllowed"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_function_app" "negative2" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   site_config {
    ftps_state = "Disabled"
  }
}

```

```terraform
resource "azurerm_function_app" "negative1" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   site_config {
    ftps_state = "FtpsOnly"
  }
}

```
---
title: "Function App Authentication Disabled"
meta:
  name: "azure/function_app_authentication_disabled"
  id: "e65a0733-94a0-4826-82f4-df529f4c593f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/function_app_authentication_disabled`

**Id:** `e65a0733-94a0-4826-82f4-df529f4c593f`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Azure Function App authentication settings should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#auth_settings)

## Non-Compliant Code Examples
```terraform
resource "azurerm_function_app" "positive2" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   auth_settings {
    enabled = false
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
}

```

## Compliant Code Examples
```terraform
resource "azurerm_function_app" "negative" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

   auth_settings {
    enabled = true
  }
}

```
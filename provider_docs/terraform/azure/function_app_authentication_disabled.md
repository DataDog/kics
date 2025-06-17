---
title: "Function App Authentication Disabled"
meta:
  name: "terraform/function_app_authentication_disabled"
  id: "e65a0733-94a0-4826-82f4-df529f4c593f"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/function_app_authentication_disabled`
**Id:** `e65a0733-94a0-4826-82f4-df529f4c593f`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Azure Function App authentication settings should be enabled to ensure that only authorized users and services can access the deployed function endpoints. Leaving authentication (`auth_settings { enabled = true }`) disabled, as shown below, permits unauthenticated, potentially malicious access to the function app, increasing the risk of data exposure or unauthorized actions:

```
resource "azurerm_function_app" "example" {
  // ...other config...
  auth_settings {
    enabled = true
  }
}
```

Enabling authentication protects sensitive operations and helps prevent unauthorized or anonymous interaction with serverless workloads.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#auth_settings)

## Non-Compliant Code Examples
```azure
resource "azurerm_function_app" "positive1" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
}

```

```azure
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

## Compliant Code Examples
```azure
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
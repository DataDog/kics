---
title: "Function App Managed Identity Disabled"
meta:
  name: "azure/function_app_managed_identity_disabled"
  id: "c87749b3-ff10-41f5-9df2-c421e8151759"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/function_app_managed_identity_disabled`

**Id:** `c87749b3-ff10-41f5-9df2-c421e8151759`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Azure Function App should have managed identity enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#identity)

## Non-Compliant Code Examples
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

  identity {
    type = "SystemAssigned"
  }
}

```
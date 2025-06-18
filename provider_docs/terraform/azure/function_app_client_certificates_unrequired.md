---
title: "Function App Client Certificates Unrequired"
meta:
  name: "terraform/function_app_client_certificates_unrequired"
  id: "9bb3c639-5edf-458c-8ee5-30c17c7d671d"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/function_app_client_certificates_unrequired`
**Id:** `9bb3c639-5edf-458c-8ee5-30c17c7d671d`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Azure Function Apps should require client certificates for incoming requests by setting the `client_cert_mode` attribute to `"Required"`. Without this setting, as seen below, the Function App allows unauthenticated traffic, increasing the risk of unauthorized access to sensitive business logic or data processed by the Function App:

```
resource "azurerm_function_app" "example" {
  // ... other configuration ...
  client_cert_mode = "Required"
}
```

Enforcing client certificate authentication ensures that only trusted clients can interact with the Function App, reducing the attack surface and protecting against various unauthorized access vectors.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#client_cert_mode)

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

  client_cert_mode = "Optional"
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

  client_cert_mode = "Required"
}

```
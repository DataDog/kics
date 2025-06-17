---
title: "Function App HTTP2 Disabled"
meta:
  name: "terraform/function_app_http2_disabled"
  id: "ace823d1-4432-4dee-945b-cdf11a5a6bd0"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/function_app_http2_disabled`
**Id:** `ace823d1-4432-4dee-945b-cdf11a5a6bd0`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Enabling HTTP/2 for Azure Function Apps improves security and performance by providing better encryption and more efficient use of network resources compared to HTTP/1.1. If the `http2_enabled` attribute is not set to `true` in the `site_config` block, as shown below, the Function App defaults to HTTP/1.1, exposing it to potential vulnerabilities like protocol downgrade attacks and inefficient data transfer:

```
site_config {
  http2_enabled = true
}
```

Leaving HTTP/2 disabled can result in degraded application security and responsiveness, making services more vulnerable to exploitation and less performant for end users.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#http2_enabled)

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
resource "azurerm_function_app" "positive3" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
    http2_enabled = false
  }
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

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
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

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
    http2_enabled = true
  }
}

```
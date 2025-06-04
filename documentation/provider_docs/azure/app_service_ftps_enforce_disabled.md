---
title: "App Service FTPS Enforce Disabled"
meta:
  name: "azure/app_service_ftps_enforce_disabled"
  id: "85da374f-b00f-4832-9d44-84a1ca1e89f8"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/app_service_ftps_enforce_disabled`

**Id:** `85da374f-b00f-4832-9d44-84a1ca1e89f8`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Azure App Service should only enforce FTPS when 'ftps_state' is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#ftps_state)

## Non-Compliant Code Examples
```terraform
resource "azurerm_app_service" "positive1" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    ftps_state = "AllAllowed"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_app_service" "negative2" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    ftps_state = "Disabled"
  }
}

```

```terraform
resource "azurerm_app_service" "negative1" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    ftps_state = "FtpsOnly"
  }
}

```
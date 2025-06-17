---
title: "App Service FTPS Enforce Disabled"
meta:
  name: "terraform/app_service_ftps_enforce_disabled"
  id: "85da374f-b00f-4832-9d44-84a1ca1e89f8"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/app_service_ftps_enforce_disabled`
**Id:** `85da374f-b00f-4832-9d44-84a1ca1e89f8`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
App Service FTPS enforcement should be configured by setting the `ftps_state` attribute to `"FtpsOnly"` in the `site_config` block for the `azurerm_app_service` resource. Allowing `"AllAllowed"` in this setting permits both unencrypted FTP and encrypted FTPS connections, exposing sensitive data to potential interception or tampering during transit. To ensure data confidentiality and compliance, always use:

```
ftps_state = "FtpsOnly"
```
which enforces encrypted connections to the Azure App Service.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#ftps_state)

## Non-Compliant Code Examples
```azure
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
```azure
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

```azure
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
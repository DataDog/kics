---
title: "App Service Authentication Disabled"
meta:
  name: "azure/app_service_authentication_disabled"
  id: "c7fc1481-2899-4490-bbd8-544a3a61a2f3"
  display_name: "App Service Authentication Disabled"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Name:** `azure/app_service_authentication_disabled`

**Query Name** `App Service Authentication Disabled`

**Id:** `c7fc1481-2899-4490-bbd8-544a3a61a2f3`

**Cloud Provider:** azure

**Platform** Terraform

**Severity:** Medium

**Category:** Access Control

## Description
Azure App Service authentication settings should be enabled to ensure that access to the application is restricted to authenticated users. Without enabling the `auth_settings { enabled = true }` block in Terraform, anyone can anonymously access the app, which exposes it to unauthorized access and potential misuse of sensitive data or resources. Properly configuring authentication securely helps mitigate risks from attacks such as data exfiltration, account takeover, or service abuse.

```
auth_settings = {
  enabled = true
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#enabled)


## Compliant Code Examples
```terraform
resource "azurerm_app_service" "negative1" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  auth_settings = {
    enabled = true
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "azurerm_app_service" "positive2" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  auth_settings = {
    enabled = false
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}

```

```terraform
resource "azurerm_app_service" "positive1" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}

```
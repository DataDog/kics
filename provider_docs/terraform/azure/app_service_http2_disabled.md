---
title: "App Service HTTP2 Disabled"
meta:
  name: "terraform/app_service_http2_disabled"
  id: "525b53be-62ed-4244-b4df-41aecfcb4071"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/app_service_http2_disabled`
**Id:** `525b53be-62ed-4244-b4df-41aecfcb4071`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Enabling HTTP/2 for Azure App Service is essential for improved security and performance, as HTTP/2 offers better data encryption, reduced latency, and protection against certain protocol-level attacks. If the `http2_enabled` attribute is not set to `true` in the `site_config` block, as shown below, the app service will only support HTTP/1.1, making it more vulnerable to downgrade attacks and less efficient in handling modern web traffic.

```
site_config {
  http2_enabled = true
}
```
Neglecting to enable HTTP/2 may expose applications to increased risks and degrade the overall performance and user experience.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#http2_enabled)

## Non-Compliant Code Examples
```azure
resource "azurerm_app_service" "positive1" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

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

```azure
resource "azurerm_app_service" "positive3" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
    http2_enabled = false
  }
}

```

```azure
resource "azurerm_app_service" "positive2" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
  }
}

```

## Compliant Code Examples
```azure
resource "azurerm_app_service" "negative" {
  name                = "example-app-service"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  app_service_plan_id = azurerm_app_service_plan.example.id

  app_settings = {
    "SOME_KEY" = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
    http2_enabled = true
  }
}

```
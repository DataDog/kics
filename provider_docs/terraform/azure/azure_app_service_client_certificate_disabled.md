---
title: "Azure App Service Client Certificate Disabled"
meta:
  name: "terraform/azure_app_service_client_certificate_disabled"
  id: "a81573f9-3691-4d83-88a0-7d4af63e17a3"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/azure_app_service_client_certificate_disabled`
**Id:** `a81573f9-3691-4d83-88a0-7d4af63e17a3`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Enabling client certificates for Azure App Service ensures that only authenticated clients can access the application by requiring a client SSL certificate in all HTTPS requests. If the `client_cert_enabled` attribute is not set to `true`, unauthorized users could potentially connect to the service, increasing the risk of data leaks or abuse. To secure the App Service, always configure the resource as follows:

```
resource "azurerm_app_service" "example" {
  // ... other configuration ...
  client_cert_enabled = true
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/app_service#client_cert_enabled)

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
  }

  app_settings = {
    SOME_KEY = "some-value"
  }

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}

```

```azure
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
    SOME_KEY = "some-value"
  }

  client_cert_enabled = false

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
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

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }

  app_settings = {
    SOME_KEY = "some-value"
  }

  client_cert_enabled = true

  connection_string {
    name  = "Database"
    type  = "SQLServer"
    value = "Server=some-server.mydomain.com;Integrated Security=SSPI"
  }
}


```
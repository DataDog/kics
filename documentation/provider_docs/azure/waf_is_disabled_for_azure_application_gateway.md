---
title: "WAF Is Disabled For Azure Application Gateway"
meta:
  name: "azure/waf_is_disabled_for_azure_application_gateway"
  id: "2e48d91c-50e4-45c8-9312-27b625868a72"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/waf_is_disabled_for_azure_application_gateway`

**Id:** `2e48d91c-50e4-45c8-9312-27b625868a72`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Check if Web Application Firewall is disabled or not configured for Azure's Application Gateway.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/application_gateway)

## Non-Compliant Code Examples
```terraform
resource "azurerm_application_gateway" "positive1" {
  name                = "example-appgateway"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location

  waf_configuration {
    enabled = false
  }
}

resource "azurerm_application_gateway" "positive2" {
  name                = "example-appgateway"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
}
```

## Compliant Code Examples
```terraform
resource "azurerm_application_gateway" "negative1" {
  name                = "example-appgateway"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location

  waf_configuration {
    enabled = true
  }
}
```
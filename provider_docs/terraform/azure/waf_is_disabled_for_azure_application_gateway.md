---
title: "WAF Is Disabled For Azure Application Gateway"
meta:
  name: "terraform/waf_is_disabled_for_azure_application_gateway"
  id: "2e48d91c-50e4-45c8-9312-27b625868a72"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/waf_is_disabled_for_azure_application_gateway`
**Id:** `2e48d91c-50e4-45c8-9312-27b625868a72`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
This check ensures that the Azure Application Gateway has its Web Application Firewall (WAF) correctly configured and enabled, as indicated by the `waf_configuration { enabled = true }` attribute in Terraform. If WAF is not enabled or omitted from the configuration, the application gateway is left unprotected against common web attacks, such as SQL injection and cross-site scripting, increasing the risk of a successful attack. To mitigate this vulnerability, always configure WAF with `enabled = true` as shown below:

```
resource "azurerm_application_gateway" "example" {
  // ... other settings ...
  waf_configuration {
    enabled = true
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/application_gateway)

## Non-Compliant Code Examples
```azure
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
```azure
resource "azurerm_application_gateway" "negative1" {
  name                = "example-appgateway"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location

  waf_configuration {
    enabled = true
  }
}
```
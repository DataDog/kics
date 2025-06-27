---
title: "Security Center Pricing Tier Is Not Standard"
meta:
  name: "azure/security_center_pricing_tier_is_not_standard"
  id: "819d50fd-1cdf-45c3-9936-be408aaad93e"
  display_name: "Security Center Pricing Tier Is Not Standard"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Name:** `azure/security_center_pricing_tier_is_not_standard`

**Query Name** `Security Center Pricing Tier Is Not Standard`

**Id:** `819d50fd-1cdf-45c3-9936-be408aaad93e`

**Cloud Provider:** azure

**Platform** Terraform

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Selecting the appropriate pricing tier for Azure Security Center is crucial for ensuring comprehensive security monitoring and protection. If the `tier` attribute in the `azurerm_security_center_subscription_pricing` resource is set to `"Free"`, as shown below:

```
resource "azurerm_security_center_subscription_pricing" "example" {
   tier = "Free"
}
```

the configuration provides only limited security features and monitoring capabilities, leaving the environment more vulnerable to threats. Upgrading to the `"Standard"` tier, as recommended:

```
resource "azurerm_security_center_subscription_pricing" "example" {
   tier = "Standard"
}
```

enables advanced security features such as threat detection and automated response, significantly reducing the risk of undetected attacks and data breaches.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/security_center_subscription_pricing)


## Compliant Code Examples
```terraform
resource "azurerm_security_center_subscription_pricing" "negative1" {
   tier = "Standard"
}
```
## Non-Compliant Code Examples
```terraform
resource "azurerm_security_center_subscription_pricing" "positive1" {
   tier = "Free"
}
```
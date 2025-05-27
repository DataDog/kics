---
title: "Security Center Pricing Tier Is Not Standard"
meta:
  name: "azure/security_center_pricing_tier_is_not_standard"
  id: "819d50fd-1cdf-45c3-9936-be408aaad93e"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/security_center_pricing_tier_is_not_standard`

**Id:** `819d50fd-1cdf-45c3-9936-be408aaad93e`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Make sure that the 'Standard' pricing tiers were selected.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/security_center_subscription_pricing)

## Non-Compliant Code Examples
```terraform
resource "azurerm_security_center_subscription_pricing" "positive1" {
   tier = "Free"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_security_center_subscription_pricing" "negative1" {
   tier = "Standard"
}
```
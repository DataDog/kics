---
title: "Email Alerts Disabled"
meta:
  name: "azure/email_alerts_disabled"
  id: "9db38e87-f6aa-4b5e-a1ec-7266df259409"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/email_alerts_disabled`

**Id:** `9db38e87-f6aa-4b5e-a1ec-7266df259409`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that alerts notifications are set to 'On' in the Azure Security Center Contact

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/security_center_contact)

## Non-Compliant Code Examples
```terraform
resource "azurerm_security_center_contact" "positive1" {
    email = "contact@example.com"
    phone = "+1-555-555-5555"
   alert_notifications = false
}
```

## Compliant Code Examples
```terraform
resource "azurerm_security_center_contact" "negative1" {
    email = "contact@example.com"
    phone = "+1-555-555-5555"
    alert_notifications = true
}
```
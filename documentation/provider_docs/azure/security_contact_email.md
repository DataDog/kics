---
title: "Security Contact Email"
meta:
  name: "azure/security_contact_email"
  id: "34664094-59e0-4524-b69f-deaa1a68cce3"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/security_contact_email`

**Id:** `34664094-59e0-4524-b69f-deaa1a68cce3`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Best Practices

## Description
Security Contact Email should be defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/security_center_contact#email)

## Non-Compliant Code Examples
```terraform
resource "azurerm_security_center_contact" "positive" {
  phone = "+1-555-555-5555"

  alert_notifications = true
  alerts_to_admins    = true
}

```

## Compliant Code Examples
```terraform
resource "azurerm_security_center_contact" "negative" {
  email = "contact@example.com"
  phone = "+1-555-555-5555"

  alert_notifications = true
  alerts_to_admins    = true
}

```
---
title: "Secret Expiration Not Set"
meta:
  name: "azure/secret_expiration_not_set"
  id: "dfa20ffa-f476-428f-a490-424b41e91c7f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Secret Management"
---

## Metadata
**Name:** `azure/secret_expiration_not_set`

**Id:** `dfa20ffa-f476-428f-a490-424b41e91c7f`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Secret Management

## Description
Make sure that for all secrets the expiration date is set

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_secret)

## Non-Compliant Code Examples
```terraform
resource "azurerm_key_vault_secret" "positive1" {
    name         = "secret-sauce"
    value        = "szechuan"
    key_vault_id = azurerm_key_vault.example.id

    tags = {
    environment = "Production"
    }
}
```

## Compliant Code Examples
```terraform
resource "azurerm_key_vault_secret" "negative1" {
    name         = "secret-sauce"
    value        = "szechuan"
    key_vault_id = azurerm_key_vault.example.id

    tags = {
    environment = "Production"
    }
    expiration_date = "2020-12-30T20:00:00Z"
}
```
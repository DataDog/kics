---
title: "Key Vault Secrets Content Type Undefined"
meta:
  name: "azure/key_vault_secrets_content_type_undefined"
  id: "f8e08a38-fc6e-4915-abbe-a7aadf1d59ef"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/key_vault_secrets_content_type_undefined`

**Id:** `f8e08a38-fc6e-4915-abbe-a7aadf1d59ef`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Best Practices

## Description
Key Vault Secrets should have set Content Type

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_secret#content_type)

## Non-Compliant Code Examples
```terraform
resource "azurerm_key_vault_secret" "positive" {
  name         = "secret-sauce"
  value        = "szechuan"
  key_vault_id = azurerm_key_vault.example.id
}

```

## Compliant Code Examples
```terraform
resource "azurerm_key_vault_secret" "negative" {
  name         = "secret-sauce"
  value        = "szechuan"
  key_vault_id = azurerm_key_vault.example.id
  content_type = "password"
}

```
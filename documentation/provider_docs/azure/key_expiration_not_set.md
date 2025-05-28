---
title: "Key Expiration Not Set"
meta:
  name: "azure/key_expiration_not_set"
  id: "4d080822-5ee2-49a4-8984-68f3d4c890fc"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Secret Management"
---

## Metadata
**Name:** `azure/key_expiration_not_set`

**Id:** `4d080822-5ee2-49a4-8984-68f3d4c890fc`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Secret Management

## Description
Make sure that for all keys the expiration date is set

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/key_vault_key)

## Non-Compliant Code Examples
```terraform
resource "azurerm_key_vault_key" "positive1" {
    name         = "generated-certificate"
    key_vault_id = azurerm_key_vault.example.id
    key_type     = "RSA"
    key_size     = 2048

    key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
    ]
}
```

## Compliant Code Examples
```terraform
resource "azurerm_key_vault_key" "negative1" {
    name         = "generated-certificate"
    key_vault_id = azurerm_key_vault.example.id
    key_type     = "RSA"
    key_size     = 2048

    key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
    ]
  expiration_date = "2020-12-30T20:00:00Z"
}
```
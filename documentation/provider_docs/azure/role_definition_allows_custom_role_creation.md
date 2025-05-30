---
title: "Role Definition Allows Custom Role Creation"
meta:
  name: "azure/role_definition_allows_custom_role_creation"
  id: "3fa5900f-9aac-4982-96b2-a6143d9c99fb"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/role_definition_allows_custom_role_creation`

**Id:** `3fa5900f-9aac-4982-96b2-a6143d9c99fb`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Role Definition should not allow custom role creation (Microsoft.Authorization/roleDefinitions/write)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/role_definition#actions)

## Non-Compliant Code Examples
```terraform
resource "azurerm_role_definition" "example" {
  name        = "my-custom-role"
  scope       = data.azurerm_subscription.primary.id
  description = "This is a custom role created via Terraform"

  permissions {
    actions     = ["*"]
    not_actions = []
  }
}

```

```terraform
resource "azurerm_role_definition" "example2" {
  role_definition_id = "00000000-0000-0000-0000-000000000000"
  name               = "my-custom-role-definition"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions     = ["Microsoft.Authorization/roleDefinitions/write"]
    not_actions = []
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_role_definition" "example3" {
  role_definition_id = "00000000-0000-0000-0000-000000000000"
  name               = "my-custom-role-definition"
  scope              = data.azurerm_subscription.primary.id

  permissions {
    actions     = ["Microsoft.Authorization/roleDefinitions/read"]
    not_actions = []
  }
}

```
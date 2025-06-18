---
title: "Role Definition Allows Custom Role Creation"
meta:
  name: "terraform/role_definition_allows_custom_role_creation"
  id: "3fa5900f-9aac-4982-96b2-a6143d9c99fb"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/role_definition_allows_custom_role_creation`
**Id:** `3fa5900f-9aac-4982-96b2-a6143d9c99fb`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing the `Microsoft.Authorization/roleDefinitions/write` action in a custom Azure role definition grants users the ability to create, modify, or delete role definitions within the assigned scope. This level of permission can lead to privilege escalation, where a malicious or compromised user can create overly permissive roles or alter existing ones. To mitigate this risk, custom roles should be limited to `Microsoft.Authorization/roleDefinitions/read`, as shown below:

```
permissions {
  actions = ["Microsoft.Authorization/roleDefinitions/read"]
  not_actions = []
}
```

Restricting write access helps protect against unauthorized changes to role definitions and reduces the attack surface for privilege escalation.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/role_definition#actions)

## Non-Compliant Code Examples
```azure
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

```azure
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

## Compliant Code Examples
```azure
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
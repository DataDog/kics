---
title: "Role Assignment Not Limit Guest User Permissions"
meta:
  name: "azure/role_assignment_not_limit_guest_users_permissions"
  id: "8e75e431-449f-49e9-b56a-c8f1378025cf"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/role_assignment_not_limit_guest_users_permissions`

**Id:** `8e75e431-449f-49e9-b56a-c8f1378025cf`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Role Assignment should limit guest user permissions

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/role_assignment)

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

  assignable_scopes = [
    data.azurerm_subscription.primary.id, 
  ]
}

resource "azurerm_role_assignment" "example" {
  name               = "00000000-0000-0000-0000-000000000000"
  scope              = data.azurerm_subscription.primary.id
  role_definition_name = "Guest"
  role_definition_id = azurerm_role_definition.example.role_definition_resource_id
  principal_id       = data.azurerm_client_config.example.object_id
}

```

## Compliant Code Examples
```terraform
resource "azurerm_role_definition" "example2" {
  name        = "my-custom-role"
  scope       = data.azurerm_subscription.primary.id
  description = "This is a custom role created via Terraform"

  permissions {
    actions     = []
    not_actions = ["*"]
  }

  assignable_scopes = [
    data.azurerm_subscription.primary.id, 
  ]
}

resource "azurerm_role_assignment" "example2" {
  name               = "00000000-0000-0000-0000-000000000000"
  scope              = data.azurerm_subscription.primary.id
  role_definition_name = "Guest"
  role_definition_id = azurerm_role_definition.example2.role_definition_resource_id
  principal_id       = data.azurerm_client_config.example.object_id
}

```
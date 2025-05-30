---
title: "Admin User Enabled For Container Registry"
meta:
  name: "azure/admin_user_enabled_for_container_registry"
  id: "b897dfbf-322c-45a8-b67c-1e698beeaa51"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/admin_user_enabled_for_container_registry`

**Id:** `b897dfbf-322c-45a8-b67c-1e698beeaa51`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Admin user is enabled for Container Registry

#### Learn More

 - [Provider Reference](https://www.terraform.io/docs/providers/azurerm/r/container_registry.html)

## Non-Compliant Code Examples
```terraform
resource "azurerm_resource_group" "positive1" {
  name     = "resourceGroup1"
  location = "West US"
}

resource "azurerm_container_registry" "positive2" {
  name                     = "containerRegistry1"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  sku                      = "Premium"
  admin_enabled            = true
  georeplication_locations = ["East US", "West Europe"]
}
```

## Compliant Code Examples
```terraform
resource "azurerm_resource_group" "negative1" {
  name     = "resourceGroup1"
  location = "West US"
}

resource "azurerm_container_registry" "negative2" {
  name                     = "containerRegistry1"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  sku                      = "Premium"
  admin_enabled            = false
  georeplication_locations = ["East US", "West Europe"]
}
```
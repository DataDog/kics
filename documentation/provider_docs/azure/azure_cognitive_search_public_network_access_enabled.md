---
title: "Azure Cognitive Search Public Network Access Enabled"
meta:
  name: "azure/azure_cognitive_search_public_network_access_enabled"
  id: "4a9e0f00-0765-4f72-a0d4-d31110b78279"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/azure_cognitive_search_public_network_access_enabled`

**Id:** `4a9e0f00-0765-4f72-a0d4-d31110b78279`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Public Network Access should be disabled for Azure Cognitive Search

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/search_service#public_network_access_enabled)

## Non-Compliant Code Examples
```terraform
resource "azurerm_search_service" "positive2" {
  name                = "example-search-service"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku                 = "standard"
}

```

```terraform
resource "azurerm_search_service" "positive1" {
  name                = "example-search-service"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku                 = "standard"
  public_network_access_enabled = true
}

```

## Compliant Code Examples
```terraform
resource "azurerm_search_service" "example" {
  name                = "example-search-service"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  sku                 = "standard"
  public_network_access_enabled = false
}

```
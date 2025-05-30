---
title: "Cosmos DB Account Without Tags"
meta:
  name: "azure/cosmos_db_account_without_tags"
  id: "56dad03e-e94f-4dd6-93a4-c253a03ff7a0"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Build Process"
---

## Metadata
**Name:** `azure/cosmos_db_account_without_tags`

**Id:** `56dad03e-e94f-4dd6-93a4-c253a03ff7a0`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Build Process

## Description
Cosmos DB Account must have a mapping of tags.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/cosmosdb_account)

## Non-Compliant Code Examples
```terraform
resource "azurerm_cosmosdb_account" "positive1" {
  name                = "tfex-cosmos-db-${random_integer.ri.result}"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  offer_type          = "Standard"
  kind                = "GlobalDocumentDB"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_cosmosdb_account" "negative1" {
  name                = "tfex-cosmos-db-${random_integer.ri.result}"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  offer_type          = "Standard"
  kind                = "GlobalDocumentDB"
  tags                = "tag_1"
}
```
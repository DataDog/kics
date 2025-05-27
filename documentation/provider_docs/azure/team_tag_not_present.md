---
title: "Ensure that cloud resource has a 'team' tag"
meta:
  name: "azure/team_tag_not_present"
  id: "e8f4d3c2-b1a0-4e5f-8d7c-9a0b1c2d3e4f"
  cloud_provider: "azure"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/team_tag_not_present`

**Id:** `e8f4d3c2-b1a0-4e5f-8d7c-9a0b1c2d3e4f`

**Cloud Provider:** azure

**Severity:** Info

**Category:** Best Practices

## Description
All Azure resources must include a 'team' tag in their tags block to ensure proper ownership and accountability.

#### Learn More

 - [Provider Reference](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/tag-resources)

## Non-Compliant Code Examples
```terraform
# Example 1: Missing tags block entirely
resource "azurerm_storage_account" "bad_example_no_tags" {
  name                     = "badstorageacct"
  resource_group_name      = "example-rg"
  location                 = "East US"
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

# Example 2: Tags block exists, but missing the "team" tag
resource "azurerm_storage_account" "bad_example_missing_team" {
  name                     = "badstorageacct2"
  resource_group_name      = "example-rg"
  location                 = "East US"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags = {
    environment = "prod"
  }
}

```

## Compliant Code Examples
```terraform
# ✅ "team" label is not a valid attribute for this resource type

resource "azurerm_postgresql_test" "good_example" {
  name                = "good-postgresql-server"
  location            = "East US"
  resource_group_name = "example-rg"

  public_network_access_enabled = [false]

  version                 = "9.6"
  ssl_enforcement_enabled = true
  sku_name                = "B_Gen5_1"
}

```

```terraform
resource "azurerm_postgresql_server" "good_example" {
  name                = "good-postgresql-server"
  location            = "East US"
  resource_group_name = "example-rg"

  public_network_access_enabled = [false]

  version                 = "9.6"
  ssl_enforcement_enabled = true
  sku_name                = "B_Gen5_1"

  tags = {
    Team        = "DevOps" # ✅ Correct setting
    environment = "prod"
  }
}

```

```terraform
resource "azurerm_storage_account" "good_example" {
  name                     = "goodstorageacct"
  resource_group_name      = "example-rg"
  location                 = "East US"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags = {
    team        = "DevOps" # Required tag is present
    environment = "prod"
  }
}

```
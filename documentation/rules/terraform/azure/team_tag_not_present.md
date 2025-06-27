---
title: "Ensure that cloud resource has a 'team' tag"
meta:
  name: "azure/team_tag_not_present"
  id: "e8f4d3c2-b1a0-4e5f-8d7c-9a0b1c2d3e4f"
  display_name: "Ensure that cloud resource has a 'team' tag"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "INFO"
  category: "Best Practices"
---
## Metadata
**Name:** `azure/team_tag_not_present`
**Query Name** `Ensure that cloud resource has a 'team' tag`
**Id:** `e8f4d3c2-b1a0-4e5f-8d7c-9a0b1c2d3e4f`
**Cloud Provider:** azure
**Platform** Terraform
**Severity:** Info
**Category:** Best Practices
## Description
To ensure proper resource ownership and management accountability in Azure environments, all resources should include a `team` tag within their `tags` block. Without this tag, as shown below, it becomes difficult to identify who is responsible for the resource, increasing the risk of unmanaged assets, security oversights, and operational inefficiencies:

```
resource "azurerm_storage_account" "bad_example_no_tags" {
  name                     = "badstorageacct"
  resource_group_name      = "example-rg"
  location                 = "East US"
  account_tier             = "Standard"
  account_replication_type = "LRS"
}
```

Properly tagging resources with a `team` value, such as in this example, helps organizations implement cost controls, streamline incident response, and maintain compliance by ensuring every asset has clear ownership:

```
resource "azurerm_storage_account" "good_example" {
  name                     = "goodstorageacct"
  resource_group_name      = "example-rg"
  location                 = "East US"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  tags = {
    team = "DevOps"
  }
}
```

#### Learn More

 - [Provider Reference](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/tag-resources)


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
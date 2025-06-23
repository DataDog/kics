---
title: "Trusted Microsoft Services Not Enabled"
meta:
  name: "azure/trusted_microsoft_services_not_enabled"
  id: "5400f379-a347-4bdd-a032-446465fdcc6f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `azure/trusted_microsoft_services_not_enabled`
**Id:** `5400f379-a347-4bdd-a032-446465fdcc6f`
**Cloud Provider:** azure
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Trusted Microsoft Services should be enabled for Storage Account access to ensure that Azure resources such as Azure Backup, Azure Monitor, and others can securely interact with the Storage Account without exposing it more broadly. When the `bypass` attribute in `azurerm_storage_account` or `azurerm_storage_account_network_rules` is not set to include `"AzureServices"`, essential Azure services may be denied access or administrators may set overly permissive network rules, increasing the attack surface. Securely configuring your storage account should look like:

```
network_rules {
  default_action = "Deny"
  bypass         = ["AzureServices"]
}
```

Failing to enable Trusted Microsoft Services can hinder platform functionality or lead to weaker network restrictions that unnecessarily expose the storage account to risk.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_account#bypass)


## Compliant Code Examples
```terraform
resource "azurerm_storage_account" "negative1" {
  name                = "storageaccountname"
  resource_group_name = azurerm_resource_group.example.name

  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  network_rules {
    default_action             = "Deny"
	bypass					   = ["AzureServices"]
    ip_rules                   = ["100.0.0.1"]
    virtual_network_subnet_ids = [azurerm_subnet.example.id]
  }

  tags = {
    environment = "staging"
  }
}

resource "azurerm_storage_account_network_rules" "negative2" {
  resource_group_name  = azurerm_resource_group.test.name
  storage_account_name = azurerm_storage_account.test.name

  default_action             = "Allow"
  ip_rules                   = ["127.0.0.1"]
  virtual_network_subnet_ids = [azurerm_subnet.test.id]
  bypass                     = ["Metrics", "AzureServices"]
}
```
## Non-Compliant Code Examples
```terraform
resource "azurerm_storage_account_network_rules" "positive1" {
  resource_group_name  = azurerm_resource_group.test.name
  storage_account_name = azurerm_storage_account.test.name

  default_action             = "Allow"
  ip_rules                   = ["127.0.0.1"]
  virtual_network_subnet_ids = [azurerm_subnet.test.id]
  bypass                     = ["Metrics"]
}

resource "azurerm_storage_account" "positive2" {
  name                = "storageaccountname"
  resource_group_name = azurerm_resource_group.example.name

  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  network_rules {
    default_action             = "Deny"
	bypass					   = ["None"]
    ip_rules                   = ["100.0.0.1"]
    virtual_network_subnet_ids = [azurerm_subnet.example.id]
  }

  tags = {
    environment = "staging"
  }
}
```
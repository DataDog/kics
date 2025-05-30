---
title: "Small Flow Logs Retention Period"
meta:
  name: "azure/small_flow_logs_retention_period"
  id: "7750fcca-dd03-4d38-b663-4b70289bcfd4"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/small_flow_logs_retention_period`

**Id:** `7750fcca-dd03-4d38-b663-4b70289bcfd4`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Flow logs enable capturing information about IP traffic flowing in and out of the network security groups. Network Security Group Flow Logs must be enabled with retention period greater than or equal to 90 days. This is important, because these logs are used to check for anomalies and give information of suspected breaches

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_watcher_flow_log)

## Non-Compliant Code Examples
```terraform
resource "azurerm_network_watcher_flow_log" "positive1" {
  network_watcher_name      = azurerm_network_watcher.test.name
  resource_group_name       = azurerm_resource_group.test.name
  network_security_group_id = azurerm_network_security_group.test.id
  storage_account_id        = azurerm_storage_account.test.id
  enabled                   = true

  retention_policy {
    enabled = true
    days    = 89
  }
}

resource "azurerm_network_watcher_flow_log" "positive2" {
  network_watcher_name      = azurerm_network_watcher.test.name
  resource_group_name       = azurerm_resource_group.test.name
  network_security_group_id = azurerm_network_security_group.test.id
  storage_account_id        = azurerm_storage_account.test.id
  enabled                   = true

  retention_policy {
    enabled = true
    days    = 3
  }
}

resource "azurerm_network_watcher_flow_log" "positive3" {
  network_watcher_name      = azurerm_network_watcher.test.name
  resource_group_name       = azurerm_resource_group.test.name
  network_security_group_id = azurerm_network_security_group.test.id
  storage_account_id        = azurerm_storage_account.test.id
  enabled                   = true
}

resource "azurerm_network_watcher_flow_log" "positive4" {
  network_watcher_name      = azurerm_network_watcher.test.name
  resource_group_name       = azurerm_resource_group.test.name
  network_security_group_id = azurerm_network_security_group.test.id
  storage_account_id        = azurerm_storage_account.test.id
  enabled                   = true

  retention_policy {
    enabled = false
    days    = 900
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_network_watcher_flow_log" "negative1" {
    network_watcher_name = azurerm_network_watcher.test.name
    resource_group_name  = azurerm_resource_group.test.name
    network_security_group_id = azurerm_network_security_group.test.id
    storage_account_id        = azurerm_storage_account.test.id
    enabled                   = true

    retention_policy {
    enabled = true
    days    = 90
    }
}

resource "azurerm_network_watcher_flow_log" "negative2" {
    network_watcher_name = azurerm_network_watcher.test.name
    resource_group_name  = azurerm_resource_group.test.name
    network_security_group_id = azurerm_network_security_group.test.id
    storage_account_id        = azurerm_storage_account.test.id
    enabled                   = true

    retention_policy {
    enabled = true
    days    = 900
    }
}
```
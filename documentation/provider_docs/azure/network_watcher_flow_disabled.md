---
title: "Network Watcher Flow Disabled"
meta:
  name: "azure/network_watcher_flow_disabled"
  id: "b90842e5-6779-44d4-9760-972f4c03ba1c"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/network_watcher_flow_disabled`

**Id:** `b90842e5-6779-44d4-9760-972f4c03ba1c`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Check if enable field in the resource azurerm_network_watcher_flow_log is false.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_watcher_flow_log)

## Non-Compliant Code Examples
```terraform
resource "azurerm_network_watcher_flow_log" "positive1" {
  network_watcher_name = azurerm_network_watcher.test.name
  resource_group_name  = azurerm_resource_group.test.name

  network_security_group_id = azurerm_network_security_group.test.id
  storage_account_id        = azurerm_storage_account.test.id
  enabled                   = false

  retention_policy {
    enabled = true
    days    = 7
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
    days    = 7
  }
}
```
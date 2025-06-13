---
title: "Small Activity Log Retention Period"
meta:
  name: "azure/small_activity_log_retention_period"
  id: "2b856bf9-8e8c-4005-875f-303a8cba3918"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `azure/small_activity_log_retention_period`

**Id:** `2b856bf9-8e8c-4005-875f-303a8cba3918`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Observability

## Description
Ensure that Activity Log Retention is set 365 days or greater

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/monitor_log_profile)

## Non-Compliant Code Examples
```terraform
resource "azurerm_monitor_log_profile" "positive1" {
  name = "default"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  servicebus_rule_id = "${azurerm_eventhub_namespace.example.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.example.id

  retention_policy {
    enabled = true
    days    = 7
  }
}

resource "azurerm_monitor_log_profile" "positive2" {
  name = "default"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  servicebus_rule_id = "${azurerm_eventhub_namespace.example.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.example.id

  retention_policy {
    enabled = true
  }
}

resource "azurerm_monitor_log_profile" "positive3" {
  name = "default"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  servicebus_rule_id = "${azurerm_eventhub_namespace.example.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.example.id

  retention_policy {
    enabled = false
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_monitor_log_profile" "negative1" {
  name = "default"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  servicebus_rule_id = "${azurerm_eventhub_namespace.example.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.example.id

  retention_policy {
    enabled = true
    days    = 367
  }
}

resource "azurerm_monitor_log_profile" "negative2" {
  name = "default"

  categories = [
    "Action",
    "Delete",
    "Write",
  ]

  locations = [
    "westus",
    "global",
  ]

  servicebus_rule_id = "${azurerm_eventhub_namespace.example.id}/authorizationrules/RootManageSharedAccessKey"
  storage_account_id = azurerm_storage_account.example.id

  retention_policy {
    enabled = true
    days    = 0
  }
}

```
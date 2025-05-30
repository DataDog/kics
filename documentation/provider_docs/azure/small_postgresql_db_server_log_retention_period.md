---
title: "Small PostgreSQL DB Server Log Retention Period"
meta:
  name: "azure/small_postgresql_db_server_log_retention_period"
  id: "261a83f8-dd72-4e8c-b5e1-ebf06e8fe606"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `azure/small_postgresql_db_server_log_retention_period`

**Id:** `261a83f8-dd72-4e8c-b5e1-ebf06e8fe606`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Observability

## Description
Check if PostgreSQL Database Server retains logs for less than 3 Days

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "positive1" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_postgresql_server.example.name
  value               = 2
}
```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "negative1" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_postgresql_server.example.name
  value               = 5
}
```
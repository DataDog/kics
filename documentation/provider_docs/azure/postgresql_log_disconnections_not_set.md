---
title: "PostgreSQL Log Disconnections Not Set"
meta:
  name: "azure/postgresql_log_disconnections_not_set"
  id: "07f7134f-9f37-476e-8664-670c218e4702"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/postgresql_log_disconnections_not_set`

**Id:** `07f7134f-9f37-476e-8664-670c218e4702`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that for PostgreSQL Database, server parameter 'log_disconnections' is set to 'ON'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "positive1" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "off"
}

resource "azurerm_postgresql_configuration" "positive2" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "Off"
}

resource "azurerm_postgresql_configuration" "positive3" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "OFF"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "negative1" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "on"
}

resource "azurerm_postgresql_configuration" "negative2" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "On"
}

resource "azurerm_postgresql_configuration" "negative3" {
    name                = "log_disconnections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "ON"
}
```
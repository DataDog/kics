---
title: "PostgreSQL Log Connections Not Set"
meta:
  name: "azure/postgresql_log_connections_not_set"
  id: "c640d783-10c5-4071-b6c1-23507300d333"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/postgresql_log_connections_not_set`

**Id:** `c640d783-10c5-4071-b6c1-23507300d333`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that for PostgreSQL Database, server parameter 'log_connections' is set to 'ON'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "positive1" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "off"
}

resource "azurerm_postgresql_configuration" "positive2" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "Off"
}

resource "azurerm_postgresql_configuration" "positive3" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "OFF"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "negative1" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "on"
}

resource "azurerm_postgresql_configuration" "negative2" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "On"
}

resource "azurerm_postgresql_configuration" "negative3" {
    name                = "log_connections"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "ON"
}
```
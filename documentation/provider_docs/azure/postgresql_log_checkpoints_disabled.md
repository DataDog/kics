---
title: "PostgreSQL Log Checkpoints Disabled"
meta:
  name: "azure/postgresql_log_checkpoints_disabled"
  id: "3790d386-be81-4dcf-9850-eaa7df6c10d9"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/postgresql_log_checkpoints_disabled`

**Id:** `3790d386-be81-4dcf-9850-eaa7df6c10d9`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that for Postgre SQL Database Server, parameter 'log_checkpoints' is set to 'ON'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "positive1" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "off"
}

resource "azurerm_postgresql_configuration" "positive2" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "Off"
}

resource "azurerm_postgresql_configuration" "positive3" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "OFF"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "negative1" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "on"
}

resource "azurerm_postgresql_configuration" "negative2" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "On"
}

resource "azurerm_postgresql_configuration" "negative3" {
    name                = "log_checkpoints"
    resource_group_name = data.azurerm_resource_group.example.name
    server_name         = azurerm_postgresql_server.example.name
    value               = "ON"
}
```
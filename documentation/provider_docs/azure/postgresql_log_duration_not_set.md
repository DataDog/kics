---
title: "PostgreSQL Log Duration Not Set"
meta:
  name: "azure/postgresql_log_duration_not_set"
  id: "16e0879a-c4ae-4ff8-a67d-a2eed5d67b8f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `azure/postgresql_log_duration_not_set`

**Id:** `16e0879a-c4ae-4ff8-a67d-a2eed5d67b8f`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Observability

## Description
Make sure that for PostgreSQL Database, server parameter 'log_duration' is set to 'ON'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "azurerm_postgresql_configuration" "positive1" {
    name                = "log_duration"
    resource_group_name = "example1_resource_group_name"
    server_name         = "example1_server_name"
    value               = "off"
}

resource "azurerm_postgresql_configuration" "positive2" {
    name                = "log_duration"
    resource_group_name = "example2_resource_group_name"
    server_name         = "example2_server_name"
    value               = "Off"
}

resource "azurerm_postgresql_configuration" "positive3" {
    name                = "log_duration"
    resource_group_name = "example3_resource_group_name"
    server_name         = "example3_server_name"
    value               = "OFF"
}
```

## Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "azurerm_postgresql_configuration" "negative1" {
    name                = "log_duration"
    resource_group_name = "example1_resource_group_name"
    server_name         = "example1_server_name"
    value               = "on"
}

resource "azurerm_postgresql_configuration" "negative2" {
    name                = "log_duration"
    resource_group_name = "example2_resource_group_name"
    server_name         = "example2_server_name"
    value               = "On"
}

resource "azurerm_postgresql_configuration" "negative3" {
    name                = "log_duration"
    resource_group_name = "example3_resource_group_name"
    server_name         = "example3_server_name"
    value               = "ON"
}
```
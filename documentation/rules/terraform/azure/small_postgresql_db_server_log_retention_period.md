---
title: "Small PostgreSQL DB Server Log Retention Period"
group-id: "rules/terraform/azure"
meta:
  name: "azure/small_postgresql_db_server_log_retention_period"
  id: "261a83f8-dd72-4e8c-b5e1-ebf06e8fe606"
  display_name: "Small PostgreSQL DB Server Log Retention Period"
  cloud_provider: "azure"
  framework: "Terraform"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `261a83f8-dd72-4e8c-b5e1-ebf06e8fe606`

**Cloud Provider:** azure

**Framework:** Terraform

**Severity:** Low

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_configuration)

### Description

 This check verifies whether the `log_retention_days` configuration for an Azure PostgreSQL Database Server retains logs for at least 3 days. Insufficient log retention, such as setting `value = 2` in the Terraform resource

```
resource "azurerm_postgresql_configuration" "positive1" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_postgresql_server.example.name
  value               = 2
}
```

can hinder the ability to investigate security incidents or troubleshoot issues, as critical audit and activity logs may be deleted too quickly. Increasing the retention period to a secure value (such as `value = 5`) helps ensure logs are available for effective monitoring and forensic analysis.


## Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "negative1" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_postgresql_server.example.name
  value               = 5
}
```
## Non-Compliant Code Examples
```terraform
resource "azurerm_postgresql_configuration" "positive1" {
  name                = "log_retention_days"
  resource_group_name = azurerm_resource_group.example.name
  server_name         = azurerm_postgresql_server.example.name
  value               = 2
}
```
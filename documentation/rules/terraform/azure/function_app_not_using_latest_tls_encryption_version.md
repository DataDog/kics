---
title: "Function App Not Using Latest TLS Encryption Version"
meta:
  name: "azure/function_app_not_using_latest_tls_encryption_version"
  id: "45fc717a-bd86-415c-bdd8-677901be1aa6"
  display_name: "Function App Not Using Latest TLS Encryption Version"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata
**Name:** `azure/function_app_not_using_latest_tls_encryption_version`
**Query Name** `Function App Not Using Latest TLS Encryption Version`
**Id:** `45fc717a-bd86-415c-bdd8-677901be1aa6`
**Cloud Provider:** azure
**Platform** Terraform
**Severity:** Medium
**Category:** Encryption
## Description
Azure Function Apps should be configured to use the latest supported TLS version to ensure encrypted communications and protect data in transit. If the `min_tls_version` attribute is set to an outdated value such as `1.1`, as in:

```
site_config {
  min_tls_version = 1.1
}
```

the application becomes susceptible to known TLS vulnerabilities and exploits. Setting `min_tls_version` to at least `1.2`, as in

```
site_config {
  min_tls_version = 1.2
}
```

mitigates these risks and enforces stronger secure connections with clients.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/function_app#min_tls_version)


## Compliant Code Examples
```terraform
resource "azurerm_function_app" "negative3" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key
}

```

```terraform
resource "azurerm_function_app" "negative2" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
  }
}

```

```terraform
resource "azurerm_function_app" "negative1" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.2
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "azurerm_function_app" "positive2" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = "1.1"
  }
}

```

```terraform
resource "azurerm_function_app" "positive1" {
  name                       = "test-azure-functions"
  location                   = azurerm_resource_group.example.location
  resource_group_name        = azurerm_resource_group.example.name
  app_service_plan_id        = azurerm_app_service_plan.example.id
  storage_account_name       = azurerm_storage_account.example.name
  storage_account_access_key = azurerm_storage_account.example.primary_access_key

  site_config {
    dotnet_framework_version = "v4.0"
    scm_type                 = "LocalGit"
    min_tls_version = 1.1
  }
}

```
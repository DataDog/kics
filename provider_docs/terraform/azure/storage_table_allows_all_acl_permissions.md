---
title: "Storage Table Allows All ACL Permissions"
meta:
  name: "terraform/storage_table_allows_all_acl_permissions"
  id: "3ac3e75c-6374-4a32-8ba0-6ed69bda404e"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/storage_table_allows_all_acl_permissions`
**Id:** `3ac3e75c-6374-4a32-8ba0-6ed69bda404e`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing all ACL (Access Control List) permissions (`rwdl` for read, write, delete, and list) on an Azure Storage Table exposes the resource to excessive access, increasing the risk of unauthorized data modification or deletion. This misconfiguration may lead to data leakage, loss, or manipulation if the credentials are compromised or abused. To enhance security, permissions should be limited to only those operations necessary for the application's function, such as using only `r` for read access:

```
resource "azurerm_storage_table" "table_resource2" {
  name                 = "my_table_name"
  storage_account_name = "mystoragexxx"
  acl {
    id = "someid-1XXXXXXXXX"
    access_policy {
      expiry = "2022-10-03T05:05:00.0000000Z"
      permissions = "r"
      start = "2021-05-28T04:05:00.0000000Z"
    }
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_table#permissions)

## Non-Compliant Code Examples
```azure
resource "azurerm_storage_table" "table_resource" {
  name                 = "my_table_name"
  storage_account_name = "mystoragexxx"
  acl {
    id = "someid-1XXXXXXXXX"
    access_policy {
      expiry = "2022-10-03T05:05:00.0000000Z"
      permissions = "rwdl"
      start = "2021-05-28T04:05:00.0000000Z"
    }
  }
}

```

## Compliant Code Examples
```azure
resource "azurerm_storage_table" "table_resource2" {
  name                 = "my_table_name"
  storage_account_name = "mystoragexxx"
  acl {
    id = "someid-1XXXXXXXXX"
    access_policy {
      expiry = "2022-10-03T05:05:00.0000000Z"
      permissions = "r"
      start = "2021-05-28T04:05:00.0000000Z"
    }
  }
}

```
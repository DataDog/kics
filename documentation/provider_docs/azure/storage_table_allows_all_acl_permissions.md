---
title: "Storage Table Allows All ACL Permissions"
meta:
  name: "azure/storage_table_allows_all_acl_permissions"
  id: "3ac3e75c-6374-4a32-8ba0-6ed69bda404e"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/storage_table_allows_all_acl_permissions`

**Id:** `3ac3e75c-6374-4a32-8ba0-6ed69bda404e`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Azure Storage Table should not allow all ACL (Access Control List) permissions - r (read), w (write), d (delete), and l (list).

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_table#permissions)

## Non-Compliant Code Examples
```terraform
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
```terraform
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
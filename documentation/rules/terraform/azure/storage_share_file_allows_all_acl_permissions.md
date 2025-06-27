---
title: "Storage Share File Allows All ACL Permissions"
meta:
  name: "azure/storage_share_file_allows_all_acl_permissions"
  id: "48bbe0fd-57e4-4678-a4a1-119e79c90fc3"
  display_name: "Storage Share File Allows All ACL Permissions"
  cloud_provider: "azure"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `azure/storage_share_file_allows_all_acl_permissions`
**Query Name** `Storage Share File Allows All ACL Permissions`
**Id:** `48bbe0fd-57e4-4678-a4a1-119e79c90fc3`
**Cloud Provider:** azure
**Platform** Terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing all ACL (Access Control List) permissions—read (`r`), write (`w`), delete (`d`), and list (`l`)—on an Azure Storage Share File resource exposes the storage share to excessive privileges, increasing the risk of unauthorized access, data leakage, or malicious data manipulation. This misconfiguration could allow any user or process with the relevant access policy to not only read and list files, but also modify or delete important data, potentially leading to service disruption or data loss. To mitigate this risk, permissions should be set following the principle of least privilege; for example, granting only `r` (read) when read-only access is required, as shown below:

```
resource "azurerm_storage_share" "example" {
  name                 = "sharename"
  storage_account_name = "mystorageaccount"
  quota                = 50

  acl {
    id = "unique-acl-id"

    access_policy {
      permissions = "r"
      start       = "2024-06-07T09:38:21.0000000Z"
      expiry      = "2025-06-07T09:38:21.0000000Z"
    }
  }
}

resource "azurerm_storage_share_file" "example" {
  name             = "read-only-file.zip"
  storage_share_id = azurerm_storage_share.example.id
  source           = "some-local-file.zip"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/storage_share_file)


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
## Non-Compliant Code Examples
```terraform
resource "azurerm_storage_share" "example" {
  name                 = "sharename"
  storage_account_name = azurerm_storage_account.example.name
  quota                = 50

  acl {
    id = "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI"

    access_policy {
      permissions = "rwdl"
      start       = "2022-07-02T09:38:21.0000000Z"
      expiry      = "2021-07-02T10:38:21.0000000Z"
    }
  }
}

resource "azurerm_storage_share_file" "example" {
  name             = "my-awesome-content.zip"
  storage_share_id = azurerm_storage_share.example.id
  source           = "some-local-file.zip"
}

```
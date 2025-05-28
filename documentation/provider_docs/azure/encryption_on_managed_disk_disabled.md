---
title: "Encryption On Managed Disk Disabled"
meta:
  name: "azure/encryption_on_managed_disk_disabled"
  id: "a99130ab-4c0e-43aa-97f8-78d4fcb30024"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `azure/encryption_on_managed_disk_disabled`

**Id:** `a99130ab-4c0e-43aa-97f8-78d4fcb30024`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Encryption

## Description
Ensure that the encryption is active on the disk

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/managed_disk#encryption_settings)

## Non-Compliant Code Examples
```terraform
resource "azurerm_managed_disk" "positive1" {
  name                 = "acctestmd"
  location             = "West US 2"
  resource_group_name  = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"

  encryption_settings = {
      enabled = false
  }

  tags = {
    environment = "staging"
  }
}

resource "azurerm_managed_disk" "positive2" {
  name                 = "acctestmd"
  location             = "West US 2"
  resource_group_name  = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"
  

  tags = {
    environment = "staging"
  }
}
```

## Compliant Code Examples
```terraform

resource "azurerm_managed_disk" "negative1" {
  name                 = "acctestmd"
  location             = "West US 2"
  resource_group_name  = azurerm_resource_group.example.name
  storage_account_type = "Standard_LRS"
  create_option        = "Empty"
  disk_size_gb         = "1"
  
  encryption_settings = {
      enabled = true
  }
  tags = {
    environment = "staging"
  }
}
```
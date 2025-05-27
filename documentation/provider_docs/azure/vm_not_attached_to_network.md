---
title: "VM Not Attached To Network"
meta:
  name: "azure/vm_not_attached_to_network"
  id: "bbf6b3df-4b65-4f87-82cc-da9f30f8c033"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/vm_not_attached_to_network`

**Id:** `bbf6b3df-4b65-4f87-82cc-da9f30f8c033`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
No Network Security Group is attached to the Virtual Machine

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_machine#network_interface_ids)

## Non-Compliant Code Examples
```terraform
resource "azurerm_virtual_machine" "positive1" {
  name                  = "${var.prefix}-vm"
  location              = azurerm_resource_group.main.location
  resource_group_name   = azurerm_resource_group.main.name
  network_interface_ids = []
  vm_size               = "Standard_DS1_v2"

  os_profile_linux_config {
    disable_password_authentication = false
  }
}
```

## Compliant Code Examples
```terraform
resource "azurerm_network_interface" "negative1" {
  name                = "${var.prefix}-nic"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  ip_configuration {
    name                          = "testconfiguration1"
    subnet_id                     = azurerm_subnet.internal.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "negative2" {
  name                  = "${var.prefix}-vm"
  location              = azurerm_resource_group.main.location
  resource_group_name   = azurerm_resource_group.main.name
  network_interface_ids = [azurerm_network_interface.main.id]
  vm_size               = "Standard_DS1_v2"

  os_profile_linux_config {
    disable_password_authentication = false
  }
}
```
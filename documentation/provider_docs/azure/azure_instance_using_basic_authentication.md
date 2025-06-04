---
title: "Azure Instance Using Basic Authentication"
meta:
  name: "azure/azure_instance_using_basic_authentication"
  id: "dafe30ec-325d-4516-85d1-e8e6776f012c"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/azure_instance_using_basic_authentication`

**Id:** `dafe30ec-325d-4516-85d1-e8e6776f012c`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Best Practices

## Description
Azure Instances should use SSH Key instead of basic authentication

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_virtual_machine#admin_ssh_key)

## Non-Compliant Code Examples
```terraform
resource "azurerm_linux_virtual_machine" "positive1" {
  name                  = "${var.prefix}-vm"
  location              = azurerm_resource_group.main.location
  resource_group_name   = azurerm_resource_group.main.name
  network_interface_ids = []
  vm_size               = "Standard_DS1_v2"
  disable_password_authentication = false
}

```

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
resource "azurerm_virtual_machine" "negative1" {
  name                  = "${var.prefix}-vm"
  location              = azurerm_resource_group.main.location
  resource_group_name   = azurerm_resource_group.main.name
  network_interface_ids = [azurerm_network_interface.main.id]
  vm_size               = "Standard_DS1_v2"

  os_profile_linux_config {
    disable_password_authentication = true
  }

  admin_ssh_key {
    username   = "adminuser"
    public_key = file("~/.ssh/id_rsa.pub")
  }
}

```

```terraform
resource "azurerm_linux_virtual_machine" "negative1" {
  name                  = "${var.prefix}-vm"
  location              = azurerm_resource_group.main.location
  resource_group_name   = azurerm_resource_group.main.name
  network_interface_ids = [azurerm_network_interface.main.id]
  vm_size               = "Standard_DS1_v2"

  admin_ssh_key {
    username   = "adminuser"
    public_key = file("~/.ssh/id_rsa.pub")
  }
}

```
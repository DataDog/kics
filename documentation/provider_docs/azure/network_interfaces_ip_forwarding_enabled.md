---
title: "Network Interfaces IP Forwarding Enabled"
meta:
  name: "azure/network_interfaces_ip_forwarding_enabled"
  id: "4216ebac-d74c-4423-b437-35025cb88af5"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/network_interfaces_ip_forwarding_enabled`

**Id:** `4216ebac-d74c-4423-b437-35025cb88af5`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Network Interfaces IP Forwarding should be disabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface#enable_ip_forwarding)

## Non-Compliant Code Examples
```terraform
resource "azurerm_network_interface" "positive" {
  name                = "example-nic"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.example.id
    private_ip_address_allocation = "Dynamic"
  }

  enable_ip_forwarding = true
}

```

## Compliant Code Examples
```terraform
resource "azurerm_network_interface" "negative2" {
  name                = "example-nic"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.example.id
    private_ip_address_allocation = "Dynamic"
  }
}

```

```terraform
resource "azurerm_network_interface" "negative1" {
  name                = "example-nic"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.example.id
    private_ip_address_allocation = "Dynamic"
  }

  enable_ip_forwarding = false
}

```
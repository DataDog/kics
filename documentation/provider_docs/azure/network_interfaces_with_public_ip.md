---
title: "Network Interfaces With Public IP"
meta:
  name: "azure/network_interfaces_with_public_ip"
  id: "c1573577-e494-4417-8854-7e119368dc8b"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/network_interfaces_with_public_ip`

**Id:** `c1573577-e494-4417-8854-7e119368dc8b`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Network Interfaces should not be exposed with a public IP address. If configured, additional security baselines should be followed (https://docs.microsoft.com/en-us/security/benchmark/azure/baselines/virtual-network-security-baseline, https://docs.microsoft.com/en-us/security/benchmark/azure/baselines/public-ip-security-baseline)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface#public_ip_address_id)

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
    public_ip_address_id          = "IP"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_network_interface" "negative" {
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
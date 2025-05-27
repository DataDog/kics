---
title: "Ensure that UDP Services are restricted from the Internet"
meta:
  name: "azure/udp_services_not_restricted_from_internet"
  id: "a3b4c5d6-e7f8-9012-3456-7890abcdef12"
  cloud_provider: "azure"
  severity: "HIGH"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/udp_services_not_restricted_from_internet`

**Id:** `a3b4c5d6-e7f8-9012-3456-7890abcdef12`

**Cloud Provider:** azure

**Severity:** High

**Category:** Networking and Firewall

## Description
Network Security Group (NSG) rules that allow unrestricted UDP traffic from the internet (0.0.0.0/0) create significant security vulnerabilities in Azure environments. UDP is a connectionless protocol that doesn't provide built-in security controls, making services using it particularly susceptible to DDoS attacks, packet spoofing, and unauthorized access when exposed publicly. Instead of using broad source address prefixes like '0.0.0.0/0', restrict inbound UDP traffic to specific IP ranges or CIDR blocks that require access.

Secure example:
```
security_rule {
  protocol                  = "udp"
  source_address_prefix     = "192.168.1.0/24"
  destination_port_range    = "53"
}
```

Vulnerable example:
```
security_rule {
  protocol                  = "udp"
  source_address_prefix     = "0.0.0.0/0"
  destination_port_range    = "53"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group)

## Non-Compliant Code Examples
```terraform
resource "azurerm_network_security_group" "bad_example" {
  name                = "bad-nsg"
  location            = "East US"
  resource_group_name = "example-rg"

  security_rule {
    name                       = "Allow UDP Inbound"
    protocol                   = "udp"
    direction                  = "inbound"
    access                     = "allow"
    priority                   = 100
    source_address_prefix      = "0.0.0.0/0"
    destination_address_prefix = "*"
    source_port_range          = "*"
    destination_port_range     = "53"
  }
}

resource "azurerm_network_security_rule" "bad_example_rule" {
  name                        = "Allow UDP Inbound"
  resource_group_name         = "example-rg"
  network_security_group_name = "bad-nsg"
  priority                    = 100
  direction                   = "inbound"
  access                      = "allow"
  protocol                    = "udp"
  source_address_prefix       = "0.0.0.0/0"
  destination_address_prefix  = "*"
  source_port_range           = "*"
  destination_port_range      = "53"
}

```

## Compliant Code Examples
```terraform
# Example for azurerm_network_security_rule (standalone)
resource "azurerm_network_security_rule" "good_example_rule" {
  name                        = "Allow UDP Inbound"
  resource_group_name         = "example-rg"
  network_security_group_name = "good-nsg"
  priority                    = 100
  direction                   = "inbound"
  access                      = "allow"
  protocol                    = "udp"
  source_address_prefix       = "192.168.1.0/24" # ✅ Restricted access
  destination_address_prefix  = "*"
  source_port_range           = "*"
  destination_port_range      = "53"
}

```

```terraform
# Example for azurerm_network_security_group (embedded security_rule)
resource "azurerm_network_security_group" "good_example" {
  name                = "good-nsg"
  location            = "East US"
  resource_group_name = "example-rg"

  security_rule {
    name                       = "Allow UDP Inbound"
    protocol                   = "udp"
    direction                  = "inbound"
    access                     = "allow"
    priority                   = 100
    source_address_prefix      = "192.168.1.0/24" # ✅ Restricted access
    destination_address_prefix = "*"
    source_port_range          = "*"
    destination_port_range     = "53"
  }
}


```
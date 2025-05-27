---
title: "Redis Publicly Accessible"
meta:
  name: "azure/redis_publicly_accessible"
  id: "5089d055-53ff-421b-9482-a5267bdce629"
  cloud_provider: "azure"
  severity: "CRITICAL"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/redis_publicly_accessible`

**Id:** `5089d055-53ff-421b-9482-a5267bdce629`

**Cloud Provider:** azure

**Severity:** Critical

**Category:** Networking and Firewall

## Description
Azure Redis Cache instances with firewall rules that allow access from public IP addresses are vulnerable to unauthorized access and potential data breaches. When configuring firewall rules, using public IP ranges (like '1.2.3.4' to '2.3.4.5') exposes your Redis Cache to the internet, allowing potential attackers to attempt brute force attacks or exploit vulnerabilities.

Instead, limit access to private IP ranges within your internal network (like '10.2.3.4' to '10.3.4.5') as shown below:

```terraform
resource "azurerm_redis_firewall_rule" "example" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "10.2.3.4"  // Private IP range
  end_ip              = "10.3.4.5"  // Private IP range
}
```

This ensures your Redis Cache is only accessible from within your virtual network, significantly reducing the attack surface.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_firewall_rule)

## Non-Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "positive1" {
  name                = "redis${random_id.server.hex}"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  capacity            = 1
  family              = "P"
  sku_name            = "Premium"
  enable_non_ssl_port = false

  redis_configuration {
    maxclients         = 256
    maxmemory_reserved = 2
    maxmemory_delta    = 2
    maxmemory_policy   = "allkeys-lru"
  }
}

resource "azurerm_redis_firewall_rule" "positive2" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.2.3.4"
  end_ip              = "2.3.4.5"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "negative1" {
  name                = "redis${random_id.server.hex}"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  capacity            = 1
  family              = "P"
  sku_name            = "Premium"
  enable_non_ssl_port = false

  redis_configuration {
    maxclients         = 256
    maxmemory_reserved = 2
    maxmemory_delta    = 2
    maxmemory_policy   = "allkeys-lru"
  }
}

resource "azurerm_redis_firewall_rule" "negative2" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "10.2.3.4"
  end_ip              = "10.3.4.5"
}
```
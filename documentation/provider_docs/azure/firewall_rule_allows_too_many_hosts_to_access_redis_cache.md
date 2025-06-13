---
title: "Firewall Rule Allows Too Many Hosts To Access Redis Cache"
meta:
  name: "azure/firewall_rule_allows_too_many_hosts_to_access_redis_cache"
  id: "a829b715-cf75-4e92-b645-54c9b739edfb"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `azure/firewall_rule_allows_too_many_hosts_to_access_redis_cache`

**Id:** `a829b715-cf75-4e92-b645-54c9b739edfb`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Check if any firewall rule allows too many hosts to access Redis Cache

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_firewall_rule)

## Non-Compliant Code Examples
```terraform
resource "azurerm_redis_firewall_rule" "positive1" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.0.0.0"
  end_ip              = "3.0.0.0"
}
```

## Compliant Code Examples
```terraform
resource "azurerm_redis_firewall_rule" "negative1" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.2.3.4"
  end_ip              = "1.2.3.8"
}
```
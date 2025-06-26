---
title: "Firewall Rule Allows Too Many Hosts To Access Redis Cache"
meta:
  name: "terraform/firewall_rule_allows_too_many_hosts_to_access_redis_cache"
  id: "a829b715-cf75-4e92-b645-54c9b739edfb"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/firewall_rule_allows_too_many_hosts_to_access_redis_cache`
**Id:** `a829b715-cf75-4e92-b645-54c9b739edfb`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
This check assesses whether any Redis Cache firewall rule is configured to allow access from an overly broad IP range, such as specifying `start_ip = "1.0.0.0"` and `end_ip = "3.0.0.0"`. Allowing too many hosts to access the Redis Cache can expose sensitive data or enable unauthorized users to exploit the cache service. Firewall rules should narrowly define permitted IPs, for example:

```
resource "azurerm_redis_firewall_rule" "secure_example" {
  name                = "limitedAccess"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.2.3.4"
  end_ip              = "1.2.3.8"
}
```

Restricting access to only the necessary hosts mitigates the risk of data breaches and service misuse.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_firewall_rule)

## Non-Compliant Code Examples
```azure
resource "azurerm_redis_firewall_rule" "positive1" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.0.0.0"
  end_ip              = "3.0.0.0"
}
```

## Compliant Code Examples
```azure
resource "azurerm_redis_firewall_rule" "negative1" {
  name                = "someIPrange"
  redis_cache_name    = azurerm_redis_cache.example.name
  resource_group_name = azurerm_resource_group.example.name
  start_ip            = "1.2.3.4"
  end_ip              = "1.2.3.8"
}
```
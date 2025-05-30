---
title: "Redis Not Updated Regularly"
meta:
  name: "azure/redis_not_updated_regularly"
  id: "b947809d-dd2f-4de9-b724-04d101c515aa"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/redis_not_updated_regularly`

**Id:** `b947809d-dd2f-4de9-b724-04d101c515aa`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Redis Cache is not configured to be updated regularly with security and operational updates

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_cache#patch_schedule)

## Non-Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "positive1" {
    name                = "timeout-redis"
    location            = "West Europe"
    resource_group_name = azurerm_resource_group.example_rg.name
    subnet_id           = azurerm_subnet.example_redis_snet.id

    family              = "P"
    capacity            = 1
    sku_name            = "Premium"
    shard_count         = 1

    enable_non_ssl_port = false
    minimum_tls_version = "1.2"

    redis_configuration {
        enable_authentication   = true
        maxmemory_policy        = "volatile-lru"
    }
}
```

## Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "negative1" {
    name                = "timeout-redis"
    location            = "West Europe"
    resource_group_name = azurerm_resource_group.example_rg.name
    subnet_id           = azurerm_subnet.example_redis_snet.id

    family              = "P"
    capacity            = 1
    sku_name            = "Premium"
    shard_count         = 1

    enable_non_ssl_port = false
    minimum_tls_version = "1.2"

    redis_configuration {
        enable_authentication   = true
        maxmemory_policy        = "volatile-lru"
    }

    patch_schedule {
        day_of_week     = "Thursday"
        start_hour_utc  = 7
    }
}
```
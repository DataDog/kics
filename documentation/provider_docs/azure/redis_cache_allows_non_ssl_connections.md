---
title: "Redis Cache Allows Non SSL Connections"
meta:
  name: "azure/redis_cache_allows_non_ssl_connections"
  id: "e29a75e6-aba3-4896-b42d-b87818c16b58"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/redis_cache_allows_non_ssl_connections`

**Id:** `e29a75e6-aba3-4896-b42d-b87818c16b58`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Redis Cache resources should not allow non-SSL connections

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_cache)

## Non-Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "positive1" {
  name                = "example-cache"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  capacity            = 2
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = true
  minimum_tls_version = "1.2"

  redis_configuration {
  }
}
```

## Compliant Code Examples
```terraform
resource "azurerm_redis_cache" "negative1" {
  name                = "example-cache"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  capacity            = 2
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = false
  minimum_tls_version = "1.2"

  redis_configuration {
  }
}

resource "azurerm_redis_cache" "negative2" {
  name                = "example-cache"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  capacity            = 2
  family              = "C"
  sku_name            = "Standard"
 
  minimum_tls_version = "1.2"

  redis_configuration {
  }
}
```
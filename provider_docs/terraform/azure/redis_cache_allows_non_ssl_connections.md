---
title: "Redis Cache Allows Non SSL Connections"
meta:
  name: "terraform/redis_cache_allows_non_ssl_connections"
  id: "e29a75e6-aba3-4896-b42d-b87818c16b58"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/redis_cache_allows_non_ssl_connections`
**Id:** `e29a75e6-aba3-4896-b42d-b87818c16b58`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Allowing non-SSL connections to Azure Redis Cache resources exposes sensitive data in transit to potential interception or man-in-the-middle attacks, as information exchanged between clients and the Redis service will not be encrypted. To mitigate this risk, the `enable_non_ssl_port` attribute in the Terraform resource should be set to `false`, ensuring all connections use secure TLS communication. For example:

```
resource "azurerm_redis_cache" "secure_example" {
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
```

Neglecting this configuration can lead to exposure of authentication credentials and cached data, increasing the risk of data breaches and compliance violations.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/redis_cache)

## Non-Compliant Code Examples
```azure
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
```azure
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
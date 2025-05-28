---
title: "Azure Active Directory Authentication"
meta:
  name: "azure/azure_active_directory_authentication"
  id: "a21c8da9-41bf-40cf-941d-330cf0d11fc7"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Access Control"
---

## Metadata
**Name:** `azure/azure_active_directory_authentication`

**Id:** `a21c8da9-41bf-40cf-941d-330cf0d11fc7`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Access Control

## Description
Azure Active Directory must be used for authentication for Service Fabric

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/service_fabric_cluster#tenant_id)

## Non-Compliant Code Examples
```terraform
resource "azurerm_service_fabric_cluster" "positive2" {
  name                 = "example-servicefabric"
  resource_group_name  = azurerm_resource_group.example.name
  location             = azurerm_resource_group.example.location
  reliability_level    = "Bronze"
  upgrade_mode         = "Manual"
  cluster_code_version = "7.1.456.959"
  vm_image             = "Windows"
  management_endpoint  = "https://example:80"

  node_type {
    name                 = "first"
    instance_count       = 3
    is_primary           = true
    client_endpoint_port = 2020
    http_endpoint_port   = 80
  }
}

```

```terraform
resource "azurerm_service_fabric_cluster" "positive1" {
  name                 = "example-servicefabric"
  resource_group_name  = azurerm_resource_group.example.name
  location             = azurerm_resource_group.example.location
  reliability_level    = "Bronze"
  upgrade_mode         = "Manual"
  cluster_code_version = "7.1.456.959"
  vm_image             = "Windows"
  management_endpoint  = "https://example:80"

  node_type {
    name                 = "first"
    instance_count       = 3
    is_primary           = true
    client_endpoint_port = 2020
    http_endpoint_port   = 80
  }

  azure_active_directory {
    cluster_application_id = "id"
    client_application_id = "id"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_service_fabric_cluster" "negative" {
  name                 = "example-servicefabric"
  resource_group_name  = azurerm_resource_group.example.name
  location             = azurerm_resource_group.example.location
  reliability_level    = "Bronze"
  upgrade_mode         = "Manual"
  cluster_code_version = "7.1.456.959"
  vm_image             = "Windows"
  management_endpoint  = "https://example:80"

  node_type {
    name                 = "first"
    instance_count       = 3
    is_primary           = true
    client_endpoint_port = 2020
    http_endpoint_port   = 80
  }

  azure_active_directory {
    tenant_id = "id"
    cluster_application_id = "id"
    client_application_id = "id"
  }
}

```
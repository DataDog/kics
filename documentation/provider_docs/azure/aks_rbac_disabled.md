---
title: "AKS RBAC Disabled"
meta:
  name: "azure/aks_rbac_disabled"
  id: "86f92117-eed8-4614-9c6c-b26da20ff37f"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `azure/aks_rbac_disabled`

**Id:** `86f92117-eed8-4614-9c6c-b26da20ff37f`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Access Control

## Description
Azure Container Service (AKS) instance should have role-based access control (RBAC) enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster#role_based_access_control)

## Non-Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "positive1" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  role_based_access_control_enabled = false

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }

  network_profile {
    network_policy = "azure"
  }
}

resource "azurerm_kubernetes_cluster" "positive2" {
  name                = "example-aks2"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks2"

  role_based_access_control {
    enabled = false
  }

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }

  network_profile {
    network_policy = "calico"
  }
}

```

## Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "negative1" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  role_based_access_control_enabled = true

  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_D2_v2"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = {
    Environment = "Production"
  }

  network_profile {
    network_policy = "azure"
  }
}


```
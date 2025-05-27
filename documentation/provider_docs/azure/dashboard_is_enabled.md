---
title: "Dashboard Is Enabled"
meta:
  name: "azure/dashboard_is_enabled"
  id: "61c3cb8b-0715-47e4-b788-86dde40dd2db"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/dashboard_is_enabled`

**Id:** `61c3cb8b-0715-47e4-b788-86dde40dd2db`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Insecure Configurations

## Description
Check if the Kubernetes Dashboard is enabled.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster)

## Non-Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "positive1" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

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
  
  addon_profile {
    kube_dashboard {
      enabled = true
    }
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
}

resource "azurerm_kubernetes_cluster" "negative2" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

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

  addon_profile {
    kube_dashboard {
        enabled = false
    }
  }
}
```
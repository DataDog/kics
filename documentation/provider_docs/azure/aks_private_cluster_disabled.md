---
title: "AKS Private Cluster Disabled"
meta:
  name: "azure/aks_private_cluster_disabled"
  id: "599318f2-6653-4569-9e21-041d06c63a89"
  cloud_provider: "azure"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `azure/aks_private_cluster_disabled`

**Id:** `599318f2-6653-4569-9e21-041d06c63a89`

**Cloud Provider:** azure

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Azure Kubernetes Service (AKS) API should not be exposed to the internet

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster#private_cluster_enabled)

## Non-Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "positive2" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

}

```

```terraform
resource "azurerm_kubernetes_cluster" "positive1" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  private_cluster_enabled = false
}

```

## Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "negative" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  private_cluster_enabled = true
}

```
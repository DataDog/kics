---
title: "AKS Uses Azure Policies Add-On Disabled"
meta:
  name: "azure/aks_uses_azure_policies_addon_disabled"
  id: "43789711-161b-4708-b5bb-9d1c626f7492"
  cloud_provider: "azure"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `azure/aks_uses_azure_policies_addon_disabled`

**Id:** `43789711-161b-4708-b5bb-9d1c626f7492`

**Cloud Provider:** azure

**Severity:** Low

**Category:** Best Practices

## Description
Azure Container Service (AKS) should use Azure Policies Add-On

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster#azure_policy)

## Non-Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "positive2" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  azure_policy_enabled = false
}

```

```terraform
resource "azurerm_kubernetes_cluster" "positive3" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  addon_profile {}
}

```

```terraform
resource "azurerm_kubernetes_cluster" "positive4" {
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

  addon_profile {

   azure_policy {

     enabled = false

   }
 }
}


```

## Compliant Code Examples
```terraform
resource "azurerm_kubernetes_cluster" "negative" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  azure_policy_enabled = true
}

```

```terraform
resource "azurerm_kubernetes_cluster" "negative" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  addon_profile {

   azure_policy {

     enabled = true

   }
 }
}


```
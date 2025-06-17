---
title: "AKS Uses Azure Policies Add-On Disabled"
meta:
  name: "terraform/aks_uses_azure_policies_addon_disabled"
  id: "43789711-161b-4708-b5bb-9d1c626f7492"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/aks_uses_azure_policies_addon_disabled`
**Id:** `43789711-161b-4708-b5bb-9d1c626f7492`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Best Practices
## Description
Enabling the Azure Policy Add-On for Azure Kubernetes Service (AKS) clusters helps enforce organizational standards and compliance at scale by applying policy controls directly to the cluster. If the `addon_profile.azure_policy.enabled` attribute is set to `false`, as shown below, the cluster will not have Azure Policy integration, leaving it vulnerable to resource misconfigurations and violating compliance policies.

```
addon_profile {
  azure_policy {
    enabled = false
  }
}
```

To mitigate this vulnerability, the policy add-on should be enabled by setting `enabled = true`, ensuring that security and compliance policies are consistently enforced within the AKS environment.

```
addon_profile {
  azure_policy {
    enabled = true
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster#azure_policy)

## Non-Compliant Code Examples
```azure
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

```azure
resource "azurerm_kubernetes_cluster" "positive4" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"
}

```

```azure
resource "azurerm_kubernetes_cluster" "positive3" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  addon_profile {}
}

```

## Compliant Code Examples
```azure
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

```azure
resource "azurerm_kubernetes_cluster" "negative" {
  name                = "example-aks1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  dns_prefix          = "exampleaks1"

  azure_policy_enabled = true
}

```
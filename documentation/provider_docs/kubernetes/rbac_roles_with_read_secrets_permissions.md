---
title: "RBAC Roles with Read Secrets Permissions"
meta:
  name: "kubernetes/rbac_roles_with_read_secrets_permissions"
  id: "826abb30-3cd5-4e0b-a93b-67729b4f7e63"
  cloud_provider: "kubernetes"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `kubernetes/rbac_roles_with_read_secrets_permissions`

**Id:** `826abb30-3cd5-4e0b-a93b-67729b4f7e63`

**Cloud Provider:** kubernetes

**Severity:** Medium

**Category:** Access Control

## Description
Roles and ClusterRoles with get/watch/list RBAC permissions on Kubernetes secrets are dangerous and should be avoided. In case of compromise, attackers could abuse these roles to access sensitive data, such as passwords, tokens and keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/role#rule)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_role" "example1" {
  metadata {
    name = "terraform-example1"
    labels = {
      test = "MyRole"
    }
  }

  rule {
    api_groups     = [""]
    resources      = ["secrets", "namespaces"]
    resource_names = ["foo"]
    verbs          = ["get", "list", "watch"]
  }
  rule {
    api_groups = ["apps"]
    resources  = ["deployments"]
    verbs      = ["get", "list"]
  }
}

resource "kubernetes_cluster_role" "example2" {
  metadata {
    name = "terraform-example2"
  }

  rule {
    api_groups = [""]
    resources  = ["namespaces", "secrets"]
    verbs      = ["get", "list", "watch"]
  }
  rule {
    api_groups = ["apps"]
    resources  = ["deployments"]
    verbs      = ["get", "list"]
  }
}


resource "kubernetes_role" "example3" {
  metadata {
    name = "terraform-example3"
    labels = {
      test = "MyRole"
    }
  }

  rule {
    api_groups     = [""]
    resources      = ["secrets", "namespaces"]
    resource_names = ["foo"]
    verbs          = ["get", "list", "watch"]
  }

}

resource "kubernetes_cluster_role" "example4" {
  metadata {
    name = "terraform-example4"
  }

  rule {
    api_groups = [""]
    resources  = ["namespaces", "secrets"]
    verbs      = ["get", "list", "watch"]
  }

}

```

## Compliant Code Examples
```terraform
resource "kubernetes_role" "example1" {
  metadata {
    name = "terraform-example1"
    labels = {
      test = "MyRole"
    }
  }

  rule {
    api_groups     = [""]
    resources      = ["pods"]
    resource_names = ["foo"]
    verbs          = ["get", "list", "watch"]
  }
  rule {
    api_groups = ["apps"]
    resources  = ["deployments"]
    verbs      = ["get", "list"]
  }
}

resource "kubernetes_cluster_role" "example2" {
  metadata {
    name = "terraform-example2"
  }

  rule {
    api_groups = [""]
    resources  = ["namespaces", "pods"]
    verbs      = ["get", "list", "watch"]
  }
}

```
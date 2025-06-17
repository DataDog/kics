---
title: "Cluster Admin Rolebinding With Superuser Permissions"
meta:
  name: "terraform/cluster_admin_role_binding_with_super_user_permissions"
  id: "17172bc2-56fb-4f17-916f-a014147706cd"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/cluster_admin_role_binding_with_super_user_permissions`
**Id:** `17172bc2-56fb-4f17-916f-a014147706cd`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Access Control
## Description
Ensure that the cluster-admin role is only used where required (RBAC)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/cluster_role_binding#name)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_cluster_role_binding" "example2" {
  metadata {
    name = "terraform-example2"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subject {
    kind      = "User"
    name      = "admin"
    api_group = "rbac.authorization.k8s.io"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "default"
    namespace = "kube-system"
  }
  subject {
    kind      = "Group"
    name      = "system:masters"
    api_group = "rbac.authorization.k8s.io"
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_cluster_role_binding" "example1" {
  metadata {
    name = "terraform-example1"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster"
  }
  subject {
    kind      = "User"
    name      = "admin"
    api_group = "rbac.authorization.k8s.io"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "default"
    namespace = "kube-system"
  }
  subject {
    kind      = "Group"
    name      = "system:masters"
    api_group = "rbac.authorization.k8s.io"
  }
}

```
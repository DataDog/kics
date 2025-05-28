---
title: "Service Type is NodePort"
meta:
  name: "kubernetes/service_type_is_nodeport"
  id: "5c281bf8-d9bb-47f2-b909-3f6bb11874ad"
  cloud_provider: "kubernetes"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `kubernetes/service_type_is_nodeport`

**Id:** `5c281bf8-d9bb-47f2-b909-3f6bb11874ad`

**Cloud Provider:** kubernetes

**Severity:** Low

**Category:** Networking and Firewall

## Description
Service type should not be NodePort

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service#type)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_service" "example" {
  metadata {
    name = "terraform-example"
  }
  spec {
    selector = {
      app = kubernetes_pod.example.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 8080
      target_port = 80
    }

    type = "NodePort"
  }
}

```

## Compliant Code Examples
```terraform
resource "kubernetes_service" "example2" {
  metadata {
    name = "terraform-example"
  }
  spec {
    selector = {
      app = kubernetes_pod.example.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 8080
      target_port = 80
    }

    type = "LoadBalancer"
  }
}

```
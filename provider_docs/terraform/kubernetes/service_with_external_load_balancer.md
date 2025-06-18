---
title: "Service With External Load Balancer"
meta:
  name: "terraform/service_with_external_load_balancer"
  id: "2a52567c-abb8-4651-a038-52fa27c77aed"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/service_with_external_load_balancer`
**Id:** `2a52567c-abb8-4651-a038-52fa27c77aed`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Service has an external load balancer, which may cause accessibility from other networks and the Internet

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_service" "example1" {
  metadata {
    name = "terraform-example1"
    annotations = {
      "service.beta.kubernetes.io/aws-load-balancer-internal" = "false"
    }
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

resource "kubernetes_service" "example2" {
  metadata {
    name = "terraform-example2"
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

```kubernetes
resource "kubernetes_service" "example2" {
  metadata {
    name = "terraform-example2"
    annotations = {
      "service.beta.kubernetes.io/azure-load-balancer-internal" = "false"
    }
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

resource "kubernetes_service" "example3" {
  metadata {
    name = "terraform-example3"
    annotations = {
      "networking.gke.io/load-balancer-type" = "External"
    }
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

resource "kubernetes_service" "example4" {
  metadata {
    name = "terraform-example4"
    annotations = {
      "cloud.google.com/load-balancer-type" = "External"
    }
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

## Compliant Code Examples
```kubernetes
resource "kubernetes_service" "example3" {
  metadata {
    name = "terraform-example3"
    annotations = {
      "service.beta.kubernetes.io/aws-load-balancer-internal" = "true"
    }
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

```kubernetes
resource "kubernetes_service" "example2" {
  metadata {
    name = "terraform-example2"
    annotations = {
      "service.beta.kubernetes.io/azure-load-balancer-internal" = "true"
    }
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

resource "kubernetes_service" "example3" {
  metadata {
    name = "terraform-example3"
    annotations = {
      "networking.gke.io/load-balancer-type" = "Internal"
    }
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

resource "kubernetes_service" "example4" {
  metadata {
    name = "terraform-example4"
    annotations = {
      "cloud.google.com/load-balancer-type" = "Internal"
    }
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
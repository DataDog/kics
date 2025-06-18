---
title: "Shared Host IPC Namespace"
meta:
  name: "terraform/shared_host_ipc_namespace"
  id: "e94d3121-c2d1-4e34-a295-139bfeb73ea3"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata
**Name:** `terraform/shared_host_ipc_namespace`
**Id:** `e94d3121-c2d1-4e34-a295-139bfeb73ea3`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Resource Management
## Description
Container should not share the host IPC namespace

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#host_ipc)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "positive1" {
  metadata {
    name = "terraform-example"
  }

  spec {

    host_ipc = true

    container {
      image = "nginx:1.7.9"
      name  = "example"


      env {
        name  = "environment"
        value = "test"
      }

      port {
        container_port = 8080
      }

      liveness_probe {
        http_get {
          path = "/nginx_status"
          port = 80

          http_header {
            name  = "X-Custom-Header"
            value = "Awesome"
          }
        }

        initial_delay_seconds = 3
        period_seconds        = 3
      }
    }

    dns_config {
      nameservers = ["1.1.1.1", "8.8.8.8", "9.9.9.9"]
      searches    = ["example.com"]

      option {
        name  = "ndots"
        value = 1
      }

      option {
        name = "use-vc"
      }
    }

    dns_policy = "None"
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "negative1" {
  metadata {
    name = "terraform-example"
  }

  spec {

    host_ipc = false

    container {
      image = "nginx:1.7.9"
      name  = "example"


      env {
        name  = "environment"
        value = "test"
      }

      port {
        container_port = 8080
      }

      liveness_probe {
        http_get {
          path = "/nginx_status"
          port = 80

          http_header {
            name  = "X-Custom-Header"
            value = "Awesome"
          }
        }

        initial_delay_seconds = 3
        period_seconds        = 3
      }
    }

    dns_config {
      nameservers = ["1.1.1.1", "8.8.8.8", "9.9.9.9"]
      searches    = ["example.com"]

      option {
        name  = "ndots"
        value = 1
      }

      option {
        name = "use-vc"
      }
    }

    dns_policy = "None"
  }
}




```
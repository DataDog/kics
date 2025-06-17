---
title: "Readiness Probe Is Not Configured"
meta:
  name: "terraform/readiness_probe_is_not_configured"
  id: "8657197e-3f87-4694-892b-8144701d83c1"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Availability"
---
## Metadata
**Name:** `terraform/readiness_probe_is_not_configured`
**Id:** `8657197e-3f87-4694-892b-8144701d83c1`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Availability
## Description
Check if Readiness Probe is not configured.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#readiness_probe)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "test" {
  metadata {
    name = "terraform-example"
  }

  spec {
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
resource "kubernetes_pod" "test2" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      readiness_probe {
        initial_delay_seconds = 10
      }

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
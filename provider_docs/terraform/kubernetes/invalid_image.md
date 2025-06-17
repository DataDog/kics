---
title: "Invalid Image"
meta:
  name: "terraform/invalid_image"
  id: "e76cca7c-c3f9-4fc9-884c-b2831168ebd8"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Supply-Chain"
---
## Metadata
**Name:** `terraform/invalid_image`
**Id:** `e76cca7c-c3f9-4fc9-884c-b2831168ebd8`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Supply-Chain
## Description
Image must be defined and not be empty or equal to latest.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#image)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "positive1" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      image = ""
      name  = "example"

      env = {
        name  = "environment"
        value = "test"
      }

      port = {
        container_port = 8080
      }

      liveness_probe = {
        http_get  = {
          path = "/nginx_status"
          port = 80

          http_header = {
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

resource "kubernetes_pod" "positive2" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
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


resource "kubernetes_pod" "positive3" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container = [
      {
        image = "latest"
        name  = "example"

        env = {
          name  = "environment"
          value = "test"
        }

        port = {
          container_port = 8080
        }

        liveness_probe = {
          http_get = {
            path = "/nginx_status"
            port = 80

            http_header = {
              name  = "X-Custom-Header"
              value = "Awesome"
            }
          }

          initial_delay_seconds = 3
          period_seconds        = 3
        }
      }
    ]

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
resource "kubernetes_pod" "negative" {
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
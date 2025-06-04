---
title: "Image Without Digest"
meta:
  name: "kubernetes/image_without_digest"
  id: "228c4c19-feeb-4c18-848c-800ac70fdfb7"
  cloud_provider: "kubernetes"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `kubernetes/image_without_digest`

**Id:** `228c4c19-feeb-4c18-848c-800ac70fdfb7`

**Cloud Provider:** kubernetes

**Severity:** Low

**Category:** Insecure Configurations

## Description
Images should be specified together with their digests to ensure integrity

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#image)

## Non-Compliant Code Examples
```terraform
resource "kubernetes_pod" "positive1" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      image = "uses-private-image"
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
        image = "uses-private-image"
        name  = "example"

        env = {
          name  = "environment"
          value = "test"
        }

        port =  {
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
```terraform
resource "kubernetes_pod" "negative" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      image = "image@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb"
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
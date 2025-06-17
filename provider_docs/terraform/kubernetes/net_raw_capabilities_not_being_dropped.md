---
title: "NET_RAW Capabilities Not Being Dropped"
meta:
  name: "terraform/net_raw_capabilities_not_being_dropped"
  id: "e5587d53-a673-4a6b-b3f2-ba07ec274def"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/net_raw_capabilities_not_being_dropped`
**Id:** `e5587d53-a673-4a6b-b3f2-ba07ec274def`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Containers should drop 'ALL' or at least 'NET_RAW' capabilities

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#drop)

## Non-Compliant Code Examples
```kubernetes

resource "kubernetes_pod" "positive1" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container = [
     {
      image = "nginx:1.7.9"
      name  = "example22"

      security_context = {
        capabilities = {
          drop = ["NET_BIND_SERVICE"]
        }
      }

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
     ,
     {
      image = "nginx:1.7.9"
      name  = "example22222"

      security_context = {
        read_only_root_filesystem = true
      }

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
     ,
     {
      image = "nginx:1.7.9"
      name  = "example3"

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



resource "kubernetes_pod" "positive2" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      image = "nginx:1.7.9"
      name  = "example"

      security_context {
        capabilities {
          drop = ["NET_BIND_SERVICE"]
        }
      }

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

resource "kubernetes_pod" "negative3" {
  metadata {
    name = "terraform-example"
  }

  spec {

    container =  [
      {
        image = "nginx:1.7.9"
        name  = "example"

        security_context = {
          capabilities = {
            drop = ["ALL"]
          }
        }

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
      },

      {
        image = "nginx:1.7.9"
        name  = "example2"

        security_context = {
          capabilities = {
            drop = ["ALL"]
          }
        }

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



resource "kubernetes_pod" "negative4" {
  metadata {
    name = "terraform-example"
  }

  spec {
    container {
      image = "nginx:1.7.9"
      name  = "example"

      security_context {
          capabilities {
            drop = ["ALL"]
          }
      }

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
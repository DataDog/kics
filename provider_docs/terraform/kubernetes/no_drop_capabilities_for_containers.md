---
title: "No Drop Capabilities for Containers"
meta:
  name: "terraform/no_drop_capabilities_for_containers"
  id: "21cef75f-289f-470e-8038-c7cee0664164"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/no_drop_capabilities_for_containers`
**Id:** `21cef75f-289f-470e-8038-c7cee0664164`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Best Practices
## Description
Sees if Kubernetes Drop Capabilities exists to ensure containers security context

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#drop)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "test1" {
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
             add = ["NET_BIND_SERVICE"]
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

```

```kubernetes

resource "kubernetes_pod" "test3" {
  metadata {
    name = "terraform-example"
  }

  spec {

    container =  [
      {
        image = "nginx:1.7.9"
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


```

```kubernetes

resource "kubernetes_pod" "test2" {
  metadata {
    name = "terraform-example"
  }

  spec {

    container =  [
      {
        image = "nginx:1.7.9"
        name  = "example"

        security_context = {
          allow_privilege_escalation = false
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

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "negative4" {
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

```
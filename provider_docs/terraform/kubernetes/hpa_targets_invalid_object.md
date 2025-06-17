---
title: "HPA Targets Invalid Object"
meta:
  name: "terraform/hpa_targets_invalid_object"
  id: "17e52ca3-ddd0-4610-9d56-ce107442e110"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Availability"
---
## Metadata
**Name:** `terraform/hpa_targets_invalid_object`
**Id:** `17e52ca3-ddd0-4610-9d56-ce107442e110`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Availability
## Description
The Horizontal Pod Autoscaler must target a valid object

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/horizontal_pod_autoscaler#metric)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_horizontal_pod_autoscaler" "example" {
  metadata {
    name = "test"
  }

  spec {
    min_replicas = 50
    max_replicas = 100

    scale_target_ref {
      kind = "Deployment"
      name = "MyApp"
    }

    metric {
      type = "External"
      external {
        metric {
          name = "latency"
          selector {
            match_labels = {
              lb_name = "test"
            }
          }
        }
        target {
          type  = "Value"
          value = "100"
        }
      }
    }
  }
}

resource "kubernetes_horizontal_pod_autoscaler" "example2" {
  metadata {
    name = "test"
  }

  spec {
    min_replicas = 50
    max_replicas = 100

    scale_target_ref {
      kind = "Deployment"
      name = "MyApp"
    }

    metric {
      type = "Object"
      object {
        target {
          type  = "Value"
          value = "100"
        }
      }
    }
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_horizontal_pod_autoscaler" "example5" {
  metadata {
    name = "test"
  }

  spec {
    min_replicas = 50
    max_replicas = 100

    scale_target_ref {
      kind = "Deployment"
      name = "MyApp"
    }

    metric {
      type = "Object"
      object {
        metric {
          name = "latency"
        }
        described_object {
          name = "main-route"
          api_version = "networking.k8s.io/v1beta1"
          kind = "Ingress"
        }
        target {
          type  = "Value"
          value = "100"
        }
      }
    }
  }
}

```
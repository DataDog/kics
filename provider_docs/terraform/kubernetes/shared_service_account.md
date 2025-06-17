---
title: "Shared Service Account"
meta:
  name: "terraform/shared_service_account"
  id: "f74b9c43-161a-4799-bc95-0b0ec81801b9"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Secret Management"
---
## Metadata
**Name:** `terraform/shared_service_account`
**Id:** `f74b9c43-161a-4799-bc95-0b0ec81801b9`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Secret Management
## Description
A Service Account token is shared between workloads

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/pod#service_account_name)

## Non-Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "with_pod_affinity" {
  metadata {
    name = "with-pod-affinity"
  }

  spec {
    affinity {
      pod_affinity {
        required_during_scheduling_ignored_during_execution {
          label_selector {
            match_expressions {
              key      = "security"
              operator = "In"
              values   = ["S1"]
            }
          }

          topology_key = "failure-domain.beta.kubernetes.io/zone"
        }
      }

      pod_anti_affinity {
        preferred_during_scheduling_ignored_during_execution {
          weight = 100

          pod_affinity_term {
            label_selector {
              match_expressions {
                key      = "security"
                operator = "In"
                values   = ["S2"]
              }
            }

            topology_key = "failure-domain.beta.kubernetes.io/zone"
          }
        }
      }
    }

    container {
      name  = "with-pod-affinity"
      image = "k8s.gcr.io/pause:2.0"
    }

    service_account_name = "terraform-example"
  }
}

resource "kubernetes_service_account" "terraform-example" {
  metadata {
    name = "terraform-example"
  }
  secret {
    name = kubernetes_secret.example.metadata.0.name
  }
}

resource "kubernetes_secret" "example" {
  metadata {
    name = "terraform-example"
  }
}

```

## Compliant Code Examples
```kubernetes
resource "kubernetes_pod" "with_pod_affinity_1" {
  metadata {
    name = "with-pod-affinity-1"
  }

  spec {
    affinity {
      pod_affinity {
        required_during_scheduling_ignored_during_execution {
          label_selector {
            match_expressions {
              key      = "security"
              operator = "In"
              values   = ["S1"]
            }
          }

          topology_key = "failure-domain.beta.kubernetes.io/zone"
        }
      }

      pod_anti_affinity {
        preferred_during_scheduling_ignored_during_execution {
          weight = 100

          pod_affinity_term {
            label_selector {
              match_expressions {
                key      = "security"
                operator = "In"
                values   = ["S2"]
              }
            }

            topology_key = "failure-domain.beta.kubernetes.io/zone"
          }
        }
      }
    }

    container {
      name  = "with-pod-affinity"
      image = "k8s.gcr.io/pause:2.0"
    }
  }
}

resource "kubernetes_pod" "with_pod_affinity_2" {
  metadata {
    name = "with-pod-affinity-2"
  }

  spec {
    affinity {
      pod_affinity {
        required_during_scheduling_ignored_during_execution {
          label_selector {
            match_expressions {
              key      = "security"
              operator = "In"
              values   = ["S1"]
            }
          }

          topology_key = "failure-domain.beta.kubernetes.io/zone"
        }
      }

      pod_anti_affinity {
        preferred_during_scheduling_ignored_during_execution {
          weight = 100

          pod_affinity_term {
            label_selector {
              match_expressions {
                key      = "security"
                operator = "In"
                values   = ["S2"]
              }
            }

            topology_key = "failure-domain.beta.kubernetes.io/zone"
          }
        }
      }
    }

    container {
      name  = "with-pod-affinity"
      image = "k8s.gcr.io/pause:2.0"
    }

    service_account_name = "service-name"
  }
}

resource "kubernetes_service_account" "example" {
  metadata {
    name = "example"
  }
  secret {
    name = kubernetes_secret.example.metadata.0.name
  }
}

resource "kubernetes_secret" "example" {
  metadata {
    name = "terraform-example"
  }
}

```
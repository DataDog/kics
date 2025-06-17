---
title: "Cluster Labels Disabled"
meta:
  name: "terraform/cluster_labels_disabled"
  id: "65c1bc7a-4835-4ac4-a2b6-13d310b0648d"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/cluster_labels_disabled`
**Id:** `65c1bc7a-4835-4ac4-a2b6-13d310b0648d`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Insecure Configurations
## Description
Kubernetes clusters should be configured with labels by defining the `resource_labels` attribute in the `google_container_cluster` resource. Missing cluster labels make it harder to organize, identify, and apply policies to Kubernetes environments, potentially leading to management issues and security policy gaps. To mitigate this, clusters must include the `resource_labels` block as shown below:

```
resource "google_container_cluster" "example" {
  name               = "my-cluster"
  location           = "us-central1-a"
  initial_node_count = 3

  resource_labels {
    environment = "production"
    team        = "devops"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```gcp
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```

## Compliant Code Examples
```gcp
#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  resource_labels {
      
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```
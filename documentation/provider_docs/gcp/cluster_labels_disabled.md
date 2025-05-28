---
title: "Cluster Labels Disabled"
meta:
  name: "gcp/cluster_labels_disabled"
  id: "65c1bc7a-4835-4ac4-a2b6-13d310b0648d"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/cluster_labels_disabled`

**Id:** `65c1bc7a-4835-4ac4-a2b6-13d310b0648d`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Insecure Configurations

## Description
Kubernetes Clusters must be configured with labels, which means the attribute 'resource_labels' must be defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```terraform
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
```terraform
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
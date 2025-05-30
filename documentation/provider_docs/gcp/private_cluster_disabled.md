---
title: "Private Cluster Disabled"
meta:
  name: "gcp/private_cluster_disabled"
  id: "6ccb85d7-0420-4907-9380-50313f80946b"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/private_cluster_disabled`

**Id:** `6ccb85d7-0420-4907-9380-50313f80946b`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Kubernetes Clusters must be created with Private Clusters enabled, meaning the 'private_cluster_config' must be defined and the attributes 'enable_private_nodes' and 'enable_private_endpoint' must be true

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "positive1" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive2" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_endpoint = true
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive3" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_nodes = true
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive4" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {

  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive5" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_endpoint = false
        enable_private_nodes = true
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive6" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_endpoint = true
        enable_private_nodes = false
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive7" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_endpoint = false
        enable_private_nodes = false
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```

## Compliant Code Examples
```terraform
resource "google_container_cluster" "negative1" {
  name = "marcellus-wallace"
  location = "us-central1-a"
  initial_node_count = 3
  private_cluster_config {
        enable_private_endpoint = true
        enable_private_nodes = true
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```
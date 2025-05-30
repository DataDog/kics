---
title: "Google Container Node Pool Auto Repair Disabled"
meta:
  name: "gcp/google_container_node_pool_auto_repair_disabled"
  id: "acfdbec6-4a17-471f-b412-169d77553332"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/google_container_node_pool_auto_repair_disabled`

**Id:** `acfdbec6-4a17-471f-b412-169d77553332`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Google Container Node Pool Auto Repair should be enabled. This service periodically checks for failing nodes and repairs them to ensure a smooth running state.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool)

## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "positive1" {
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "positive2" {
  name       = "my-node-pool"
  location   = "us-central1"
  cluster    = google_container_cluster.primary.name
  node_count = 1

  management {
    auto_repair  = false
  }
}

resource "google_container_node_pool" "positive3" {
  name       = "my-node-pool"
  location   = "us-central1"
  cluster    = google_container_cluster.primary.name
  node_count = 1
}
```

## Compliant Code Examples
```terraform
resource "google_container_cluster" "negative1" {
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "negative2" {
  name       = "my-node-pool"
  location   = "us-central1"
  cluster    = google_container_cluster.primary.name
  node_count = 1

  management {
    auto_repair  = true
  }
}

```
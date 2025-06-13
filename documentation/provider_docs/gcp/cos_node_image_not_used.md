---
title: "COS Node Image Not Used"
meta:
  name: "gcp/cos_node_image_not_used"
  id: "8a893e46-e267-485a-8690-51f39951de58"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/cos_node_image_not_used`

**Id:** `8a893e46-e267-485a-8690-51f39951de58`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Insecure Configurations

## Description
The node image should be Container-Optimized OS(COS)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool#node_config)

## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "positive1" {
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1
}


resource "google_container_node_pool" "positive2" {
  project = "gcp_project"
  name    = "primary-pool"
  region  = "us-west1"
  cluster = google_container_cluster.primary.name

  node_config {
    image_type   = "WINDOWS_LTSC"
  }
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
  project = "gcp_project"
  name    = "primary-pool"
  region  = "us-west1"
  cluster = google_container_cluster.primary.name

  node_config {
    image_type   = "COS"
  }
}

 resource "google_container_node_pool" "negative3" {
  project = "gcp_project"
  name    = "primary-pool2"
  region  = "us-west1"
  cluster = google_container_cluster.primary.name
 }

resource "google_container_node_pool" "negative4" {
  project = "gcp_project"
  name    = "primary-pool2"
  region  = "us-west1"
  cluster = google_container_cluster.primary.name

  node_config {
    image_type   = "COS_CONTAINERD"
  }
}
```
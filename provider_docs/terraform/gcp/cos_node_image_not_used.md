---
title: "COS Node Image Not Used"
meta:
  name: "terraform/cos_node_image_not_used"
  id: "8a893e46-e267-485a-8690-51f39951de58"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/cos_node_image_not_used`
**Id:** `8a893e46-e267-485a-8690-51f39951de58`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Insecure Configurations
## Description
The node image type should be set to Container-Optimized OS (COS) to enhance security and streamline workloads in Google Kubernetes Engine (GKE). Using other image types, such as `WINDOWS_LTSC` or failing to specify the `image_type` attribute, can introduce unnecessary vulnerabilities or increase the attack surface by including unneeded components. To ensure nodes use a hardened and minimal OS, configure the `image_type` field in your node pool's `node_config` block to `"COS"` or `"COS_CONTAINERD"`:

```
node_config {
  image_type = "COS"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool#node_config)

## Non-Compliant Code Examples
```gcp
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
```gcp
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
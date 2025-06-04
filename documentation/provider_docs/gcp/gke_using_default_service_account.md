---
title: "GKE Using Default Service Account"
meta:
  name: "gcp/gke_using_default_service_account"
  id: "1c8eef02-17b1-4a3e-b01d-dcc3292d2c38"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Defaults"
---

## Metadata
**Name:** `gcp/gke_using_default_service_account`

**Id:** `1c8eef02-17b1-4a3e-b01d-dcc3292d2c38`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Defaults

## Description
Kubernetes Engine Clusters should not be configured to use the default service account

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#node_config)

## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "positive2" {
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1

  node_config {
    service_account = google_service_account.default.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
    labels = {
      foo = "bar"
    }
    tags = ["foo", "bar"]
  }
  timeouts {
    create = "30m"
    update = "40m"
  }
}

```

```terraform
resource "google_container_cluster" "positive1" {
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
    labels = {
      foo = "bar"
    }
    tags = ["foo", "bar"]
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
  name     = "my-gke-cluster"
  location = "us-central1"
  remove_default_node_pool = true
  initial_node_count       = 1

  node_config {
    service_account = google_service_account.myserviceaccount.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
    labels = {
      foo = "bar"
    }
    tags = ["foo", "bar"]
  }
  timeouts {
    create = "30m"
    update = "40m"
  }
}

```
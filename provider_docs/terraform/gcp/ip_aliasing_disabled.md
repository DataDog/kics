---
title: "IP Aliasing Disabled"
meta:
  name: "terraform/ip_aliasing_disabled"
  id: "c606ba1d-d736-43eb-ac24-e16108f3a9e0"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/ip_aliasing_disabled`
**Id:** `c606ba1d-d736-43eb-ac24-e16108f3a9e0`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Kubernetes clusters on Google Kubernetes Engine (GKE) should be created with Alias IP ranges enabled by defining the `ip_allocation_policy` block and setting the `networking_mode` attribute to `"VPC_NATIVE"`. Without these settings, as shown in the example below, the cluster may fall back to legacy networking modes, which do not provide the same isolation or scalability benefits and increase the risk of network conflicts and unauthorized access:

```
resource "google_container_cluster" "example" {
  name               = "my-legacy-cluster"
  location           = "us-central1-a"
  initial_node_count = 3
  // Missing ip_allocation_policy and/or incorrect networking_mode
}
```

A secure configuration should include both:

```
resource "google_container_cluster" "example" {
  name               = "my-secure-cluster"
  location           = "us-central1-a"
  initial_node_count = 3
  ip_allocation_policy {}
  networking_mode = "VPC_NATIVE"
}
```

Failure to enforce alias IP allocation can lead to reduced network segmentation and potentially exposes pods and services to unintended network access.

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

resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  networking_mode = "VPC_NATIVE"

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive3" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  ip_allocation_policy {

  }
  networking_mode = "ROUTES"

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
  ip_allocation_policy {

  }
  networking_mode = "VPC_NATIVE"

  timeouts {
    create = "30m"
    update = "40m"
  }
}
```
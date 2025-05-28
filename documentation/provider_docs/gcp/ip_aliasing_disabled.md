---
title: "IP Aliasing Disabled"
meta:
  name: "gcp/ip_aliasing_disabled"
  id: "c606ba1d-d736-43eb-ac24-e16108f3a9e0"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/ip_aliasing_disabled`

**Id:** `c606ba1d-d736-43eb-ac24-e16108f3a9e0`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Kubernetes Clusters must be created with Alias IP ranges enabled, which means the attribut 'ip_allocation_policy' must be defined and, if defined, the attribute 'networking_mode' must be VPC_NATIVE 

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
```terraform
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
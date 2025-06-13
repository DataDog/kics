---
title: "Shielded GKE Nodes Disabled"
meta:
  name: "gcp/shielded_gke_nodes_disabled"
  id: "579a0727-9c29-4d58-8195-fc5802a8bdb4"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/shielded_gke_nodes_disabled`

**Id:** `579a0727-9c29-4d58-8195-fc5802a8bdb4`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
GKE cluster nodes must be launched with Shielded VM enabled, which means the attribute 'enable_shielded_nodes' must be set to 'true'.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#enable_shielded_nodes)

## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "false" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = false
}
```

## Compliant Code Examples
```terraform
resource "google_container_cluster" "negative1" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = true
}
```
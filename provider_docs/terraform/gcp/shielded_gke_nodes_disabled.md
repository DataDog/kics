---
title: "Shielded GKE Nodes Disabled"
meta:
  name: "terraform/shielded_gke_nodes_disabled"
  id: "579a0727-9c29-4d58-8195-fc5802a8bdb4"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/shielded_gke_nodes_disabled`
**Id:** `579a0727-9c29-4d58-8195-fc5802a8bdb4`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
GKE cluster nodes should be launched with Shielded VM enabled by setting the `enable_shielded_nodes` attribute to `true` in the `google_container_cluster` resource. Failing to enable Shielded VM features exposes cluster nodes to potential rootkit and boot-level malware attacks, as these features help ensure node integrity through secure boot and trusted platform module (TPM) protections. For secure configuration, use:

```
resource "google_container_cluster" "secure" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = true
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#enable_shielded_nodes)

## Non-Compliant Code Examples
```gcp
resource "google_container_cluster" "false" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = false
}
```

## Compliant Code Examples
```gcp
resource "google_container_cluster" "negative1" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = true
}
```
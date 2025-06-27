---
title: "Shielded GKE Nodes Disabled"
meta:
  name: "gcp/shielded_gke_nodes_disabled"
  id: "579a0727-9c29-4d58-8195-fc5802a8bdb4"
  display_name: "Shielded GKE Nodes Disabled"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Name:** `gcp/shielded_gke_nodes_disabled`

**Query Name** `Shielded GKE Nodes Disabled`

**Id:** `579a0727-9c29-4d58-8195-fc5802a8bdb4`

**Cloud Provider:** gcp

**Platform** Terraform

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


## Compliant Code Examples
```terraform
resource "google_container_cluster" "negative1" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = true
}
```
## Non-Compliant Code Examples
```terraform
resource "google_container_cluster" "false" {
  name                  = "my-gke-cluster"
  location              = "us-central1"
  enable_shielded_nodes = false
}
```
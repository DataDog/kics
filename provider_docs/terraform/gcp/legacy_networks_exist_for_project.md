---
title: "Ensure legacy networks do not exist for a project"
meta:
  name: "terraform/legacy_networks_exist_for_project"
  id: "550e8400-e29b-41d4-a716-446655440000"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/legacy_networks_exist_for_project`
**Id:** `550e8400-e29b-41d4-a716-446655440000`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Networking and Firewall
## Description
Legacy networks in Google Cloud Platform (GCP) with auto_create_subnetworks enabled represent a significant security risk as they automatically create subnets with predefined IP ranges in every region. This can lead to overly permissive network configurations and potentially expose internal services to unauthorized access or lateral movement within your infrastructure.

The secure configuration (as shown below) explicitly avoids enabling auto-created subnetworks, giving administrators complete control over subnet creation and IP addressing:
```hcl
resource "google_compute_network" "legacy_network_2" {
  name = "legacy-network"
}
```

Insecure configuration example with auto_create_subnetworks enabled:
```hcl
resource "google_compute_network" "legacy_network" {
  name                    = "legacy-network"
  auto_create_subnetworks = true
}
```

#### Learn More

 - [Provider Reference](https://cloud.google.com/vpc/docs/legacy)

## Non-Compliant Code Examples
```gcp
resource "google_compute_network" "legacy_network" {
  name                    = "legacy-network"
  auto_create_subnetworks = true
}

```

## Compliant Code Examples
```gcp
resource "google_compute_network" "legacy_network_2" {
  name = "legacy-network"
}

```

```gcp
resource "google_compute_network" "modern_network" {
  name                    = "modern-network"
  auto_create_subnetworks = false
}

```
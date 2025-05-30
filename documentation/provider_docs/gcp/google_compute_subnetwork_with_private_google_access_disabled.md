---
title: "Google Compute Subnetwork with Private Google Access Disabled"
meta:
  name: "gcp/google_compute_subnetwork_with_private_google_access_disabled"
  id: "ee7b93c1-b3f8-4a3b-9588-146d481814f5"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/google_compute_subnetwork_with_private_google_access_disabled`

**Id:** `ee7b93c1-b3f8-4a3b-9588-146d481814f5`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Networking and Firewall

## Description
Google Compute Subnetwork should have Private Google Access enabled, which means 'private_ip_google_access' should be set to true

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_subnetwork#private_ip_google_access)

## Non-Compliant Code Examples
```terraform
resource "google_compute_subnetwork" "positive2" {
  name          = "test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
  private_ip_google_access = false
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network"
  auto_create_subnetworks = false
}

```

```terraform
resource "google_compute_subnetwork" "positive1" {
  name          = "test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network"
  auto_create_subnetworks = false
}

```

## Compliant Code Examples
```terraform
resource "google_compute_subnetwork" "negative1" {
  name          = "test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
  secondary_ip_range {
    range_name    = "tf-test-secondary-range-update1"
    ip_cidr_range = "192.168.10.0/24"
  }
  private_ip_google_access = true
}

resource "google_compute_network" "custom-test" {
  name                    = "test-network"
  auto_create_subnetworks = false
}

```
---
title: "VM With Full Cloud Access"
meta:
  name: "gcp/vm_with_full_cloud_access"
  id: "bc280331-27b9-4acb-a010-018e8098aa5d"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `gcp/vm_with_full_cloud_access`

**Id:** `bc280331-27b9-4acb-a010-018e8098aa5d`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Access Control

## Description
A VM instance is configured to use the default service account with full access to all Cloud APIs

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance#scopes)

## Non-Compliant Code Examples
```terraform
resource "google_compute_instance" "positive1" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral IP
    }
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro", "cloud-platform"]
  }
}
```

## Compliant Code Examples
```terraform
resource "google_compute_instance" "negative1" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral IP
    }
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}
```
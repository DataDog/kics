---
title: "OSLogin Is Disabled For VM Instance"
meta:
  name: "gcp/os_login_is_disabled_for_vm_instance"
  id: "d0b4d550-c001-46c3-bbdb-d5d75d33f05f"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/os_login_is_disabled_for_vm_instance`

**Id:** `d0b4d550-c001-46c3-bbdb-d5d75d33f05f`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Check if any VM instance disables OSLogin

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_instance)

## Non-Compliant Code Examples
```terraform
resource "google_compute_instance" "positive1" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }

  metadata = {
    #... some other metadata

    enable-oslogin = "FALSE"
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

```

## Compliant Code Examples
```terraform
resource "google_compute_instance" "negative1" {
  name         = "test"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  tags = ["foo", "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  // Local SSD disk
  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }

  metadata = {
    #... some other metadata

    # or if not undefined
    enable-oslogin = true
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

```
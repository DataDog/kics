---
title: "Serial Ports Are Enabled For VM Instances"
meta:
  name: "gcp/vm_serial_ports_are_enabled_for_vm_instances"
  id: "97fa667a-d05b-4f16-9071-58b939f34751"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/vm_serial_ports_are_enabled_for_vm_instances`

**Id:** `97fa667a-d05b-4f16-9071-58b939f34751`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Google Compute Engine VM instances should not enable serial ports. When enabled, anyone can access your VM, if they know the username, project ID, SSH key, instance name and zone

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

  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
    }
  }

  metadata = {
    serial-port-enable = true
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_project_metadata" "positive2" {
  metadata = {
    serial-port-enable = "TRUE"
  }
}

resource "google_compute_project_metadata_item" "positive3" {
  key   = "serial-port-enable"
  value = "TRUE"
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

  scratch_disk {
    interface = "SCSI"
  }

  network_interface {
    network = "default"

    access_config {
    }
  }

  metadata = {
    serial-port-enable = "FALSE"
  }

  metadata_startup_script = "echo hi > /test.txt"

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_project_metadata" "negative2" {
  metadata = {
    serial-port-enable = false
  }
}

resource "google_compute_project_metadata_item" "negative3" {
  key   = "my_metadata"
  value = "my_value"
}

```
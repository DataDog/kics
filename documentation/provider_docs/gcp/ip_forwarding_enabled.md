---
title: "IP Forwarding Enabled"
meta:
  name: "gcp/ip_forwarding_enabled"
  id: "f34c0c25-47b4-41eb-9c79-249b4dd47b89"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/ip_forwarding_enabled`

**Id:** `f34c0c25-47b4-41eb-9c79-249b4dd47b89`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Instances must not have IP forwarding enabled, which means the attribute 'can_ip_forward' must not be true

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/compute_instance)

## Non-Compliant Code Examples
```terraform
resource "google_compute_instance" "appserver" {
  name           = "primary-application-server"
  machine_type   = "e2-medium"
  can_ip_forward = true

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

```

## Compliant Code Examples
```terraform
resource "google_compute_instance" "appserver" {
  name           = "primary-application-server"
  machine_type   = "e2-medium"
  can_ip_forward = false

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

```
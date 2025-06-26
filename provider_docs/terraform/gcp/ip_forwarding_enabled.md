---
title: "IP Forwarding Enabled"
meta:
  name: "terraform/ip_forwarding_enabled"
  id: "f34c0c25-47b4-41eb-9c79-249b4dd47b89"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/ip_forwarding_enabled`
**Id:** `f34c0c25-47b4-41eb-9c79-249b4dd47b89`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
This check ensures that the `can_ip_forward` attribute for Google Compute Engine instances is set to `false`, preventing instances from forwarding packets that are not addressed to themselves. If `can_ip_forward` is set to `true`, as shown below, the instance could be misused as a routing or proxy device, increasing the risk of data exfiltration or man-in-the-middle attacks:

```
resource "google_compute_instance" "appserver" {
  name           = "primary-application-server"
  machine_type   = "e2-medium"
  can_ip_forward = true
  ...
}
```

To mitigate this risk, configure the attribute as `false`:

```
resource "google_compute_instance" "appserver" {
  name           = "primary-application-server"
  machine_type   = "e2-medium"
  can_ip_forward = false
  ...
}
```
Disabling IP forwarding hardens network boundaries and reduces the attack surface of the cloud environment.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/compute_instance)

## Non-Compliant Code Examples
```gcp
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
```gcp
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
---
title: "IP Forwarding Enabled"
meta:
  name: "gcp/ip_forwarding_enabled"
  id: "f34c0c25-47b4-41eb-9c79-249b4dd47b89"
  display_name: "IP Forwarding Enabled"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Name:** `gcp/ip_forwarding_enabled`

**Query Name** `IP Forwarding Enabled`

**Id:** `f34c0c25-47b4-41eb-9c79-249b4dd47b89`

**Cloud Provider:** gcp

**Platform** Terraform

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
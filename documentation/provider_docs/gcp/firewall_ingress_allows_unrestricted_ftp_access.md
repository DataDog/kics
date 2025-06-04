---
title: "Google compute firewall ingress allows unrestricted FTP access"
meta:
  name: "gcp/firewall_ingress_allows_unrestricted_ftp_access"
  id: "d3f8e9c1-7a2b-4d5f-90e2-123456789abc"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/firewall_ingress_allows_unrestricted_ftp_access`

**Id:** `d3f8e9c1-7a2b-4d5f-90e2-123456789abc`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Networking and Firewall

## Description
The firewall should not allow ingress from 0.0.0.0/0 on port 21 (FTP). Unrestricted access can lead to unauthorized access.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall)

## Non-Compliant Code Examples
```terraform
resource "google_compute_firewall" "bad_example" {
  name    = "bad-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["21"]
  }

  source_ranges = ["0.0.0.0/0"] # Unrestricted ingress for FTP
}

```

## Compliant Code Examples
```terraform
resource "google_compute_firewall" "good_example" {
  name    = "good-firewall"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["21"]
  }

  source_ranges = ["192.168.1.0/24"] # Restricted ingress for FTP
}

```
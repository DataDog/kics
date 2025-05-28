---
title: "Google compute firewall ingress allows unrestricted mysql access"
meta:
  name: "gcp/firewall_ingress_allows_unrestricted_mysql_access"
  id: "d0a1b2c3-d4e5-6789-abcd-ef0123456789"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/firewall_ingress_allows_unrestricted_mysql_access`

**Id:** `d0a1b2c3-d4e5-6789-abcd-ef0123456789`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Firewall rules must not allow ingress from 0.0.0.0/0 on port 3306 (MySQL) to prevent unauthorized access.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall)

## Non-Compliant Code Examples
```terraform
resource "google_compute_firewall" "bad_example" {
  name    = "bad-firewall-mysql"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["3306"]
  }

  source_ranges = ["0.0.0.0/0"] # Unrestricted ingress for MySQL
}

```

## Compliant Code Examples
```terraform
resource "google_compute_firewall" "good_example" {
  name    = "good-firewall-mysql"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["3306"]
  }

  source_ranges = ["192.168.1.0/24"] # Restricted ingress for MySQL
}

```
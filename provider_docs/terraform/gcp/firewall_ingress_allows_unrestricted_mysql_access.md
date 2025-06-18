---
title: "Google compute firewall ingress allows unrestricted mysql access"
meta:
  name: "terraform/firewall_ingress_allows_unrestricted_mysql_access"
  id: "d0a1b2c3-d4e5-6789-abcd-ef0123456789"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/firewall_ingress_allows_unrestricted_mysql_access`
**Id:** `d0a1b2c3-d4e5-6789-abcd-ef0123456789`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Allowing ingress traffic from `0.0.0.0/0` on port 3306, as shown in the Terraform attribute `source_ranges = ["0.0.0.0/0"]`, exposes MySQL databases to the internet, making them susceptible to unauthorized access and potential attacks. This misconfiguration can lead to data breaches, data loss, or system compromise if malicious actors exploit the open MySQL port. Restricting access to trusted IP ranges, for example `source_ranges = ["192.168.1.0/24"]`, significantly reduces this risk by limiting who can attempt to connect to the database.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall)

## Non-Compliant Code Examples
```gcp
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
```gcp
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
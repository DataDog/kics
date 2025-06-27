---
title: "Google compute firewall ingress allows unrestricted mysql access"
meta:
  name: "gcp/firewall_ingress_allows_unrestricted_mysql_access"
  id: "d0a1b2c3-d4e5-6789-abcd-ef0123456789"
  display_name: "Google compute firewall ingress allows unrestricted mysql access"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `gcp/firewall_ingress_allows_unrestricted_mysql_access`
**Query Name** `Google compute firewall ingress allows unrestricted mysql access`
**Id:** `d0a1b2c3-d4e5-6789-abcd-ef0123456789`
**Cloud Provider:** gcp
**Platform** Terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Allowing ingress traffic from `0.0.0.0/0` on port 3306, as shown in the Terraform attribute `source_ranges = ["0.0.0.0/0"]`, exposes MySQL databases to the internet, making them susceptible to unauthorized access and potential attacks. This misconfiguration can lead to data breaches, data loss, or system compromise if malicious actors exploit the open MySQL port. Restricting access to trusted IP ranges, for example `source_ranges = ["192.168.1.0/24"]`, significantly reduces this risk by limiting who can attempt to connect to the database.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall)


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
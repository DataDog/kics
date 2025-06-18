---
title: "Google Compute Network Using Default Firewall Rule"
meta:
  name: "terraform/google_compute_network_using_default_firewall_rule"
  id: "40abce54-95b1-478c-8e5f-ea0bf0bb0e33"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/google_compute_network_using_default_firewall_rule`
**Id:** `40abce54-95b1-478c-8e5f-ea0bf0bb0e33`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Google Compute Network resources should avoid using the default firewall rule, as it typically allows overly permissive access to network resources and does not restrict traffic according to least-privilege principles. If left unaddressed, using the default firewall may expose internal resources to unauthorized external access, increasing the risk of lateral movement or compromise within a project. Instead, firewall rules should be defined explicitly with restricted protocols, source ranges, and tags, as shown below:

```
resource "google_compute_firewall" "secure_example" {
  name    = "restricted-firewall"
  network = google_compute_network.secure_example.name

  allow {
    protocol = "tcp"
    ports    = ["80", "443"]
  }

  source_ranges = ["203.0.113.0/24"]
}

resource "google_compute_network" "secure_example" {
  name = "test-network"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall#name)

## Non-Compliant Code Examples
```gcp
resource "google_compute_firewall" "positive1" {
  name    = "default"
  network = google_compute_network.positive1.name
}

resource "google_compute_network" "positive1" {
  name = "test-network"
}

```

## Compliant Code Examples
```gcp
resource "google_compute_firewall" "negative1" {
  name    = "test-firewall"
  network = google_compute_network.negative1.name

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["80", "8080"]
  }

  source_tags = ["web"]
}

resource "google_compute_network" "negative1" {
  name = "test-network"
}

```
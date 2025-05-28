---
title: "Google Compute Network Using Default Firewall Rule"
meta:
  name: "gcp/google_compute_network_using_default_firewall_rule"
  id: "40abce54-95b1-478c-8e5f-ea0bf0bb0e33"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `gcp/google_compute_network_using_default_firewall_rule`

**Id:** `40abce54-95b1-478c-8e5f-ea0bf0bb0e33`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Google Compute Network should not use default firewall rule

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall#name)

## Non-Compliant Code Examples
```terraform
resource "google_compute_firewall" "positive1" {
  name    = "default"
  network = google_compute_network.positive1.name
}

resource "google_compute_network" "positive1" {
  name = "test-network"
}

```

## Compliant Code Examples
```terraform
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
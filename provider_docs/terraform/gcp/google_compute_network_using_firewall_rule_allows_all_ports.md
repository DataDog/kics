---
title: "Google Compute Network Using Firewall Rule that Allows All Ports"
meta:
  name: "terraform/google_compute_network_using_firewall_rule_allows_all_ports"
  id: "22ef1d26-80f8-4a6c-8c15-f35aab3cac78"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/google_compute_network_using_firewall_rule_allows_all_ports`
**Id:** `22ef1d26-80f8-4a6c-8c15-f35aab3cac78`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Networking and Firewall
## Description
Allowing a Google Compute Network firewall rule to permit traffic on all TCP ports (using `ports = ["0-65535"]`) exposes instances to a broad range of attacks and unauthorized access, increasing the risk of exploitation across unused and potentially vulnerable services. By not restricting allowed ports to only those necessary—such as `ports = ["80", "8080"]` for web services—the configuration creates a large attack surface. To minimize security risks, firewall rules should define only the specific protocols and ports required for service functionality.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall#allow)

## Non-Compliant Code Examples
```gcp
resource "google_compute_firewall" "positive1" {
  name    = "test-firewall"
  network = google_compute_network.positive1.name

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["0-65535"]
  }

  source_tags = ["web"]
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
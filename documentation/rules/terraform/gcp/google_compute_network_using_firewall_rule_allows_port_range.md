---
title: "Google Compute Network Using Firewall Rule that Allows Port Range"
group-id: "rules/terraform/gcp"
meta:
  name: "gcp/google_compute_network_using_firewall_rule_allows_port_range"
  id: "e6f61c37-106b-449f-a5bb-81bfcaceb8b4"
  display_name: "Google Compute Network Using Firewall Rule that Allows Port Range"
  cloud_provider: "gcp"
  framework: "Terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `e6f61c37-106b-449f-a5bb-81bfcaceb8b4`

**Cloud Provider:** gcp

**Framework:** Terraform

**Severity:** Low

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_firewall#allow)

### Description

 Allowing a port range in a Google Compute Network firewall rule, such as `ports = ["80", "8080", "1000-2000"]`, can expose unnecessary services and increase the attack surface of your cloud environment. Attackers may exploit any open ports within the specified range, leading to potential unauthorized access or compromise of resources. To reduce risk, firewall rules should restrict access to only required ports, as shown in a secure configuration:

```
allow {
  protocol = "tcp"
  ports    = ["80", "8080"]
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
## Non-Compliant Code Examples
```terraform
resource "google_compute_firewall" "positive1" {
  name    = "test-firewall"
  network = google_compute_network.positive1.name

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["80", "8080", "1000-2000"]
  }

  source_tags = ["web"]
}

resource "google_compute_network" "positive1" {
  name = "test-network"
}

```
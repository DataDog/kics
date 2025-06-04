---
title: "Google Compute Subnetwork Logging Disabled"
meta:
  name: "gcp/google_compute_subnetwork_logging_disabled"
  id: "40430747-442d-450a-a34f-dc57149f4609"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `gcp/google_compute_subnetwork_logging_disabled`

**Id:** `40430747-442d-450a-a34f-dc57149f4609`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Observability

## Description
This query checks if logs are enabled for a Google Compute Subnetwork resource.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_subnetwork)

## Non-Compliant Code Examples
```terraform
resource "google_compute_subnetwork" "positive1" {
  name          = "log-test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id
}
```

## Compliant Code Examples
```terraform
resource "google_compute_subnetwork" "negative1" {
  name          = "log-test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.custom-test.id

  log_config {
    aggregation_interval = "INTERVAL_10_MIN"
    flow_sampling        = 0.5
    metadata             = "INCLUDE_ALL_METADATA"
  }
}
```
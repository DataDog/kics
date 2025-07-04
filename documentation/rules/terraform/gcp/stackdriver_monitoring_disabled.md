---
title: "Stackdriver Monitoring Disabled"
group-id: "rules/terraform/gcp"
meta:
  name: "gcp/stackdriver_monitoring_disabled"
  id: "30e8dfd2-3591-4d19-8d11-79e93106c93d"
  display_name: "Stackdriver Monitoring Disabled"
  cloud_provider: "gcp"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Id:** `30e8dfd2-3591-4d19-8d11-79e93106c93d`

**Cloud Provider:** gcp

**Framework:** Terraform

**Severity:** Medium

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#monitoring_service)

### Description

 Kubernetes Engine Clusters must have Stackdriver Monitoring enabled by setting the `monitoring_service` attribute to `"monitoring.googleapis.com/kubernetes"` or leaving it undefined, which allows Google Cloud’s default monitoring. Disabling monitoring by setting `monitoring_service = "none"` leaves clusters without visibility into performance, health, or security events, increasing the risk of undetected failures or malicious activity. 

Secure configuration should look like:

```
resource "google_container_cluster" "example" {
  name               = "secure-cluster"
  location           = "us-central1-a"
  initial_node_count = 3
  monitoring_service = "monitoring.googleapis.com/kubernetes"
}
```
If left unaddressed, this misconfiguration can prevent prompt detection and remediation of operational or security incidents, potentially leading to service outages or breaches.


## Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  monitoring_service = "monitoring.googleapis.com/kubernetes"

  timeouts {
    create = "30m"
    update = "40m"
  }
}

# Monitoring service defaults to Stackdriver, so it's okay to be undefined
resource "google_container_cluster" "negative1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  
  timeouts {
    create = "30m"
    update = "40m"
  }
}
```
## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  monitoring_service = "none"

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  monitoring_service = "monitoring.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }
}
```
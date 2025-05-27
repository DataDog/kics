---
title: "Stackdriver Logging Disabled"
meta:
  name: "gcp/stackdriver_logging_disabled"
  id: "4c7ebcb2-eae2-461e-bc83-456ee2d4f694"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `gcp/stackdriver_logging_disabled`

**Id:** `4c7ebcb2-eae2-461e-bc83-456ee2d4f694`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Observability

## Description
Kubernetes Engine Clusters must have Stackdriver Logging enabled, which means the attribute 'logging_service' must be defined and different from 'none'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#logging_service)

## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  logging_service = "none"

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  logging_service = "logging.googleapis.com"

  timeouts {
    create = "30m"
    update = "40m"
  }
}
```

## Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  logging_service = "logging.googleapis.com/kubernetes"

  timeouts {
    create = "30m"
    update = "40m"
  }
}

# Logging service defaults to Stackdriver, so it's okay to be undefined
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
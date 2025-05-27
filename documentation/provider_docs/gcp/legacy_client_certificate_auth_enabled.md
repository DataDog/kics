---
title: "Legacy Client Certificate Auth Enabled"
meta:
  name: "gcp/legacy_client_certificate_auth_enabled"
  id: "73fb21a1-b19a-45b1-b648-b47b1678681e"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/legacy_client_certificate_auth_enabled`

**Id:** `73fb21a1-b19a-45b1-b648-b47b1678681e`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Insecure Configurations

## Description
Kubernetes Clusters must use the default OAuth authentication, which means 'master_auth' must either be undefined or have 'client_certificate_config' with the attribute 'issue_client_certificate' equal to false

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    client_certificate_config {
      issue_client_certificate = true
    }
  }

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

  master_auth {
    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

# leaving the field undefined is acceptable
resource "google_container_cluster" "negative2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```
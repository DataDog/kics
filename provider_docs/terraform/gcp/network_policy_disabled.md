---
title: "Network Policy Disabled"
meta:
  name: "terraform/network_policy_disabled"
  id: "11e7550e-c4b6-472e-adff-c698f157cdd7"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/network_policy_disabled`
**Id:** `11e7550e-c4b6-472e-adff-c698f157cdd7`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Kubernetes Engine Clusters must have network policies enabled to restrict traffic between pods and control communication within the cluster. In Terraform, this means setting `network_policy.enabled = true` and ensuring `addons_config.network_policy_config.disabled = false`, as shown below:

```
network_policy {
  enabled = true
}
addons_config {
  network_policy_config {
    disabled = false
  }
}
```

If these settings are not properly configured, unauthorized traffic between pods may be allowed, increasing the risk of lateral movement and the potential compromise of sensitive applications or data within the cluster.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```gcp
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  network_policy {
      enabled = true
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
  network_policy {
      enabled = true
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive3" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive4" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  network_policy {
      enabled = true
  }
  addons_config {

  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive5" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  network_policy {
      enabled = false
  }
  addons_config {
    network_policy_config {
        disabled = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive6" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  network_policy {
      enabled = true
  }
  addons_config {
    network_policy_config {
        disabled = true
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}
```

## Compliant Code Examples
```gcp
#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  network_policy {
      enabled = true
  }
  addons_config {
    network_policy_config {
        disabled = false
    }
  }
  networking_mode = "VPC_NATIVE"

  timeouts {
    create = "30m"
    update = "40m"
  }
}
```
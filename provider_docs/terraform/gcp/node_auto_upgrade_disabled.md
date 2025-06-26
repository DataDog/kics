---
title: "Node Auto Upgrade Disabled"
meta:
  name: "terraform/node_auto_upgrade_disabled"
  id: "b139213e-7d24-49c2-8025-c18faa21ecaa"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Resource Management"
---
## Metadata
**Name:** `terraform/node_auto_upgrade_disabled`
**Id:** `b139213e-7d24-49c2-8025-c18faa21ecaa`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Resource Management
## Description
Kubernetes nodes should have automatic upgrades enabled to ensure that critical security patches, bug fixes, and feature updates are applied without manual intervention. In Terraform, this is configured with the `auto_upgrade` attribute inside the `management` block set to `true`:

```
management {
  auto_upgrade = true
}
```

If `auto_upgrade` is not enabled, as in the following example, nodes may remain outdated and vulnerable to known security flaws:

```
management {
  auto_upgrade = false
}
```

Leaving auto upgrade disabled can expose your cluster to exploits and instability due to unpatched vulnerabilities in the underlying infrastructure.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_node_pool#auto_upgrade)

## Non-Compliant Code Examples
```gcp
resource "google_container_node_pool" "positive1" {
  name       = "my-node-pool"
  location   = "us-central1-a"
  cluster    = google_container_cluster.primary.name
  node_count = 3

  timeouts {
    create = "30m"
    update = "20m"
  }
}

resource "google_container_node_pool" "positive2" {
  name       = "my-node-pool"
  location   = "us-central1-a"
  cluster    = google_container_cluster.primary.name
  node_count = 3

  management {
    auto_repair = true
  }

  timeouts {
    create = "30m"
    update = "20m"
  }
}

resource "google_container_node_pool" "positive3" {
  name       = "my-node-pool"
  location   = "us-central1-a"
  cluster    = google_container_cluster.primary.name
  node_count = 3

  management {
    auto_upgrade = false
  }

  timeouts {
    create = "30m"
    update = "20m"
  }
}

```

## Compliant Code Examples
```gcp
resource "google_container_node_pool" "negative1" {
  name       = "my-node-pool"
  location   = "us-central1-a"
  cluster    = google_container_cluster.primary.name
  node_count = 3

  management {
    auto_upgrade = true
  }

  timeouts {
    create = "30m"
    update = "20m"
  }
}
```
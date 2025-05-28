---
title: "team label missing"
meta:
  name: "gcp/team_label_not_present"
  id: "b2b3c4d5-e6f7-8901-gh23-ijkl456m7890"
  cloud_provider: "gcp"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `gcp/team_label_not_present`

**Id:** `b2b3c4d5-e6f7-8901-gh23-ijkl456m7890`

**Cloud Provider:** gcp

**Severity:** Info

**Category:** Best Practices

## Description
Ensures that every cloud resource has a 'team' label for ownership tracking.

#### Learn More

 - [Provider Reference](https://cloud.google.com/storage/docs/tags-and-labels)

## Non-Compliant Code Examples
```terraform
resource "google_bigtable_instance" "positive2" {
  name = "marcellus-wallace"
  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_compute_instance" "good_example_2" {
  name         = "good-instance"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }

  labels = {
    environment = "prod"
  }
}

```

## Compliant Code Examples
```terraform
# ✅ "team" label is not a valid attribute for this resource type
resource "google_container_cluster" "good_example" {
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

```terraform
resource "google_compute_instance" "good_example" {
  name         = "good-instance"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }

  labels = {
    Team        = "DevOps" # ✅ "Team" tag is present
    environment = "prod"
  }
}

```

```terraform
resource "google_compute_instance" "good_example" {
  name         = "good-instance"
  machine_type = "e2-medium"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }

  labels = {
    team        = "DevOps" # ✅ "team" tag is present
    environment = "prod"
  }
}

```
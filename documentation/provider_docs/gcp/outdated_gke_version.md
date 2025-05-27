---
title: "Outdated GKE Version"
meta:
  name: "gcp/outdated_gke_version"
  id: "128df7ec-f185-48bc-8913-ce756a3ccb85"
  cloud_provider: "gcp"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `gcp/outdated_gke_version`

**Id:** `128df7ec-f185-48bc-8913-ce756a3ccb85`

**Cloud Provider:** gcp

**Severity:** Low

**Category:** Best Practices

## Description
Running outdated versions of Google Kubernetes Engine (GKE) can expose it to known vulnerabilities and attacks. To reduce these risks, it is recommended to ensure that GKE is always running the latest version.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster#master_version)

## Non-Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }

  min_master_version = "1.24"
}

#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }

  
  node_version = "1.24"
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
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }

  min_master_version = "latest"
}

#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }

  min_master_version = "1.25"
}

#this code is a correct code for which the query should not find any result
resource "google_container_cluster" "negative3" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  timeouts {
    create = "30m"
    update = "40m"
  }

  min_master_version = "1.25"
  node_version = "1.25"
}

```
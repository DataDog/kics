---
title: "Legacy Client Certificate Auth Enabled"
meta:
  name: "gcp/legacy_client_certificate_auth_enabled"
  id: "73fb21a1-b19a-45b1-b648-b47b1678681e"
  display_name: "Legacy Client Certificate Auth Enabled"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "LOW"
  category: "Insecure Configurations"
---
## Metadata

**Name:** `gcp/legacy_client_certificate_auth_enabled`

**Query Name** `Legacy Client Certificate Auth Enabled`

**Id:** `73fb21a1-b19a-45b1-b648-b47b1678681e`

**Cloud Provider:** gcp

**Platform** Terraform

**Severity:** Low

**Category:** Insecure Configurations

## Description
Kubernetes clusters in Google Kubernetes Engine (GKE) should use the default OAuth authentication, ensuring that client certificates are not issued for cluster authentication. This is enforced in Terraform by setting `master_auth.client_certificate_config.issue_client_certificate` to `false` or by omitting the attribute entirely. Allowing client certificate issuance (`issue_client_certificate = true`) increases the cluster’s attack surface by enabling users to authenticate with potentially compromised or unmanaged certificates, potentially leading to unauthorized access.

For a secure configuration, ensure the relevant block in Terraform looks like:

```
master_auth {
  client_certificate_config {
    issue_client_certificate = false
  }
}
```
or is simply omitted.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)


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
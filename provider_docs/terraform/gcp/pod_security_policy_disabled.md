---
title: "Pod Security Policy Disabled"
meta:
  name: "terraform/pod_security_policy_disabled"
  id: "9192e0f9-eca5-4056-9282-ae2a736a4088"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/pod_security_policy_disabled`
**Id:** `9192e0f9-eca5-4056-9282-ae2a736a4088`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Kubernetes clusters managed by Terraform should have the Pod Security Policy (PSP) controller enabled by setting the `pod_security_policy_config { enabled = true }` attribute in the `google_container_cluster` resource. Enabling PSP helps enforce fine-grained security controls over pod behavior, reducing the risk of privilege escalation or unauthorized access within your cluster. If left unconfigured or disabled, as in `pod_security_policy_config { enabled = false }`, workloads in the cluster may bypass key security restrictions, increasing the potential attack surface. 

A secure configuration looks like:

```
resource "google_container_cluster" "example" {
  // cluster configuration
  pod_security_policy_config {
    enabled = true
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster)

## Non-Compliant Code Examples
```gcp
#this is a problematic code where the query should report a result(s)
resource "google_container_cluster" "positive1" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3

  timeouts {
    create = "30m"
    update = "40m"
  }
}

resource "google_container_cluster" "positive2" {
  name               = "marcellus-wallace"
  location           = "us-central1-a"
  initial_node_count = 3
  pod_security_policy_config {
        enabled = false
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
  pod_security_policy_config {
        enabled = "true"
  }

  timeouts {
    create = "30m"
    update = "40m"
  }
}

```
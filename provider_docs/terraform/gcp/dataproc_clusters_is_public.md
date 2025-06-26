---
title: "Dataproc Clusters Publicly Accessible"
meta:
  name: "terraform/dataproc_clusters_is_public"
  id: "e3f7a9b1-c2d3-4e5f-8901-23456789abcd"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/dataproc_clusters_is_public`
**Id:** `e3f7a9b1-c2d3-4e5f-8901-23456789abcd`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Insecure Configurations
## Description
Google Cloud Dataproc clusters should not be publicly accessible as this could allow unauthorized access to sensitive data and resources. When IAM bindings or members include public principals like 'allUsers' or 'allAuthenticatedUsers', the Dataproc cluster becomes exposed to anyone on the internet or any authenticated Google account, respectively. To secure Dataproc clusters, use specific identities in IAM policies instead of public principals. For example, use `members = ["user:someone@example.com", "group:admins@example.com"]` instead of `members = ["allAuthenticatedUsers", "user:someone@example.com"]` or `member = "allUsers"`. Limiting access to specific users and groups significantly reduces the attack surface and helps maintain the principle of least privilege.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/dataproc_cluster_iam_member)

## Non-Compliant Code Examples
```gcp
# Positive test case# Failing IAM Member example
resource "google_dataproc_cluster_iam_member" "bad_member" {
  name   = "bad-dataproc-member"
  member = "allUsers" # ❌ Public principal
}

# Failing IAM Binding example
resource "google_dataproc_cluster_iam_binding" "bad_binding" {
  name    = "bad-dataproc-binding"
  members = ["allAuthenticatedUsers", "user:someone@example.com"] # ❌ Contains public principal
}

```

## Compliant Code Examples
```gcp
# Passing IAM Member example
resource "google_dataproc_cluster_iam_member" "good_member" {
  name   = "good-dataproc-member"
  member = "user:someone@example.com" # ✅ Non-public principal
}

```

```gcp
# Passing IAM Binding example
resource "google_dataproc_cluster_iam_binding" "good_binding" {
  name    = "good-dataproc-binding"
  members = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
}

```
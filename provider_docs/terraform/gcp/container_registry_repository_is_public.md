---
title: "Container Registry Repo Is Public"
meta:
  name: "terraform/container_registry_repository_is_public"
  id: "f6a7b8c9-d0e1-2345-f678-90abcdef1234"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/container_registry_repository_is_public`
**Id:** `f6a7b8c9-d0e1-2345-f678-90abcdef1234`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Insecure Configurations
## Description
Allowing public access to Container Registry repositories creates significant security risks by exposing potentially sensitive container images and artifacts to anyone on the internet. When IAM configurations include public principals like 'allUsers' or 'allAuthenticatedUsers', it bypasses access controls and may lead to data exfiltration, intellectual property theft, or deployment of compromised containers. Instead of using public principals (e.g., `member = "allUsers"` or `members = ["allAuthenticatedUsers", ...]`), implement proper access controls by explicitly specifying authorized users and groups, such as `members = ["user:someone@example.com", "group:admins@example.com"]` to ensure only legitimate users can access your container registry resources.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_registry)

## Non-Compliant Code Examples
```gcp
# Failing IAM Member Example
resource "google_storage_bucket_iam_member" "bad_example_member" {
  bucket = "example-bucket"
  member = "allUsers" # ❌ Public principal
  role   = "roles/storage.objectViewer"
}

# Failing IAM Binding Example
resource "google_storage_bucket_iam_binding" "bad_example_binding" {
  bucket  = "example-bucket"
  members = ["allAuthenticatedUsers", "user:someone@example.com"] # ❌ Contains public principal
  role    = "roles/storage.objectViewer"
}

```

## Compliant Code Examples
```gcp
# Passing IAM Member Example
resource "google_storage_bucket_iam_member" "good_example_member" {
  bucket = "example-bucket"
  member = "user:someone@example.com" # ✅ Non-public principal
  role   = "roles/storage.objectViewer"
}

```

```gcp
# Passing IAM Binding Example
resource "google_storage_bucket_iam_binding" "good_example_binding" {
  bucket  = "example-bucket"
  members = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
  role    = "roles/storage.objectViewer"
}

```
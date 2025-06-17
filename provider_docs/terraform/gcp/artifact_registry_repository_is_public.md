---
title: "Artifact Registry Repo Is Public"
meta:
  name: "terraform/artifact_registry_repository_is_public"
  id: "a7b8c9d0-e1f2-3a4b-5c6d-7e8f90123456"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/artifact_registry_repository_is_public`
**Id:** `a7b8c9d0-e1f2-3a4b-5c6d-7e8f90123456`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
Google Cloud Artifact Registry repositories store container images, language packages, and other artifacts that may contain sensitive code or data. When IAM bindings or members include public principals like 'allUsers' or 'allAuthenticatedUsers', these artifacts become accessible to anyone on the internet or any authenticated Google account, respectively, potentially exposing proprietary code or sensitive configurations.

This security risk could lead to intellectual property theft, discovery of hardcoded secrets, or provide attackers information about your infrastructure that could be used in further attacks. To properly secure your artifact repositories, specify only the necessary users and groups that require access, as shown in the following example:

```terraform
resource "google_artifact_registry_repository_iam_binding" "good_example_binding" {
  repository = "example-repo"
  members    = ["user:someone@example.com", "group:admins@example.com"]
  role       = "roles/artifactregistry.admin"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/artifact_registry_repository_iam_member)

## Non-Compliant Code Examples
```gcp
# IAM Member violation
resource "google_artifact_registry_repository_iam_member" "bad_example_member" {
  repository = "example-repo"
  member     = "allUsers" # ❌ Public principal
  role       = "roles/artifactregistry.reader"
}

# IAM Binding violation
resource "google_artifact_registry_repository_iam_binding" "bad_example_binding" {
  repository = "example-repo"
  members    = ["allAuthenticatedUsers", "user:someone@example.com"] # ❌ Contains public principal
  role       = "roles/artifactregistry.admin"
}

```

## Compliant Code Examples
```gcp
# IAM Member compliant
resource "google_artifact_registry_repository_iam_member" "good_example_member" {
  repository = "example-repo"
  member     = "user:someone@example.com" # ✅ Non-public principal
  role       = "roles/artifactregistry.reader"
}

```

```gcp

# IAM Binding compliant
resource "google_artifact_registry_repository_iam_binding" "good_example_binding" {
  repository = "example-repo"
  members    = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
  role       = "roles/artifactregistry.admin"
}

```
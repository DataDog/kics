---
title: "Cloud Run Service Is Public"
meta:
  name: "gcp/cloud_run_service_is_public"
  id: "7e3c1a2b-9d4f-4c8e-8a5b-0f1e2d3c4b6a"
  display_name: "Cloud Run Service Is Public"
  cloud_provider: "gcp"
  framework: "Terraform"
  severity: "HIGH"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `7e3c1a2b-9d4f-4c8e-8a5b-0f1e2d3c4b6a`

**Cloud Provider:** gcp

**Framework:** Terraform

**Severity:** High

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_service_iam_member)

### Description

 Cloud Run services with IAM bindings or members that include public principals like 'allUsers' or 'allAuthenticatedUsers' expose your service to anyone on the internet, which creates a significant security risk. When a Cloud Run service is publicly accessible, it could lead to unauthorized access, data breaches, or exploitation of vulnerabilities within your application. Instead of using public principals, you should restrict access to specific users or service accounts as shown in this secure example: `members = ["user:someone@example.com", "group:admins@example.com"]` rather than the insecure approach: `members = ["allAuthenticatedUsers", "user:someone@example.com"]`.


## Compliant Code Examples
```terraform
# Passing Terraform Example for IAM Binding
resource "google_cloud_run_service_iam_binding" "good_example_binding" {
  service = "my-cloud-run-service"
  members = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
  role    = "roles/run.invoker"
}

```

```terraform
# Passing Terraform Example for IAM Member
resource "google_cloud_run_service_iam_member" "good_example_member" {
  service = "my-cloud-run-service"
  member  = "user:someone@example.com" # ✅ Non-public principal
  role    = "roles/run.invoker"
}


```
## Non-Compliant Code Examples
```terraform
# Failing Terraform Example for IAM Member
resource "google_cloud_run_service_iam_member" "bad_example_member" {
  service = "my-cloud-run-service"
  member  = "allUsers" # ❌ Public principal
  role    = "roles/run.invoker"
}

# Failing Terraform Example for IAM Binding
resource "google_cloud_run_service_iam_binding" "bad_example_binding" {
  service = "my-cloud-run-service"
  members = ["allAuthenticatedUsers", "user:someone@example.com"] # ❌ Contains public principal
  role    = "roles/run.invoker"
}

```
---
title: "Google Project IAM Binding Service Account has Token Creator or Account User Role"
meta:
  name: "terraform/google_project_iam_binding_service_account_has_token_creator_or_account_user_role"
  id: "617ef6ff-711e-4bd7-94ae-e965911b1b40"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/google_project_iam_binding_service_account_has_token_creator_or_account_user_role`
**Id:** `617ef6ff-711e-4bd7-94ae-e965911b1b40`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
This check identifies if a Google Project IAM binding grants 'serviceAccountTokenCreator' or 'serviceAccountUser' roles, which provide excessive privileges to service accounts. These roles allow the specified members to impersonate service accounts, creating security risks through privilege escalation paths and potential lateral movement across your GCP environment.

Members with the 'serviceAccountTokenCreator' role can create OAuth2 access tokens, while those with 'serviceAccountUser' role can use service accounts for their operations. To remediate this issue, assign more restrictive roles that follow the principle of least privilege instead, as shown in the secure example:

```
resource "google_project_iam_binding" "secure_example" {
  project = "your-project-id"
  role    = "roles/editor"

  members = [
    "user:jane@example.com",
  ]
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/google_project_iam#google_project_iam_binding)

## Non-Compliant Code Examples
```gcp
resource "google_project_iam_binding" "positive1" {
  project = "your-project-id"
  role    = "roles/iam.serviceAccountTokenCreator"

  members = [
    "user:jane@example.com",
    "serviceAccount:my-other-app@appspot.gserviceacccount.com"
  ]
}

resource "google_project_iam_binding" "positive2" {
  project = "your-project-id"
  role    = "roles/iam.serviceAccountTokenCreator"
  member = "serviceAccount:my-other-app@appspot.gserviceacccount.com"
}

resource "google_project_iam_binding" "positive3" {
  project = "your-project-id"
  role    = "roles/iam.serviceAccountUser"

  members = [
    "user:jane@example.com",
    "serviceAccount:my-other-app@appspot.gserviceacccount.com"
  ]
}

resource "google_project_iam_binding" "positive4" {
  project = "your-project-id"
  role    = "roles/iam.serviceAccountUser"
  member = "serviceAccount:my-other-app@appspot.gserviceacccount.com"
}
```

## Compliant Code Examples
```gcp
resource "google_project_iam_binding" "negative1" {
  project = "your-project-id"
  role    = "roles/editor"

  members = [
    "user:jane@example.com",
  ]
}
```
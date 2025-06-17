---
title: "Cloud Storage Anonymous or Publicly Accessible"
meta:
  name: "terraform/cloud_storage_anonymous_or_publicly_accessible"
  id: "a6cd52a1-3056-4910-96a5-894de9f3f3b3"
  cloud_provider: "terraform"
  severity: "CRITICAL"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/cloud_storage_anonymous_or_publicly_accessible`
**Id:** `a6cd52a1-3056-4910-96a5-894de9f3f3b3`
**Cloud Provider:** terraform
**Severity:** Critical
**Category:** Access Control
## Description
Cloud Storage Buckets configured with anonymous or public access create significant security risks by allowing anyone on the internet to access potentially sensitive data. When 'allUsers' is included in IAM bindings, it grants access to anyone, while 'allAuthenticatedUsers' grants access to anyone with a Google account, both creating overly permissive access controls that violate the principle of least privilege.

Insecure configuration example:
```
resource "google_storage_bucket_iam_binding" "insecure" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = ["user:jane@example.com", "allUsers"]
}
```

Secure configuration example:
```
resource "google_storage_bucket_iam_binding" "secure" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = ["user:jane@example.com"]
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket_iam#google_storage_bucket_iam_binding)

## Non-Compliant Code Examples
```gcp
#this is a problematic code where the query should report a result(s)
resource "google_storage_bucket_iam_binding" "positive1" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = []
}

resource "google_storage_bucket_iam_binding" "positive2" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = ["user:jane@example.com","allUsers"]
}

resource "google_storage_bucket_iam_binding" "positive3" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = ["user:jane@example.com", "allAuthenticatedUsers"]
}
```

## Compliant Code Examples
```gcp
#this code is a correct code for which the query should not find any result
resource "google_storage_bucket_iam_binding" "negative1" {
  bucket = google_storage_bucket.default.name
  role = "roles/storage.admin"
  members = [
    "user:jane@example.com",
  ]
}
```
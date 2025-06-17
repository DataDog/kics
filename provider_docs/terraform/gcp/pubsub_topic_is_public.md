---
title: "Pub/Sub Topics are anonymously or publicly accessible"
meta:
  name: "terraform/pubsub_topic_is_public"
  id: "7sdj7dsj8-f348-4f95-845c-1090e41a615c"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/pubsub_topic_is_public`
**Id:** `7sdj7dsj8-f348-4f95-845c-1090e41a615c`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
Google Cloud Pub/Sub Topics should not be configured to allow public access by assigning IAM roles to the special IAM principals `allUsers` or `allAuthenticatedUsers`. Granting roles to these principals, as in the example below,

```
resource "google_pubsub_topic_iam_member" "bad_example" {
  topic  = "example-topic"
  member = "allUsers"
  role   = "roles/pubsub.publisher"
}
```

makes the topic accessible to anyone on the internet or to any authenticated Google user, exposing your data to unauthorized access or misuse. To prevent this, restrict the `member` attribute to specific users or service accounts, for example:

```
resource "google_pubsub_topic_iam_member" "good_example" {
  topic  = "example-topic"
  member = "user:someone@example.com"
  role   = "roles/pubsub.publisher"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_topic_iam_member)

## Non-Compliant Code Examples
```gcp
# IAM Member violation
resource "google_pubsub_topic_iam_member" "bad_example_member" {
  topic  = "example-topic"
  member = "allUsers" # ❌ Public principal
  role   = "roles/pubsub.publisher"
}

# IAM Binding violation
resource "google_pubsub_topic_iam_binding" "bad_example_binding" {
  topic   = "example-topic"
  members = ["allAuthenticatedUsers", "user:someone@example.com"] # ❌ Contains public principal
  role    = "roles/pubsub.subscriber"
}

```

## Compliant Code Examples
```gcp
# IAM Binding compliant
resource "google_pubsub_topic_iam_binding" "good_example_binding" {
  topic   = "example-topic"
  members = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
  role    = "roles/pubsub.subscriber"
}

```

```gcp
# IAM Member compliant
resource "google_pubsub_topic_iam_member" "good_example_member" {
  topic  = "example-topic"
  member = "user:someone@example.com" # ✅ Non-public principal
  role   = "roles/pubsub.publisher"
}

```
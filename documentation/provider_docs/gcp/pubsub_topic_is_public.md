---
title: "Pub/Sub Topics are anonymously or publicly accessible"
meta:
  name: "gcp/pubsub_topic_is_public"
  id: "7sdj7dsj8-f348-4f95-845c-1090e41a615c"
  cloud_provider: "gcp"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `gcp/pubsub_topic_is_public`

**Id:** `7sdj7dsj8-f348-4f95-845c-1090e41a615c`

**Cloud Provider:** gcp

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Pub/Sub Topics must not be publicly accessible. IAM members or bindings should not use 'allUsers' or 'allAuthenticatedUsers'.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_topic_iam_member)

## Non-Compliant Code Examples
```terraform
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
```terraform
# IAM Binding compliant
resource "google_pubsub_topic_iam_binding" "good_example_binding" {
  topic   = "example-topic"
  members = ["user:someone@example.com", "group:admins@example.com"] # ✅ No public principals
  role    = "roles/pubsub.subscriber"
}

```

```terraform
# IAM Member compliant
resource "google_pubsub_topic_iam_member" "good_example_member" {
  topic  = "example-topic"
  member = "user:someone@example.com" # ✅ Non-public principal
  role   = "roles/pubsub.publisher"
}

```
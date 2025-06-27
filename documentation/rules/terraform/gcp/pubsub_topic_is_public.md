---
title: "Pub/Sub Topics are anonymously or publicly accessible"
meta:
  name: "gcp/pubsub_topic_is_public"
  id: "7sdj7dsj8-f348-4f95-845c-1090e41a615c"
  display_name: "Pub/Sub Topics are anonymously or publicly accessible"
  cloud_provider: "gcp"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Name:** `gcp/pubsub_topic_is_public`

**Query Name** `Pub/Sub Topics are anonymously or publicly accessible`

**Id:** `7sdj7dsj8-f348-4f95-845c-1090e41a615c`

**Cloud Provider:** gcp

**Platform** Terraform

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
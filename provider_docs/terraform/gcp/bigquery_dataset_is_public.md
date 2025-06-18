---
title: "BigQuery Dataset Is Public"
meta:
  name: "terraform/bigquery_dataset_is_public"
  id: "e576ce44-dd03-4022-a8c0-3906acca2ab4"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/bigquery_dataset_is_public`
**Id:** `e576ce44-dd03-4022-a8c0-3906acca2ab4`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
This check identifies BigQuery datasets that are configured to allow public or anonymous access, which exposes sensitive data to unauthorized users and increases the risk of data breaches. The vulnerability occurs when access controls use the special groups 'allAuthenticatedUsers' or 'allUsers', effectively making data available to anyone with a Google account or the general public. To secure your BigQuery dataset, restrict access to specific users, groups, or domains instead of using public access groups, as shown in the example below:

```terraform
access {
  role          = "OWNER"
  user_by_email = google_service_account.bqowner.email
}

access {
  role   = "READER"
  domain = "hashicorp.com"
}
```

#### Learn More

 - [Provider Reference](https://www.terraform.io/docs/providers/google/r/bigquery_dataset.html)

## Non-Compliant Code Examples
```gcp
resource "google_bigquery_dataset" "positive1" {
  dataset_id                  = "example_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    special_group = "allAuthenticatedUsers"
  }
}
```

## Compliant Code Examples
```gcp
# negative sample
resource "google_bigquery_dataset" "negative1" {
  dataset_id                  = "example_dataset"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }
}

```
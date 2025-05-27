---
title: "IAM Access Key Is Exposed"
meta:
  name: "aws/iam_access_key_is_exposed"
  id: "7081f85c-b94d-40fd-8b45-a4f1cac75e46"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/iam_access_key_is_exposed`

**Id:** `7081f85c-b94d-40fd-8b45-a4f1cac75e46`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
IAM Access Key should not be active for root users

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_access_key)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_access_key" "positive1" {
  user = "root"
  status = "Active"
}

resource "aws_iam_access_key" "positive2" {
  user = "root"
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_access_key" "negative1" {
  user = "some-user"
}

resource "aws_iam_access_key" "negative2" {
  user = "some-user"
  status = "Active"
}

resource "aws_iam_access_key" "negative3" {
  user = "root"
  status = "Inactive"
}

```
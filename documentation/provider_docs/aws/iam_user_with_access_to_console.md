---
title: "IAM User With Access To Console"
meta:
  name: "aws/iam_user_with_access_to_console"
  id: "9ec311bf-dfd9-421f-8498-0b063c8bc552"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/iam_user_with_access_to_console`

**Id:** `9ec311bf-dfd9-421f-8498-0b063c8bc552`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
AWS IAM Users should not have access to console

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_login_profile)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_user" "example" {
  name          = "example"
  path          = "/"
  force_destroy = true
}

resource "aws_iam_user_login_profile" "example_login" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_user" "example" {
  name          = "example"
  path          = "/"
  force_destroy = true
}

```
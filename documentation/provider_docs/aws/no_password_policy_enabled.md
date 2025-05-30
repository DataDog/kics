---
title: "No Password Policy Enabled"
meta:
  name: "aws/no_password_policy_enabled"
  id: "b592ffd4-0577-44b6-bd35-8c5ee81b5918"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/no_password_policy_enabled`

**Id:** `b592ffd4-0577-44b6-bd35-8c5ee81b5918`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
IAM password policies should be set through the password minimum length and reset password attributes

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_login_profile)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_user_login_profile" "positive2" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_reset_required = false

  password_length = 15
}

resource "aws_iam_user_login_profile" "positive3" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_reset_required = true

  password_length = 13
}

resource "aws_iam_user_login_profile" "positive6" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_length = 13
}

resource "aws_iam_user_login_profile" "positive7" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_reset_required = false
  password_length = 13
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_user_login_profile" "negative1" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_reset_required = true

  password_length = 15
}
```
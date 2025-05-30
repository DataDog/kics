---
title: "IAM Password Without Minimum Length"
meta:
  name: "aws/iam_password_without_minimum_length"
  id: "1bc1c685-e593-450e-88fb-19db4c82aa1d"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/iam_password_without_minimum_length`

**Id:** `1bc1c685-e593-450e-88fb-19db4c82aa1d`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
IAM password should have the required minimum length

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "positive1" {
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
}

resource "aws_iam_account_password_policy" "positive2" {
  minimum_password_length        = 3
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "negative1" {
  minimum_password_length        = 14
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
}

```
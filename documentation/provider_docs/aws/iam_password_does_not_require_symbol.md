---
title: "IAM Password Policy Does Not Require Symbol"
meta:
  name: "aws/iam_password_does_not_require_symbol"
  id: "bts2c3d4-e5f6-7890-ab12-cd34ef567890"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/iam_password_does_not_require_symbol`

**Id:** `bts2c3d4-e5f6-7890-ab12-cd34ef567890`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Best Practices

## Description
Ensures that the AWS IAM password policy requires at least one symbol. Without this setting, passwords may be easier to guess, leading to security vulnerabilities.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy#require_symbols)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "bad_example" {
  minimum_password_length      = 14
  require_symbols              = false
  require_numbers              = true
  require_lowercase_characters = true
  require_uppercase_characters = true
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "good_example" {
  minimum_password_length      = 14
  require_symbols              = true
  require_numbers              = true
  require_lowercase_characters = true
  require_uppercase_characters = true
}

```
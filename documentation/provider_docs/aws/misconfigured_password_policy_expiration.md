---
title: "Misconfigured Password Policy Expiration"
meta:
  name: "aws/misconfigured_password_policy_expiration"
  id: "ce60d060-efb8-4bfd-9cf7-ff8945d00d90"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/misconfigured_password_policy_expiration`

**Id:** `ce60d060-efb8-4bfd-9cf7-ff8945d00d90`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
No password expiration policy

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "positive1" {
  minimum_password_length        = 8
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
  max_password_age               = 180
}

// comment
resource "aws_iam_account_password_policy" "positive2" {
  minimum_password_length        = 8
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
  minimum_password_length        = 8
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
  max_password_age               = 10
}
```
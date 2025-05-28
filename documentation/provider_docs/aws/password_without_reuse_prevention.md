---
title: "Password Without Reuse Prevention"
meta:
  name: "aws/password_without_reuse_prevention"
  id: "89806cdc-9c2e-4bd1-a0dc-53f339bcfb2a"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/password_without_reuse_prevention`

**Id:** `89806cdc-9c2e-4bd1-a0dc-53f339bcfb2a`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
Check if IAM account password has the reuse password configured with 24

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy#password_reuse_prevention)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_account_password_policy" "positive1" {
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
  password_reuse_prevention = 20
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
  minimum_password_length        = 8
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
  password_reuse_prevention = 24
}

```
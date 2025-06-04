---
title: "AWS Password Policy With Unchangeable Passwords"
meta:
  name: "aws/aws_password_policy_with_unchangeable_passwords"
  id: "9ef7d25d-9764-4224-9968-fa321c56ef76"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/aws_password_policy_with_unchangeable_passwords`

**Id:** `9ef7d25d-9764-4224-9968-fa321c56ef76`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Insecure Configurations

## Description
Unchangeable passwords in AWS password policy

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_sqs_queue" "positive1" {
  name = "examplequeue"
}

// comment
resource "aws_iam_account_password_policy" "positive2" {
  minimum_password_length        = 8
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = false
}
```

## Compliant Code Examples
```terraform
resource "aws_sqs_queue" "negative1" {
  name = "examplequeue"
}

// comment
resource "aws_iam_account_password_policy" "negative2" {
  minimum_password_length        = 10
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
}
```
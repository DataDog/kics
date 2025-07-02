---
title: "IAM Password Policy Does Not Require Symbol"
meta:
  name: "aws/iam_password_does_not_require_symbol"
  id: "bts2c3d4-e5f6-7890-ab12-cd34ef567890"
  display_name: "IAM Password Policy Does Not Require Symbol"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Best Practices"
---
## Metadata

**Id:** `bts2c3d4-e5f6-7890-ab12-cd34ef567890`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy#require_symbols)

### Description

 This check ensures that the AWS IAM account password policy enforces the use of at least one symbol in user passwords by setting `require_symbols = true`. If `require_symbols` is set to `false`, as shown below:

```
resource "aws_iam_account_password_policy" "bad_example" {
  minimum_password_length      = 14
  require_symbols              = false
  require_numbers              = true
  require_lowercase_characters = true
  require_uppercase_characters = true
}
```

it weakens password complexity, making user accounts more susceptible to brute-force or password guessing attacks. Failing to enforce symbol usage increases the risk of unauthorized access to AWS resources.


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
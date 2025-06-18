---
title: "Misconfigured Password Policy Expiration"
meta:
  name: "terraform/misconfigured_password_policy_expiration"
  id: "ce60d060-efb8-4bfd-9cf7-ff8945d00d90"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/misconfigured_password_policy_expiration`
**Id:** `ce60d060-efb8-4bfd-9cf7-ff8945d00d90`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Best Practices
## Description
A password expiration policy enforces regular password changes, reducing the risk of compromised credentials being exploited over long periods. If the `aws_iam_account_password_policy` resource does not set the `max_password_age` attribute, as shown below,

```
resource "aws_iam_account_password_policy" "example" {
  minimum_password_length        = 8
  require_lowercase_characters   = true
  require_numbers                = true
  require_uppercase_characters   = true
  require_symbols                = true
  allow_users_to_change_password = true
  // max_password_age not set
}
```

passwords may remain valid indefinitely, increasing the chance that leaked or weak passwords can be used for unauthorized access. This exposes your AWS environment to persistent credential-related threats if not addressed.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_account_password_policy)

## Non-Compliant Code Examples
```aws
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
```aws
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
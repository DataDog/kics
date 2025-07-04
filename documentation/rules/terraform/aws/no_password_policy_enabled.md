---
title: "No Password Policy Enabled"
group-id: "rules/terraform/aws"
meta:
  name: "aws/no_password_policy_enabled"
  id: "b592ffd4-0577-44b6-bd35-8c5ee81b5918"
  display_name: "No Password Policy Enabled"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `b592ffd4-0577-44b6-bd35-8c5ee81b5918`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_login_profile)

### Description

 Ensuring strong AWS IAM password security involves configuring both the `password_length` and `password_reset_required` attributes in the `aws_iam_user_login_profile` resource. Failing to set a sufficient `password_length` or omitting the `password_reset_required = true` option, as shown below,

```
resource "aws_iam_user_login_profile" "example" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"
  password_length = 13
  password_reset_required = false
}
```

can lead to accounts being protected by weak or reused passwords, which increases the risk of unauthorized access. By requiring users to reset passwords on first use and enforcing adequate password length, as in

```
resource "aws_iam_user_login_profile" "example" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"
  password_length = 15
  password_reset_required = true
}
```

organizations can better defend against brute-force attacks and reduce credential compromise risk.


## Compliant Code Examples
```terraform
resource "aws_iam_user_login_profile" "negative1" {
  user    = aws_iam_user.example.name
  pgp_key = "keybase:some_person_that_exists"

  password_reset_required = true

  password_length = 15
}
```
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
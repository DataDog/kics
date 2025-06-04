---
title: "Cognito UserPool Without MFA"
meta:
  name: "aws/cognito_userpool_without_mfa"
  id: "ec28bf61-a474-4dbe-b414-6dd3a067d6f0"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/cognito_userpool_without_mfa`

**Id:** `ec28bf61-a474-4dbe-b414-6dd3a067d6f0`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
AWS Cognito UserPool should have MFA (Multi-Factor Authentication) defined to users

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cognito_user_pool)

## Non-Compliant Code Examples
```terraform
resource "aws_cognito_user_pool" "positive1" {
  # ... other configuration ...

  sms_authentication_message = "Your code is {####}"

  sms_configuration {
    external_id    = "example"
    sns_caller_arn = aws_iam_role.example.arn
  }

  software_token_mfa_configuration {
    enabled = true
  }
}

resource "aws_cognito_user_pool" "positive2" {
  # ... other configuration ...

  mfa_configuration          = "OFF"
  sms_authentication_message = "Your code is {####}"

  sms_configuration {
    external_id    = "example"
    sns_caller_arn = aws_iam_role.example.arn
  }

  software_token_mfa_configuration {
    enabled = true
  }
}

resource "aws_cognito_user_pool" "positive3" {
  # ... other configuration ...

  mfa_configuration          = "ON"
  sms_authentication_message = "Your code is {####}"
}

```

## Compliant Code Examples
```terraform
resource "aws_cognito_user_pool" "negative1" {
  # ... other configuration ...

  mfa_configuration          = "ON"
  sms_authentication_message = "Your code is {####}"

  sms_configuration {
    external_id    = "example"
    sns_caller_arn = aws_iam_role.example.arn
  }
}

resource "aws_cognito_user_pool" "negative2" {
  # ... other configuration ...

  mfa_configuration          = "OPTIONAL"
  sms_authentication_message = "Your code is {####}"

  software_token_mfa_configuration {
    enabled = true
  }
}

```
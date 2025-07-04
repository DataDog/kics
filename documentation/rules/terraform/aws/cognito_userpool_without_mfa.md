---
title: "Cognito UserPool Without MFA"
group-id: "rules/terraform/aws"
meta:
  name: "aws/cognito_userpool_without_mfa"
  id: "ec28bf61-a474-4dbe-b414-6dd3a067d6f0"
  display_name: "Cognito UserPool Without MFA"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata

**Id:** `ec28bf61-a474-4dbe-b414-6dd3a067d6f0`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Low

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cognito_user_pool)

### Description

 AWS Cognito User Pools should have Multi-Factor Authentication (MFA) enabled to enhance the security of user accounts. If the `mfa_configuration` attribute is set to `"OFF"` or left undefined, as in 

```
resource "aws_cognito_user_pool" "example" {
  mfa_configuration = "OFF"
  // ... other configuration ...
}
```

users are only required to use a single authentication factor, making their accounts more susceptible to unauthorized access if credentials are compromised. Enabling MFA (e.g., `mfa_configuration = "ON"` or `"OPTIONAL"`) significantly reduces the risk of account takeover by requiring an additional authentication factor.


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
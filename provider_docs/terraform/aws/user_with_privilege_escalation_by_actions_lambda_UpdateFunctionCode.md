---
title: "User With Privilege Escalation By Actions 'lambda:UpdateFunctionCode'"
meta:
  name: "terraform/user_with_privilege_escalation_by_actions_lambda_UpdateFunctionCode"
  id: "b69247e5-7e73-464e-ba74-ec9b715c6e12"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/user_with_privilege_escalation_by_actions_lambda_UpdateFunctionCode`
**Id:** `b69247e5-7e73-464e-ba74-ec9b715c6e12`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
This configuration grants a user the "lambda:UpdateFunctionCode" permission with the resource set to "*", allowing them to update the code of any Lambda function in the AWS account. Attackers or unauthorized users with this privilege can replace the code of existing Lambda functions with malicious code, which may be executed in response to legitimate triggers or schedules. This type of privilege escalation can allow an attacker to gain unauthorized access to sensitive data, further compromise the AWS environment, or establish persistent backdoors within serverless resources. If left unaddressed, this vulnerability exposes the environment to substantial risk of code injection, data exfiltration, and lateral movement within the AWS account.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy#policy)

## Non-Compliant Code Examples
```aws
resource "aws_iam_user" "cosmic" {
  name = "cosmic"
}

resource "aws_iam_user_policy" "test_inline_policy" {
  name = "test_inline_policy"
  user = aws_iam_user.cosmic.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "lambda:UpdateFunctionCode",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "aws_iam_policy_attachment" "test-attach" {
  name       = "test-attachment"
  users      = [aws_iam_user.cosmic.name]
  roles      = [aws_iam_role.role.name]
  groups     = [aws_iam_group.group.name]
}


```

## Compliant Code Examples
```aws
resource "aws_iam_user" "cosmic2" {
  name = "cosmic2"
}

resource "aws_iam_user_policy" "inline_policy_run_instances2" {
  name = "inline_policy_run_instances"
  user = aws_iam_user.cosmic2.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

```
---
title: "Group With Privilege Escalation By Actions 'lambda:UpdateFunctionCode'"
meta:
  name: "aws/group_with_privilege_escalation_by_actions_lambda_UpdateFunctionCode"
  id: "571254d8-aa6a-432e-9725-535d3ef04d69"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/group_with_privilege_escalation_by_actions_lambda_UpdateFunctionCode`

**Id:** `571254d8-aa6a-432e-9725-535d3ef04d69`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Group with privilege escalation by actions 'lambda:UpdateFunctionCode' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_group_policy#policy)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_group" "cosmic" {
  name = "cosmic"
}

resource "aws_iam_group_policy" "test_inline_policy" {
  name = "test_inline_policy"
  group = aws_iam_group.cosmic.name

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

```

## Compliant Code Examples
```terraform
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
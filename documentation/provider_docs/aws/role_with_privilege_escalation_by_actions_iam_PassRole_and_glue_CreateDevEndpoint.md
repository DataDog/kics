---
title: "Role With Privilege Escalation By Actions 'glue:CreateDevEndpoint' And 'iam:PassRole'"
meta:
  name: "aws/role_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint"
  id: "0a592060-8166-49f5-8e65-99ac6dce9871"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/role_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint`

**Id:** `0a592060-8166-49f5-8e65-99ac6dce9871`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Role with privilege escalation by actions 'glue:CreateDevEndpoint' and 'iam:PassRole' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy#policy)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_role" "cosmic" {
  name = "cosmic"
}

resource "aws_iam_role_policy" "test_inline_policy" {
  name = "test_inline_policy"
  role = aws_iam_role.cosmic.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "glue:CreateDevEndpoint",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "aws_iam_policy_attachment" "test-attach" {
  name       = "test-attachment"
  roles      = [aws_iam_role.cosmic.name]
  policy_arn = aws_iam_policy.policy.arn
}


resource "aws_iam_policy" "policy" {
  name        = "test-policy"
  description = "A test policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "iam:PassRole",
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
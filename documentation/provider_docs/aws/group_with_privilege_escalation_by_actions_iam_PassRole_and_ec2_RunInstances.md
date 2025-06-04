---
title: "Group With Privilege Escalation By Actions 'ec2:RunInstances' And 'iam:PassRole'"
meta:
  name: "aws/group_with_privilege_escalation_by_actions_iam_PassRole_and_ec2_RunInstances"
  id: "15e6ad8c-f420-49a6-bafb-074f5eb1ec74"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/group_with_privilege_escalation_by_actions_iam_PassRole_and_ec2_RunInstances`

**Id:** `15e6ad8c-f420-49a6-bafb-074f5eb1ec74`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Group with privilege escalation by actions 'ec2:RunInstances' and 'iam:PassRole' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

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
          "ec2:RunInstances",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}


resource "aws_iam_policy_attachment" "test-attach" {
  name       = "test-attachment"
  groups     = [aws_iam_group.cosmic.name]
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
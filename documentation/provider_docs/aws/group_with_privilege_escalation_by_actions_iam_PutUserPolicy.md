---
title: "Group With Privilege Escalation By Actions 'iam:PutUserPolicy'"
meta:
  name: "aws/group_with_privilege_escalation_by_actions_iam_PutUserPolicy"
  id: "60263b4a-6801-4587-911d-919c37ed733b"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/group_with_privilege_escalation_by_actions_iam_PutUserPolicy`

**Id:** `60263b4a-6801-4587-911d-919c37ed733b`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Group with privilege escalation by actions 'iam:PutUserPolicy' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

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
          "iam:PutUserPolicy",
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
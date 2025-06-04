---
title: "User With Privilege Escalation By Actions 'iam:UpdateAssumeRolePolicy' And 'sts:AssumeRole'"
meta:
  name: "aws/user_with_privilege_escalation_by_actions_iam_UpdateAssumeRolePolicy_and_sts_AssumeRole"
  id: "33627268-1445-4385-988a-318fd9d1a512"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/user_with_privilege_escalation_by_actions_iam_UpdateAssumeRolePolicy_and_sts_AssumeRole`

**Id:** `33627268-1445-4385-988a-318fd9d1a512`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
User with privilege escalation by actions 'iam:UpdateAssumeRolePolicy' and 'sts:AssumeRole' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy#policy)

## Non-Compliant Code Examples
```terraform
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
          "iam:UpdateAssumeRolePolicy",
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
          "sts:AssumeRole",
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
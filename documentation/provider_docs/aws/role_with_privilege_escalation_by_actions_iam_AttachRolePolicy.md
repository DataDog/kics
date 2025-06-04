---
title: "Role With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
meta:
  name: "aws/role_with_privilege_escalation_by_actions_iam_AttachRolePolicy"
  id: "f465fff1-0a0f-457d-aa4d-1bddb6f204ff"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/role_with_privilege_escalation_by_actions_iam_AttachRolePolicy`

**Id:** `f465fff1-0a0f-457d-aa4d-1bddb6f204ff`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Role with privilege escalation by actions 'iam:AttachRolePolicy' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

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
          "iam:AttachRolePolicy",
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
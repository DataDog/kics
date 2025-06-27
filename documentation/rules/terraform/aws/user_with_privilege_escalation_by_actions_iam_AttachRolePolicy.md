---
title: "User With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
meta:
  name: "aws/user_with_privilege_escalation_by_actions_iam_AttachRolePolicy"
  id: "e227091e-2228-4b40-b046-fc13650d8e88"
  display_name: "User With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `aws/user_with_privilege_escalation_by_actions_iam_AttachRolePolicy`
**Query Name** `User With Privilege Escalation By Actions 'iam:AttachRolePolicy'`
**Id:** `e227091e-2228-4b40-b046-fc13650d8e88`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting the `iam:AttachRolePolicy` action with a resource value of `"*"` allows the user to attach any policy to any IAM role, enabling privilege escalation if the user can attach policies granting additional permissions or even administrator-level access. This misconfiguration can lead to unauthorized access across your AWS environment, as users may obtain permissions far beyond their original scope. To remediate, policy statements should scope the `Resource` field to only those roles and policies necessary for the user’s legitimate activities, avoiding the use of a wildcard (`*`).

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy#policy)


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
          "iam:AttachRolePolicy",
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
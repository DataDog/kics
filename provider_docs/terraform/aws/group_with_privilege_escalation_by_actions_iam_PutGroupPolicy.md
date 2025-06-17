---
title: "Group With Privilege Escalation By Actions 'iam:PutGroupPolicy'"
meta:
  name: "terraform/group_with_privilege_escalation_by_actions_iam_PutGroupPolicy"
  id: "e77c89f6-9c85-49ea-b95b-5f960fe5be92"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/group_with_privilege_escalation_by_actions_iam_PutGroupPolicy`
**Id:** `e77c89f6-9c85-49ea-b95b-5f960fe5be92`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
This check identifies IAM group policies that grant the action `iam:PutGroupPolicy` with a resource set to `"*"`, allowing anyone in the group to attach arbitrary inline policies to any IAM group. This presents a privilege escalation risk, as users with this permission could assign themselves broader or unauthorized permissions by updating policies on other groups. To mitigate this, restrict the `Action` and `Resource` fields in policies and avoid assigning sensitive permissions to groups broadly.

A secure Terraform configuration should limit permissions to only what is necessary, for example:

```
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

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_group_policy#policy)

## Non-Compliant Code Examples
```aws
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
          "iam:PutGroupPolicy",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
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
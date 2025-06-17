---
title: "Group With Privilege Escalation By Actions 'iam:PutRolePolicy'"
meta:
  name: "terraform/group_with_privilege_escalation_by_actions_iam_PutRolePolicy"
  id: "c0c1e744-0f37-445e-924a-1846f0839f69"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/group_with_privilege_escalation_by_actions_iam_PutRolePolicy`
**Id:** `c0c1e744-0f37-445e-924a-1846f0839f69`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
This check looks for AWS IAM groups with policies that allow the `iam:PutRolePolicy` action on all resources (`"Resource": "*"`) in Terraform code. Granting this privilege means members of the group can attach any policy to any IAM role in the account, enabling easy privilege escalation or the creation of backdoors. If left unaddressed, attackers or unprivileged users could use this access to gain administrative permissions or compromise critical resources, leading to severe security risks.

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
          "iam:PutRolePolicy",
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
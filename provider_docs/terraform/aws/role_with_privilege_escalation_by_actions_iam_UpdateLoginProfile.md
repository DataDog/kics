---
title: "Role With Privilege Escalation By Actions 'iam:UpdateLoginProfile'"
meta:
  name: "terraform/role_with_privilege_escalation_by_actions_iam_UpdateLoginProfile"
  id: "35ccf766-0e4d-41ed-9ec4-2dab155082b4"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/role_with_privilege_escalation_by_actions_iam_UpdateLoginProfile`
**Id:** `35ccf766-0e4d-41ed-9ec4-2dab155082b4`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting the `iam:UpdateLoginProfile` action with the `Resource` attribute set to `"*"` in an IAM policy allows a role or user to change the login passwords of any IAM user in the AWS account. This creates a privilege escalation vulnerability, as a user with this permission could take over other accounts and gain unauthorized access to critical resources. 

To prevent this, use least-privilege access—for example, restricting allowed actions and resource scope, as in:

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
This sample policy grants only read-only access to EC2 information, rather than sensitive IAM actions.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy#policy)

## Non-Compliant Code Examples
```aws
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
          "iam:UpdateLoginProfile",
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
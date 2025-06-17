---
title: "User With Privilege Escalation By Actions 'iam:CreateLoginProfile'"
meta:
  name: "terraform/user_with_privilege_escalation_by_actions_iam_CreateLoginProfile"
  id: "0fd7d920-4711-46bd-aff2-d307d82cd8b7"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/user_with_privilege_escalation_by_actions_iam_CreateLoginProfile`
**Id:** `0fd7d920-4711-46bd-aff2-d307d82cd8b7`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing users the `iam:CreateLoginProfile` action with the `Resource` set to `"*"` in AWS IAM policies permits them to set a password for any IAM user, thereby enabling direct console access. This creates a privilege escalation vulnerability, as the user can potentially assign login profiles to high-privilege accounts, leading to unauthorized access and control over critical AWS resources. To mitigate this risk, restrict the `Resource` to specific user ARNs and avoid assigning broad permissions, as shown in a secure Terraform configuration:

```
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
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy#policy)

## Non-Compliant Code Examples
```aws
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
          "iam:CreateLoginProfile",
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
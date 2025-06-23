---
title: "Role With Privilege Escalation By Actions 'iam:CreateLoginProfile'"
meta:
  name: "aws/role_with_privilege_escalation_by_actions_iam_CreateLoginProfile"
  id: "9a205ba3-0dd1-42eb-8d54-2ffec836b51a"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `aws/role_with_privilege_escalation_by_actions_iam_CreateLoginProfile`
**Id:** `9a205ba3-0dd1-42eb-8d54-2ffec836b51a`
**Cloud Provider:** aws
**Severity:** Medium
**Category:** Access Control
## Description
This check identifies if an AWS IAM role policy grants the `iam:CreateLoginProfile` action with a `Resource` set to `"*"`, as shown below:

```
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
```

Allowing this action enables a user or role to create login profiles (i.e., set console passwords) for any IAM user account, potentially facilitating privilege escalation. If left unaddressed, attackers could exploit this permission to gain unauthorized console access and take full control of IAM users, posing a severe security risk to the environment.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy#policy)


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
          "iam:CreateLoginProfile",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}



```
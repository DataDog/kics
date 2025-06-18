---
title: "Role With Privilege Escalation By Actions 'iam:PutRolePolicy'"
meta:
  name: "terraform/role_with_privilege_escalation_by_actions_iam_PutRolePolicy"
  id: "eb64f1e9-f67d-4e35-8a3c-3d6a2f9efea7"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/role_with_privilege_escalation_by_actions_iam_PutRolePolicy`
**Id:** `eb64f1e9-f67d-4e35-8a3c-3d6a2f9efea7`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting a role the `iam:PutRolePolicy` action with the `Resource` set to `"*"` in Terraform, as in the following configuration,

```
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
```

allows any principal with this role to attach arbitrary permissions to any IAM role, leading to potential privilege escalation. An attacker could exploit this to grant themselves or others broad or administrative permissions, compromising the security of the AWS environment. It is critical to scope IAM permissions as narrowly as possible, restricting both actions and resources to mitigate this risk.

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
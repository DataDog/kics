---
title: "User With Privilege Escalation By Actions 'iam:PutRolePolicy'"
meta:
  name: "terraform/user_with_privilege_escalation_by_actions_iam_PutRolePolicy"
  id: "eeb4d37a-3c59-4789-a00c-1509bc3af1e5"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/user_with_privilege_escalation_by_actions_iam_PutRolePolicy`
**Id:** `eeb4d37a-3c59-4789-a00c-1509bc3af1e5`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing a user the `iam:PutRolePolicy` action on all resources (i.e., `"Resource": "*"`) enables them to attach inline policies to any IAM role in the AWS environment. This grants the user a privilege escalation path, as they could grant overly broad or administrative permissions to roles they do not own, potentially gaining full control over the AWS account. To mitigate this risk, restrict the `iam:PutRolePolicy` action using least privilege and avoid using wildcards in the `Resource` attribute, as shown below:

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
          "iam:PutRolePolicy",
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
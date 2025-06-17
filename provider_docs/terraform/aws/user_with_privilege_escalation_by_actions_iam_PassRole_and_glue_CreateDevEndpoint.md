---
title: "User With Privilege Escalation By Actions 'glue:CreateDevEndpoint' And 'iam:PassRole'"
meta:
  name: "terraform/user_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint"
  id: "94fbe150-27e3-4eba-9ca6-af32865e4503"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/user_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint`
**Id:** `94fbe150-27e3-4eba-9ca6-af32865e4503`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting an IAM user permissions for both `glue:CreateDevEndpoint` and `iam:PassRole` actions with the `Resource` set to `"*"` allows that user to create a Glue development endpoint and attach any role in the account, including those with elevated privileges. This misconfiguration enables privilege escalation, meaning the user could effectively gain administrator-level access or perform unauthorized actions across the AWS environment. To prevent this, restrict the actions and resources in IAM policies and avoid using wildcard `"*"` permissions, as in the secure example below:

```
resource "aws_iam_user_policy" "inline_policy_limited" {
  name = "inline_policy_limited"
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
          "glue:CreateDevEndpoint",
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
          "iam:PassRole",
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
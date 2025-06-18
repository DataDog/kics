---
title: "Role With Privilege Escalation By Actions 'glue:CreateDevEndpoint' And 'iam:PassRole'"
meta:
  name: "terraform/role_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint"
  id: "0a592060-8166-49f5-8e65-99ac6dce9871"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/role_with_privilege_escalation_by_actions_iam_PassRole_and_glue_CreateDevEndpoint`
**Id:** `0a592060-8166-49f5-8e65-99ac6dce9871`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting an AWS IAM role permissions for both `glue:CreateDevEndpoint` and `iam:PassRole` with the `Resource` attribute set to `"*"` allows for privilege escalation within an AWS environment. With these permissions, a user or attacker could create Glue DevEndpoints and assign any AWS IAM role to the endpoint, effectively running arbitrary code with higher privileges by passing roles they may not otherwise have access to. The use of the `"iam:PassRole"` action combined with a resource wildcard means that the role can be used to assign any role in the account, potentially including administrative or sensitive roles. If left unaddressed, this misconfiguration can lead to an attacker gaining full control over AWS resources, resulting in data breaches or the compromise of critical cloud infrastructure.

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
  roles      = [aws_iam_role.cosmic.name]
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
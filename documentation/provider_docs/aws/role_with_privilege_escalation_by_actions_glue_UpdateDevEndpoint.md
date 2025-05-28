---
title: "Role With Privilege Escalation By Actions 'glue:UpdateDevEndpoint'"
meta:
  name: "aws/role_with_privilege_escalation_by_actions_glue_UpdateDevEndpoint"
  id: "eda48c88-2b7d-4e34-b6ca-04c0194aee17"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/role_with_privilege_escalation_by_actions_glue_UpdateDevEndpoint`

**Id:** `eda48c88-2b7d-4e34-b6ca-04c0194aee17`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Role with privilege escalation by actions 'glue:UpdateDevEndpoint' and Resource set to '*'. For more information see https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy#policy)

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
          "glue:UpdateDevEndpoint",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

```

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
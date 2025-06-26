---
title: "Role With Privilege Escalation By Actions 'iam:CreatePolicyVersion'"
meta:
  name: "terraform/role_with_privilege_escalation_by_actions_iam_CreatePolicyVersion"
  id: "ee49557d-750c-4cc1-aa95-94ab36cbefde"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/role_with_privilege_escalation_by_actions_iam_CreatePolicyVersion`
**Id:** `ee49557d-750c-4cc1-aa95-94ab36cbefde`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting the `iam:CreatePolicyVersion` action with a resource of `"*"` in an IAM policy allows a user to create new policy versions for *any* policy in the AWS account, including those attached to highly privileged roles. This capability can be exploited for privilege escalation, as a malicious or compromised user could attach or update policies to grant themselves broader permissions. To mitigate this risk, restrict the `Resource` attribute to specific policy ARNs and only grant this action to trusted administrative principals.

A more secure configuration would specify only necessary actions for the specific resources required, for example:

```
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
          "iam:CreatePolicyVersion",
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
---
title: "Group With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
meta:
  name: "terraform/group_with_privilege_escalation_by_actions_iam_AttachRolePolicy"
  id: "3dd96caa-0b5f-4a85-b929-acfac4646cc2"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/group_with_privilege_escalation_by_actions_iam_AttachRolePolicy`
**Id:** `3dd96caa-0b5f-4a85-b929-acfac4646cc2`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
This configuration grants an IAM group the permission to attach any policy to any role in the AWS account, as indicated by the action 'iam:AttachRolePolicy' with the resource set to '*'. This creates a significant privilege escalation risk: any member of this group could grant themselves or others broader access, including administrative permissions, by attaching highly privileged policies to roles. If left unaddressed, this misconfiguration makes it possible for malicious or compromised users to take full control of AWS resources, potentially leading to data breaches or loss of service integrity. Restricting such permissions only to trusted entities and minimizing their scope is critical to maintaining a secure AWS environment.

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
          "iam:AttachRolePolicy",
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
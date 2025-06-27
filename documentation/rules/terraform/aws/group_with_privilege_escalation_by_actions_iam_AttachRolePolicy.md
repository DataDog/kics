---
title: "Group With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
meta:
  name: "aws/group_with_privilege_escalation_by_actions_iam_AttachRolePolicy"
  id: "3dd96caa-0b5f-4a85-b929-acfac4646cc2"
  display_name: "Group With Privilege Escalation By Actions 'iam:AttachRolePolicy'"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `aws/group_with_privilege_escalation_by_actions_iam_AttachRolePolicy`
**Query Name** `Group With Privilege Escalation By Actions 'iam:AttachRolePolicy'`
**Id:** `3dd96caa-0b5f-4a85-b929-acfac4646cc2`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Medium
**Category:** Access Control
## Description
This configuration grants an IAM group the permission to attach any policy to any role in the AWS account, as indicated by the action 'iam:AttachRolePolicy' with the resource set to '*'. This creates a significant privilege escalation risk: any member of this group could grant themselves or others broader access, including administrative permissions, by attaching highly privileged policies to roles. If left unaddressed, this misconfiguration makes it possible for malicious or compromised users to take full control of AWS resources, potentially leading to data breaches or loss of service integrity. Restricting such permissions only to trusted entities and minimizing their scope is critical to maintaining a secure AWS environment.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_group_policy#policy)


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
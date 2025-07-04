---
title: "User With Privilege Escalation By Actions 'iam:SetDefaultPolicyVersion'"
group-id: "rules/terraform/aws"
meta:
  name: "aws/user_with_privilege_escalation_by_actions_iam_SetDefaultPolicyVersion"
  id: "43a41523-386a-4cb1-becb-42af6b414433"
  display_name: "User With Privilege Escalation By Actions 'iam:SetDefaultPolicyVersion'"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `43a41523-386a-4cb1-becb-42af6b414433`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_user_policy#policy)

### Description

 Granting the `iam:SetDefaultPolicyVersion` action with a resource of `"*"` allows a user to set any version of any IAM policy as the default, including attaching more permissive versions to roles or users. This creates a serious privilege escalation risk, as an attacker with these permissions could assign themselves or others elevated privileges by setting a policy version that permits broader or unauthorized access. If left unaddressed, this vulnerability could lead to full account compromise or unapproved actions throughout the AWS environment.


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
          "iam:SetDefaultPolicyVersion",
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
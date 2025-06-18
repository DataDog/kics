---
title: "User With Privilege Escalation By Actions 'iam:SetDefaultPolicyVersion'"
meta:
  name: "terraform/user_with_privilege_escalation_by_actions_iam_SetDefaultPolicyVersion"
  id: "43a41523-386a-4cb1-becb-42af6b414433"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/user_with_privilege_escalation_by_actions_iam_SetDefaultPolicyVersion`
**Id:** `43a41523-386a-4cb1-becb-42af6b414433`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Granting the `iam:SetDefaultPolicyVersion` action with a resource of `"*"` allows a user to set any version of any IAM policy as the default, including attaching more permissive versions to roles or users. This creates a serious privilege escalation risk, as an attacker with these permissions could assign themselves or others elevated privileges by setting a policy version that permits broader or unauthorized access. If left unaddressed, this vulnerability could lead to full account compromise or unapproved actions throughout the AWS environment.

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
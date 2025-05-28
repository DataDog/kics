---
title: "SSO Policy with full privileges"
meta:
  name: "aws/sso_policy_with_full_priveleges"
  id: "132a8c31-9837-4203-9fd1-15ca210c7b73"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/sso_policy_with_full_priveleges`

**Id:** `132a8c31-9837-4203-9fd1-15ca210c7b73`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
SSO policies should be configured to grant limited administrative privileges, rather than full access to all resources. This approach allows for better security and control over the resources being accessed.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ssoadmin_permission_set_inline_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_ssoadmin_permission_set_inline_policy" "pos1" {
  instance_arn       = aws_ssoadmin_permission_set.example.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.example.arn
  inline_policy = <<POLICY
{
  "Statement": [
    {
      "Action": [
        "*"
      ],
      "Effect": "Allow",
      "Resource": [
        "*"
      ],
      "Sid": ""
    },
    {
      "Action": "s3:PutObject*",
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::b1/*",
      "Sid": ""
    }
  ],
  "Version": "2012-10-17"
}
POLICY
}

```

## Compliant Code Examples
```terraform
resource "aws_ssoadmin_permission_set_inline_policy" "neg1" {
  instance_arn       = aws_ssoadmin_permission_set.example.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.example.arn
  inline_policy = <<POLICY
{
  "Statement": [
    {
      "Action": [
        "s3:ListBucket*",
        "s3:HeadBucket",
        "s3:Get*"
      ],
      "Effect": "Allow",
      "Resource": [
        "arn:aws:s3:::b1",
        "arn:aws:s3:::b1/*",
        "arn:aws:s3:::b2",
        "arn:aws:s3:::b2/*"
      ],
      "Sid": ""
    },
    {
      "Action": "s3:PutObject*",
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::b1/*",
      "Sid": ""
    }
  ],
  "Version": "2012-10-17"
}
POLICY
}

```
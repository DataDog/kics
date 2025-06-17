---
title: "SES Policy With Allowed IAM Actions"
meta:
  name: "terraform/ses_policy_with_allowed_iam_actions"
  id: "34b921bd-90a0-402e-a0a5-dc73371fd963"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/ses_policy_with_allowed_iam_actions`
**Id:** `34b921bd-90a0-402e-a0a5-dc73371fd963`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
Amazon SES identity policies that allow access to all principals ('AWS': '*') create a significant security risk by granting any AWS account permissions to perform the specified actions on your SES resources. This overly permissive configuration can lead to unauthorized email sending, potential data breaches, and could be exploited for phishing campaigns or reputation damage.

Instead, SES policies should explicitly specify the ARNs of trusted principals that require access. For example, replace `"AWS": "*"` with `"AWS": "arn:aws:iam::ACCOUNT_ID:root"` or more specific IAM roles/users to implement the principle of least privilege.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ses_identity_policy#policy)

## Non-Compliant Code Examples
```aws
resource "aws_ses_identity_policy" "positive1" {
  identity = aws_ses_domain_identity.example.arn
  name     = "example"
  policy   = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "*",
      "Principal": {
        "AWS": "*"
      },
      "Effect": "Allow",
      "Resource": "*",
      "Sid": ""
    }
  ]
}
EOF
}

```

## Compliant Code Examples
```aws
resource "aws_ses_identity_policy" "negative1" {
  identity = aws_ses_domain_identity.example.arn
  name     = "example"
  policy   = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "*",
      "Principal": {
        "AWS": "arn:aws:iam::987654321145:root"
      },
      "Effect": "Allow",
      "Resource": "*",
      "Sid": ""
    }
  ]
}
EOF
}

```
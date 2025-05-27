---
title: "IAM Policies With Full Privileges"
meta:
  name: "aws/iam_policies_with_full_privileges"
  id: "2f37c4a3-58b9-4afe-8a87-d7f1d2286f84"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/iam_policies_with_full_privileges`

**Id:** `2f37c4a3-58b9-4afe-8a87-d7f1d2286f84`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
IAM policies shouldn't allow full administrative privileges (for all resources)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "positive1" {
  name = "apigateway-cloudwatch-logging"
  role = aws_iam_role.apigateway_cloudwatch_logging.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["*"],
      "Resource": "*"
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    effect = "Allow"
    actions = [
      "*"
    ]
    resources = [
      "*",
    ]
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_role_policy" "negative1" {
  name = "apigateway-cloudwatch-logging"
  role = aws_iam_role.apigateway_cloudwatch_logging.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["some:action"],
      "Resource": "*"
    }
  ]
}
EOF
}
data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    effect = "Allow"
    actions = [
      "*"
    ]
    resources = [
      "arn:aws:s3:::*",
    ]
  }
}

```
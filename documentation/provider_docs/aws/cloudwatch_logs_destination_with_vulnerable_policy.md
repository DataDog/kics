---
title: "CloudWatch Logs Destination With Vulnerable Policy"
meta:
  name: "aws/cloudwatch_logs_destination_with_vulnerable_policy"
  id: "db0ec4c4-852c-46a2-b4f3-7ec13cdb12a8"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Access Control"
---

## Metadata
**Name:** `aws/cloudwatch_logs_destination_with_vulnerable_policy`

**Id:** `db0ec4c4-852c-46a2-b4f3-7ec13cdb12a8`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Access Control

## Description
CloudWatch Logs destination policy should avoid wildcard in 'principals' and 'actions'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_destination_policy#access_policy)

## Non-Compliant Code Examples
```terraform
data "aws_iam_policy_document" "test_destination_policy" {
  statement {
    effect = "Allow"

    principals {
      type = "AWS"

      identifiers = [
        data.aws_caller_identity.current.id,
      ]
    }

    actions = [
      "logs:*",
    ]

  }
}

resource "aws_cloudwatch_log_destination_policy" "test_destination_policy" {
  destination_name = aws_cloudwatch_log_destination.test_destination.name
  access_policy    = data.aws_iam_policy_document.test_destination_policy.json
}

```

## Compliant Code Examples
```terraform
data "aws_iam_policy_document" "test_destination_policy2" {
  statement {
    effect = "Allow"

    principals {
      type = "AWS"

      identifiers = [
        "123456789012",
      ]
    }

    actions = [
      "logs:PutSubscriptionFilter",
    ]

    resources = [
      aws_cloudwatch_log_destination.test_destination.arn,
    ]
  }
}

resource "aws_cloudwatch_log_destination_policy" "test_destination_policy2" {
  destination_name = aws_cloudwatch_log_destination.test_destination.name
  access_policy    = data.aws_iam_policy_document.test_destination_policy2.json
}

```
---
title: "Lambda Permission Misconfigured"
meta:
  name: "aws/lambda_permission_misconfigured"
  id: "75ec6890-83af-4bf1-9f16-e83726df0bd0"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/lambda_permission_misconfigured`

**Id:** `75ec6890-83af-4bf1-9f16-e83726df0bd0`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Best Practices

## Description
Lambda permission may be misconfigured if the action field is not filled in by 'lambda:InvokeFunction'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

## Non-Compliant Code Examples
```terraform
resource "aws_lambda_permission" "positive1" {
  action        = "lambda:DeleteFunction"
  function_name = aws_lambda_function.logging.function_name
  principal     = "logs.eu-west-1.amazonaws.com"
  source_arn    = "${aws_cloudwatch_log_group.default.arn}:*"
}

```

## Compliant Code Examples
```terraform
resource "aws_lambda_permission" "negative1" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.logging.function_name
  principal     = "logs.eu-west-1.amazonaws.com"
  source_arn    = "${aws_cloudwatch_log_group.default.arn}:*"
}

```